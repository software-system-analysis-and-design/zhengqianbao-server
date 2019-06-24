package controllers

import (
	"database/sql"
	"fmt"

	"../global"
	"../models"
	_ "github.com/lib/pq"
)

type Query func(models.User) bool

type User_Interface interface {
	QueryUser(phone string) (ok bool)

	SelectUser(phone string) (user *models.User, ok bool)

	InsertUser(user *models.User) (ok bool)

	UpdateUser(phone string, user *models.User) (updatedUser *models.User, ok bool)

	UpdateMoney(phone string, money int) (ok bool)

	DeleteUser(phone string) (ok bool)
}

func (r *DBRepository) QueryUser(phone string) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	var count int64
	err = db.QueryRow("select count(*) from t_user where phone=$1", phone).Scan(&count)

	ok = true
	if err != nil || count == 0 {
		ok = false
	}

	db.Close()
	return ok
}

func (r *DBRepository) SelectUser(phone string) (user *models.User, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query("select * from t_user where phone=$1", phone)

	userObj := models.User{}
	rows.Next()
	err = rows.Scan(&userObj.Phone, &userObj.Remain, &userObj.Iscow, &userObj.Name, &userObj.Password, &userObj.Gender,
		&userObj.Age, &userObj.University, &userObj.Company, &userObj.Description, &userObj.Class)

	if err != nil {
		fmt.Printf("could not find user, %v", err)
		db.Close()
		return nil, false
	}

	db.Close()
	return &userObj, ok
}

func (r *DBRepository) InsertUser(userObj *models.User) (ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("insert into t_user (phone, remain, iscow, name, password, gender, age, university, company, description, class)" +
		" values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		ok = false
	}
	_, err = stmt.Exec(userObj.Phone, 0, userObj.Iscow, userObj.Name, userObj.Password, userObj.Gender,
		userObj.Age, userObj.University, userObj.Company, userObj.Description, userObj.Class)
	if err != nil {
		fmt.Printf("could not insert user, %v", err)
		ok = false
	}

	db.Close()
	return ok
}

func (r *DBRepository) UpdateUser(phone string, user *models.User) (updatedUser *models.User, ok bool) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_user set iscow=$1, name=$2, password=$3, gender=$4, age=$5, " +
		"university=$6, company=$7, description=$8, class=$9 WHERE phone=$10")
	if err != nil {
		ok = false
	}
	_, err = stmt.Exec(user.Iscow, user.Name, user.Password, user.Gender,
		user.Age, user.University, user.Company, user.Description, user.Class, user.Phone)
	if err != nil {
		ok = false
	}
	db.Close()

	return user, ok
}

func (r *DBRepository) UpdateMoney(phone string, money int) (ok bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ok = true

	stmt, err := db.Prepare("update t_user set remain=$1 WHERE phone=$2")
	if err != nil {
		ok = false
	}
	_, err = stmt.Exec(money, phone)
	if err != nil {
		ok = false
	}
	db.Close()

	return ok
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
