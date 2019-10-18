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
	IsActive string `json:"isactive"`
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
	AvatarDir string `json:"-"`
	IsActive  bool   `json:"isactive"`
}

type DataJSON struct {
	UserJSON  interface{} `json:"user,omitempty"`
	UsersJSON interface{} `json:"users,omitempty"`
	InfoJSON  interface{} `json:"info,omitempty"`
}

type OutJSON struct {
	BodyJSON interface{} `json:"body"`
}
