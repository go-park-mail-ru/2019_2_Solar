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
	Id       uint64
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

type User struct {
	ID        uint64 `json:"-"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Status    string `json:"status"`
	AvatarDir string `json:"avatar_dir"`
	IsActive  bool   `json:"is_active"`
}


type DataJSON struct {
	UserJSON  interface{} `json:"user,omitempty"`
	UsersJSON interface{} `json:"users,omitempty"`
	InfoJSON  interface{} `json:"info,omitempty"`
}

type OutJSON struct {
	BodyJSON interface{} `json:"body"`
}

type NewBoard struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
}

type Board struct {
	ID			uint64 `json:"id"`
	OwnerID 	uint64 `json:"owner_id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Category 	string `json:"category"`
	CreatedTime time.Time `json:"created_time"`
	IsDeleted 	bool `json:"is_deleted"`
}

type NewPin struct {
	BoardID		uint64 `json:"board_id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	PinDir 		string `json:"pin_dir"`
}

type Pin struct {
	ID			uint64 `json:"id"`
	OwnerID 	uint64 `json:"owner_id"`
	AuthorID	uint64 `json:"author_id"`
	BoardID		uint64 `json:"board_id"`
	PinDir		string `json:"pin_dir"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	CreatedTime time.Time `json:"created_time"`
	IsDeleted 	bool `json:"is_deleted"`
}
