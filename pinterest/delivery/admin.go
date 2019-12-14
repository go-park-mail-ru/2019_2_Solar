package delivery

import (
	"bytes"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func (h *HandlersStruct) HandleAdminFill(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")

	myUsers := []struct{
		Username string
		UserID	uint64
		boardID uint64
	}{
		{
			"ADN",
			0,
			0,
		},
		{
			"Anastasi",
			0,
			0,
		},
		{
			"Kolm",
			0,
			0,
		},
		{
			"BNT",
			0,
			0,
		},
		{
			"noonCom",
			0,
			0,
		},
		{
			"Five",
			0,
			0,
		},
		{
			"SysT",
			0,
			0,
		},
		{
			"BridgeTM",
			0,
			0,
		},
		{
			"Apology",
			0,
			0,
		},
		{
			"ALOHA",
			0,
			0,
		},
		{
			"Int64",
			0,
			0,
		},
		{
			"Kirik_2020",
			0,
			0,
		},
		{
			"asgentos_lux",
			0,
			0,
		},
		{
			"Mr_Fin",
			0,
			0,
		},
	}
	//usersList := []string{
	//	"ADN",
	//	"Anastasi",
	//	"Kolm",
	//	"BNT",
	//	"noonCom",
	//	"Five",
	//	"SysT",
	//	"BridgeTM",
	//	"ALOHA",
	//	"Apology",
	//}

	//encoder := json.NewEncoder(ctx.Response())
	//decoder := json.NewDecoder(ctx.Request().Body)
	pinsStrings := validation.FindJpg.FindAllString(consts.PinterestPins, -1)
	PinUrls := []string{}
	for _, url := range pinsStrings {
		PinUrls = append(PinUrls, validation.FindPinUrl.FindString (url))
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < len(myUsers); i++ {
		var err error
		myUsers[i].UserID, err = h.PUsecase.AddNewUser(myUsers[i].Username, myUsers[i].Username + "@email.eu", "vr213b4t54k3fNFem3")
		if err != nil {
			return err
		}
		user, _ := h.PUsecase.GetUserByEmail(myUsers[i].Username + "@email.eu")
		newProfile := models.EditUserProfile{
			Name:     consts.LatinoNames[rand.Intn(len(consts.LatinoNames))],
			Surname:  consts.LatinoSurnames[rand.Intn(len(consts.LatinoSurnames))],
			Age:      string(16 + rand.Intn(30)),
			Status:   consts.LatinoDescription[rand.Intn(len(consts.LatinoDescription))],
		}


		_, err = h.PUsecase.SetUser(newProfile, user)


		board := models.Board{
			OwnerID:     myUsers[i].UserID,
			Title:       myUsers[i].Username + " board",
			Description: consts.LatinoDescription[rand.Intn(len(consts.LatinoDescription))],
			Category:    "default_category",
			CreatedTime: time.Now(),
		}
		myUsers[i].boardID, err = h.PUsecase.AddBoard(board)
		if err != nil {
			return err
		}
	}
	for i := 0; i < 20; i++ {
		for i := 0; i < len(myUsers); i++ {
			var err error

			index := rand.Intn(len(PinUrls))
			url := PinUrls[index]
			PinUrls = append(PinUrls[:index], PinUrls[index+1:]...)

			resp, err := http.Get(url)
			if err != nil {
				continue
			}

			var buf bytes.Buffer
			tee := io.TeeReader(resp.Body, &buf)
			fileHash, err := h.PUsecase.CalculateMD5FromFile(tee)
			if err != nil {
				return err
			}
			if err = h.PUsecase.AddDir("static/pin/" + fileHash[:2]); err != nil {
				return err
			}
			fileName := "static/pin/" + fileHash[:2] + "/" + fileHash + ".jpg"
			if err = h.PUsecase.AddPictureFile(fileName, &buf); err != nil {
				return err
			}

			pin := models.Pin{
				OwnerID:     myUsers[i].UserID,
				AuthorID:    myUsers[i].UserID,
				BoardID:     myUsers[i].boardID,
				Title:       consts.LatinoTitles[rand.Intn(len(consts.LatinoTitles))],
				Description: consts.LatinoDescription[rand.Intn(len(consts.LatinoDescription))],
				PinDir:      fileName,
				CreatedTime: time.Now(),
			}
			_, err = h.PUsecase.AddPin(pin)
			if err != nil {
				return err
			}

			resp.Body.Close()
		}
	}

	data := models.ValeraJSONResponse{ctx.Get("token").(string),nil}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
