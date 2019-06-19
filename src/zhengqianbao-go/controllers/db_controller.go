package controllers

import (
	"database/sql"
	"fmt"
	"sync"

	"../global"
)

var once sync.Once

type DBRepository struct {
}

var dbRepository *DBRepository

func GetDBInstance() *DBRepository {
	once.Do(func() {
		dbRepository = NewDBRepository()
	})
	return dbRepository
}

// NewUserRepository returns a new user repository,
// the one and only repository type in our example.
func NewDBRepository() *DBRepository {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &DBRepository{}
}
