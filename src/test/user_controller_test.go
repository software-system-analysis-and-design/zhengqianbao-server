package test

import (
	"fmt"
	"testing"
	"time"

	"../controllers"
	"../models"
)

func TestInsert(t *testing.T) {
	userDB := controllers.GetUserDBInstance()

	user1 := models.User{Name: "user1", Password: "pwd1", Genre: "Male"}

	var ok = userDB.Insert(&user1)
	if ok == false {
		t.Errorf("Test Insert FAIL")
	}
}

func TestQuery(t *testing.T) {
	userDB := controllers.GetUserDBInstance()

	var ok = userDB.Query("user1")
	if ok == false {
		t.Errorf("Test Query FAIL in user1")
	}
	ok = userDB.Query("user2")
	if ok == true {
		t.Errorf("Test Query FAIL in user2")
	}
}

func TestSelect(t *testing.T) {
	userDB := controllers.GetUserDBInstance()

	user1, _ := userDB.Select("user1")

	fmt.Println(*user1)

	if user1 == nil {
		t.Errorf("Test Query FAIL in user1")
	} else {
		t.Logf("Test Query SUCCESS in user1: %v", *user1)
	}
	user2, _ := userDB.Select("user2")
	if user2 != nil {
		t.Errorf("Test Query FAIL in user2")
	}
}

func TestAny(t *testing.T) {
	t = time.Now().Add(time.Hour * 2).Unix()
	println(t)
}
