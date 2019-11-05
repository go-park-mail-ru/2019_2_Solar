package models

import (
	"database/sql"
	"time"
)

type DBUser struct {
	ID          uint64
	Username    string
	Name        sql.NullString
	Surname     sql.NullString
	Password    string
	Email       string
	Age         sql.NullInt32
	Status      sql.NullString
	AvatarDir   sql.NullString
	IsActive    bool
	Salt        string
	CreatedTime time.Time
}
