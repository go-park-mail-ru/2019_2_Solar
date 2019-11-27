package functions

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/balancer"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"io"
	"log"
	"strconv"
	"time"
)

type CryptToken struct {
	Secret []byte
}

type TokenData struct {
	SessionID string
	UserID    uint
	Exp       int64
}

type Session struct {
	UserID uint
	ID     string
}

func NewAesCryptHashToken(secret string) (*CryptToken, error) {
	key := []byte(secret)
	_, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("cypher problem %v", err)
	}
	return &CryptToken{Secret: key}, nil
}

func (tk *CryptToken) Create(s *Session, tokenExpTime int64) (string, error) {
	block, err := aes.NewCipher(tk.Secret)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	td := &TokenData{SessionID: s.ID, UserID: s.UserID, Exp: tokenExpTime}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

func (tk *CryptToken) Check(s *Session, inputToken string) (bool, error) {
	block, err := aes.NewCipher(tk.Secret)
	if err != nil {
		return false, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		return false, err
	}
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, fmt.Errorf("decrypt fail: %v", err)
	}

	td := TokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		return false, fmt.Errorf("bad json: %v", err)
	}

	if td.Exp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	expected := TokenData{SessionID: s.ID, UserID: s.UserID}
	td.Exp = 0
	return td == expected, nil
}

var (
	consulAddr = "127.0.0.1:8500"
)

var (
	consul       *consulapi.Client
	nameResolver *balancer.TestNameResolver
)

type Auth struct {
	 grpcConn *grpc.ClientConn
	 Client services.AuthorizationServiceClient
}

func (auth *Auth) AuthServiceCreate() () {
	var err error
	config := consulapi.DefaultConfig()
	config.Address = consulAddr
	consul, err = consulapi.NewClient(config)

	health, _, err := consul.Health().Service("authorization-service", "", false, nil)
	if err != nil {
		log.Fatalf("cant get alive services")
	}

	servers := []string{}
	for _, item := range health {
		addr := item.Service.Address +
			":" + strconv.Itoa(item.Service.Port)
		servers = append(servers, addr)
	}

	nameResolver = &balancer.TestNameResolver{
		Addr: servers[0],
	}

	auth.grpcConn, err = grpc.Dial(
		servers[0],
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

	if len(servers) > 1 {
		var updates []*naming.Update
		for i := 1; i < len(servers); i++ {
			updates = append(updates, &naming.Update{
				Op:   naming.Add,
				Addr: servers[i],
			})
		}
		nameResolver.W.Inject(updates)
	}

	auth.Client = services.NewAuthorizationServiceClient(auth.grpcConn)

	// тут мы будем периодически опрашивать консул на предмет изменений
	go RunOnlineServiceDiscovery(servers)
	//testUserReg := services.UserReg{
	//	Email:                "test@mail.ru",
	//	Password:             "12314ghMEFnk123",
	//	Username:             "Test",
	//	XXX_NoUnkeyedLiteral: struct{}{},
	//	XXX_unrecognized:     nil,
	//	XXX_sizecache:        0,
	//}

	//ctx := context.Background()
	//step := 1
	//for {
	//	// проверяем несуществуюущую сессию
	//	// потому что сейчас между сервисами нет общения
	//	// получаем загшулку
	//	sess, err := auth.Client.CheckSession(ctx,
	//		&services.Cookie{
	//			Key: "test",
	//			Value: string(step),
	//			Exp: time.Now().String(),
	//		})
	//	fmt.Println("get sess", step, sess, err)
	//
	//	time.Sleep(1500 * time.Millisecond)
	//	step++
	//}



	//TestRun(auth.grpcConn)
	//testRun(grcpConn)
	//sessManager := services.NewAuthorizationServiceClient(grcpConn)

	//return grcpConn
}

func TestRun(grcpConn *grpc.ClientConn) {
	auth := services.NewAuthorizationServiceClient(grcpConn)
	ctx := context.Background()
	step := 1
	for {
		// проверяем несуществуюущую сессию
		// потому что сейчас между сервисами нет общения
		// получаем загшулку
		sess, err := auth.CheckSession(ctx,
			&services.Cookie{
				Key: "test",
				Value: string(step),
				Exp: time.Now().String(),
			})
		fmt.Println("get sess", step, sess, err)

		time.Sleep(1500 * time.Millisecond)
		step++
	}
}

func RunOnlineServiceDiscovery(servers []string) {
	currAddrs := make(map[string]struct{}, len(servers))
	for _, addr := range servers {
		currAddrs[addr] = struct{}{}
	}
	ticker := time.Tick(5 * time.Second)
	for _ = range ticker {
		health, _, err := consul.Health().Service("authorization-service", "", false, nil)
		if err != nil {
			log.Fatalf("cant get alive services")
		}

		newAddrs := make(map[string]struct{}, len(health))
		for _, item := range health {
			addr := item.Service.Address +
				":" + strconv.Itoa(item.Service.Port)
			newAddrs[addr] = struct{}{}
		}

		var updates []*naming.Update
		// проверяем что удалилось
		for addr := range currAddrs {
			if _, exist := newAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Delete,
					Addr: addr,
				})
				delete(currAddrs, addr)
				fmt.Println("remove", addr)
			}
		}
		// проверяем что добавилось
		for addr := range newAddrs {
			if _, exist := currAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Add,
					Addr: addr,
				})
				currAddrs[addr] = struct{}{}
				fmt.Println("add", addr)
			}
		}
		if len(updates) > 0 {
			nameResolver.W.Inject(updates)
		}
	}
}

