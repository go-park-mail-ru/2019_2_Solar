package models

import (
	"time"
)

type UserCookie struct {
	Value      string    `json:"-"`
	Expiration time.Time `json:"-"`
}

type UserSession struct {
	ID     uint64 `json:"-"`
	UserID uint64 `json:"-"`
	UserCookie
}

type UserReg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserUnique struct {
	ID       uint64
	Email    string
	Username string
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserProfile struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      string `json:"age"`
	Status   string `json:"status"`
}

type AnotherUser struct {
	ID         uint64 `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Password   string `json:"-"`
	Email      string `json:"-"`
	Age        uint   `json:"age"`
	Status     string `json:"status"`
	AvatarDir  string `json:"avatar_dir"`
	IsActive   bool   `json:"is_active"`
	IsFollowee bool   `json:"is_followee"`
}

type User struct {
	ID          uint64    `json:"-"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Password    string    `json:"-"`
	Email       string    `json:"email"`
	Age         uint      `json:"age"`
	Status      string    `json:"status"`
	AvatarDir   string    `json:"avatar_dir"`
	IsActive    bool      `json:"is_active"`
	Salt        string    `json:"-"`
	CreatedTime time.Time `json:"created_time"`
}

type DataJSON struct {
	UserJSON  interface{} `json:"user,omitempty"`
	UsersJSON interface{} `json:"users,omitempty"`
	InfoJSON  interface{} `json:"info,omitempty"`
}

type OutJSON struct {
	CSRFToken string      `json:"csrf_token"`
	BodyJSON  interface{} `json:"body"`
}

type NewBoard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type Board struct {
	ID          uint64    `json:"id"`
	OwnerID     uint64    `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedTime time.Time `json:"created_time"`
	IsDeleted   bool      `json:"is_deleted"`
}

type NewPin struct {
	BoardID     uint64 `json:"board_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PinDir      string `json:"pin_dir"`
}

type PinForMainPage struct {
	ID        uint64 `json:"id"`
	PinDir    string `json:"pin_dir"`
	Title     string `json:"title"`
	IsDeleted bool   `json:"is_deleted"`
}

type PinForSearchResult struct {
	ID     uint64 `json:"id"`
	PinDir string `json:"pin_dir"`
	Title  string `json:"title"`
}

type PinDisplay struct {
	ID     uint64 `json:"id"`
	PinDir string `json:"pin_dir"`
	Title  string `json:"title"`
}

type Pin struct {
	ID          uint64    `json:"id"`
	OwnerID     uint64    `json:"owner_id"`
	AuthorID    uint64    `json:"author_id"`
	BoardID     uint64    `json:"board_id"`
	PinDir      string    `json:"pin_dir"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedTime time.Time `json:"created_time"`
	IsDeleted   bool      `json:"is_deleted"`
}

type FullPin struct {
	ID             uint64    `json:"id"`
	OwnerUsername  string    `json:"owner_username"`
	AuthorUsername string    `json:"author_username"`
	BoardID        uint64    `json:"board_id"`
	PinDir         string    `json:"pin_dir"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CreatedTime    time.Time `json:"created_time"`
	IsDeleted      bool      `json:"is_deleted"`
}

type NewNotice struct {
	Message string `json:"message"`
}

type Notice struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	ReceiverID  uint64    `json:"receiver_id"`
	Message     string    `json:"message"`
	CreatedTime time.Time `json:"created_time"`
	IsRead      bool      `json:"is_read"`
}

type NewComment struct {
	Text string `json:"text"`
}

type Comment struct {
	ID          uint64    `json:"id"`
	PinID       uint64    `json:"pin_id"`
	Text        string    `json:"text"`
	CreatedTime time.Time `json:"created_time"`
	AuthorID    uint64    `json:"author_id"`
}

type CommentDisplay struct {
	Text          string    `json:"text"`
	CreatedTime   time.Time `json:"created_time"`
	Author        string    `json:"author_username"`
	AuthorPicture string    `json:"author_dir"`
}

type NewChatMessage struct {
	IdSender          uint64 `json:"id_sender"`
	UserNameRecipient string `json:"username_recipient"`
	Message           string `json:"text"`
}

type ChatMessage struct {
	IdSender    uint64    `json:"id_sender"`
	IdRecipient uint64    `json:"id_recipient"`
	Message     string    `json:"text"`
	SendTime    time.Time `json:"send_time"`
	IsDeleted   bool      `json:"is_deleted"`
}

type Subscribe struct {
	Id           uint64
	IdSubscriber uint64
	FolloweeId   uint64
}
