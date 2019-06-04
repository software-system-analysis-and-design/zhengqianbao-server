package controllers

import (
	"encoding/json"
	"fmt"
	"sync"

	"../models"
	"github.com/boltdb/bolt"
)

type Query func(models.User) bool

type User_Interface interface {
	Query(phone string) (ok bool)

	Select(phone string) (user *models.User, ok bool)

	Insert(user *models.User) (ok bool)

	Update(phone string, user *models.User) (updatedUser *models.User, ok bool)

	Delete(phone string) (ok bool)
}

type UserDBRepository struct {
	bucketName string
}

var userDBRepository *UserDBRepository
var once sync.Once

func GetUserDBInstance() *UserDBRepository {
	once.Do(func() {
		userDBRepository = NewDBRepository()
	})
	return userDBRepository
}

// NewUserRepository returns a new user repository,
// the one and only repository type in our example.
func NewDBRepository() *UserDBRepository {
	db, err := bolt.Open("db", 0600, nil)
	bucketName := "userdb"
	if err != nil {
		fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
	if err != nil {
		fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println(bucketName + " Setup Done")

	db.Close()

	return &UserDBRepository{bucketName}
}

func (r *UserDBRepository) Query(phone string) (ok bool) {
	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
		fmt.Errorf("could not open db, %v", err)
		db.Close()
		return false
	}
	ok = false
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(r.bucketName))
		v := b.Get([]byte(phone))
		if v != nil {
			ok = true
		}
		db.Close()
		return nil
	})
	return ok
}

func (r *UserDBRepository) Select(phone string) (user *models.User, ok bool) {
	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
		fmt.Errorf("could not open db, %v", err)
		db.Close()
		return nil, false
	}
	ok = false
	var v []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(r.bucketName))
		v = b.Get([]byte(phone))
		return nil
	})

	if v == nil {
		db.Close()
		return nil, ok
	}
	userObj := models.User{}
	json.Unmarshal([]byte(v), &userObj) //json解析到结构体里面
	// fmt.Println(userObj)                //输入结构体

	db.Close()
	return &userObj, ok
}

func (r *UserDBRepository) Insert(user *models.User) (ok bool) {

	// 转换成JSON返回的是byte[]
	jsonStr, errs := json.Marshal(*user)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	// fmt.Println(string(jsonStr))

	userBytes := []byte(jsonStr)
	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
		fmt.Errorf("could not open db, %v", err)
		db.Close()
		return false
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(r.bucketName)).Put([]byte(user.Phone), []byte(userBytes))
		if err != nil {
			return fmt.Errorf("could not insert %s: %v", r.bucketName, err)
		}
		return nil
	})
	if err != nil {
		db.Close()
		return false
	}
	db.Close()
	return true
}

func (r *UserDBRepository) Update(phone string, user *models.User) (updatedUser *models.User, ok bool) {
	// 转换成JSON返回的是byte[]
	jsonStr, errs := json.Marshal(*user)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	userBytes := []byte(jsonStr)

	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
		fmt.Errorf("could not open db, %v", err)
		db.Close()
		return nil, false
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(r.bucketName)).Put([]byte(user.Phone), []byte(userBytes))
		if err != nil {
			return fmt.Errorf("could not insert %s: %v", r.bucketName, err)
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, false
	}
	db.Close()

	return user, true
}
