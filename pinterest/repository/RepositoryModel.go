package repository

import "database/sql"

type RepositoryStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type RepositoryInterface interface {
	WriteData(executeQuery string, params []interface{}) (string, error)
	UniversalRead(executeQuery string, readSlice DBReader, params []interface{}) error
}
