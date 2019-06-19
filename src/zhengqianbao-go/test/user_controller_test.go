package test

import (
	"testing"

	"../controllers"
	"../models"
)

func TestInsert(t *testing.T) {
	userDB := controllers.GetDBInstance()
	user1 := models.User{Phone: "phone1", Remain: 0, Iscow: false, Name: "name1", Password: "password1", Gender: "Male",
		Age: 12, University: "SYSU", Company: "", Description: "", Class: ""}

	var ok = userDB.InsertUser(&user1)
	if ok == false {
		t.Errorf("Test Insert FAIL")
	}
}

func TestQuery(t *testing.T) {
	userDB := controllers.GetDBInstance()

	var ok = userDB.QueryUser("phone1")
	if ok == false {
		t.Errorf("Test Query FAIL in user1")
	}
	ok = userDB.QueryUser("phone2")
	if ok == true {
		t.Errorf("Test Query FAIL in user2")
	}
}

func TestSelect(t *testing.T) {
	userDB := controllers.GetDBInstance()

	user1, _ := userDB.SelectUser("phone1")

	if user1 == nil {
		t.Errorf("Test Query FAIL in user1")
	} else {
		t.Logf("Test Query SUCCESS in user1: %v", *user1)
	}
	user2, _ := userDB.SelectUser("phone2")
	if user2 != nil {
		t.Errorf("Test Query FAIL in user2")
	}
}
