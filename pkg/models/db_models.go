package models

import "database/sql"

type DBUser struct {
	ID        uint64
	Username  string
	Name      sql.NullString
	Surname   sql.NullString
	Password  string
	Email     string
	Age       sql.NullInt32
	Status    sql.NullString
	AvatarDir sql.NullString
	IsActive  bool
}
