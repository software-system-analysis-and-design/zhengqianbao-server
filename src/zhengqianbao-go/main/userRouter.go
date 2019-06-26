package main

import (
	"strconv"
	"time"

	"../controllers"
	"../helper"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func userRegisterHandler(ctx iris.Context) {
	phone := ctx.FormValue("phone")
	iscow, _ := strconv.ParseBool(ctx.FormValue("iscow"))
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")
	gender := ctx.FormValue("gender")
	age, _ := strconv.Atoi(ctx.FormValue("age"))
	university := ctx.FormValue("university")
	company := ctx.FormValue("company")
	description := ctx.FormValue("description")
	class := ctx.FormValue("class")
	remain := 0

	dbInstance := controllers.GetDBInstance()
	var found = dbInstance.QueryUser(phone)
	if found == true {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "该手机号已注册！",
		}
		ctx.JSON(response)
		return
	}

	registerUser := models.User{Phone: phone, Remain: remain, Iscow: iscow, Name: name, Password: password, Gender: gender,
		Age: age, University: university, Company: company, Description: description, Class: class}

	var ok = dbInstance.InsertUser(&registerUser)
	if ok == true {
		response := helper.Register_Response{
			Code: 200,
			Msg:  "注册成功！",
		}
		ctx.JSON(response)
	}

}

func userLoginHandler(ctx iris.Context) {
	phone := ctx.FormValue("phone")
	password := ctx.FormValue("password")

	dbInstance := controllers.GetDBInstance()
	var found = dbInstance.QueryUser(phone)
	if found == false {
		response := helper.Login_Response{
			Code: 400,
			Msg:  "用户不存在！",
		}
		ctx.JSON(response)
		return
	}

	user, _ := dbInstance.SelectUser(phone)
	if user.Password != password {
		response := helper.Login_Response{
			Code: 401,
			Msg:  "密码错误！",
		}
		ctx.JSON(response)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":    phone,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 2000).Unix(), // 添加过期时间
	})

	tokenString, _ := token.SignedString([]byte("My Secret"))
	tokenResponse := helper.Token_Response{
		Code: 200,
		Msg:  "登录成功！",
		Data: map[string]string{"token": tokenString}}
	ctx.JSON(tokenResponse)
}

func userLogoutHandler(ctx iris.Context) {
	response := helper.Login_Response{
		Code: 200,
		Msg:  "User successfully log out!",
	}
	ctx.JSON(response)
}

func userUpdateHandler(ctx iris.Context) {
	dbInstance := controllers.GetDBInstance()
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	tokenPhone := userMsg["phone"].(string)

	iscow, _ := strconv.ParseBool(ctx.FormValue("iscow"))
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")
	gender := ctx.FormValue("gender")
	age, _ := strconv.Atoi(ctx.FormValue("age"))
	university := ctx.FormValue("university")
	company := ctx.FormValue("company")
	description := ctx.FormValue("description")
	class := ctx.FormValue("class")
	phone := ctx.FormValue("phone")

	if tokenPhone != phone {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "token错误！",
		}
		ctx.JSON(response)
		return
	}

	updatedUser := models.User{Phone: phone, Iscow: iscow, Name: name, Password: password, Gender: gender,
		Age: age, University: university, Company: company, Description: description, Class: class}

	var _, ok = dbInstance.UpdateUser(phone, &updatedUser)
	if ok == true {
		response := helper.Register_Response{
			Code: 200,
			Msg:  "更新成功！",
		}
		ctx.JSON(response)
	} else {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "更新失败！",
		}
		ctx.JSON(response)
	}

}

func moneyUpdateHandler(ctx iris.Context) {
	dbInstance := controllers.GetDBInstance()
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	tokenPhone := userMsg["phone"].(string)

	money, _ := strconv.Atoi(ctx.FormValue("money"))

	ok := dbInstance.UpdateMoney(tokenPhone, money)
	if ok == true {
		response := helper.Register_Response{
			Code: 200,
			Msg:  "更新成功！",
		}
		ctx.JSON(response)
	} else {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "更新失败！",
		}
		ctx.JSON(response)
	}
}

func userProfileHandler(ctx iris.Context) {

	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)
	dbInstance := controllers.GetDBInstance()
	user, _ := dbInstance.SelectUser(phone)
	ctx.JSON(user)
}

func userTokenHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	dbInstance := controllers.GetDBInstance()
	user, _ := dbInstance.SelectUser(userMsg["phone"].(string))

	var response helper.Gene_Response
	response.Code = 200
	if user == nil {
		response.Code = 400
		response.Msg = "User does not exists!"
	} else if user.Password != userMsg["password"].(string) {
		response.Code = 400
		response.Msg = "Wrong password!"
	} else if userMsg["exp"].(float64) < (float64)(time.Now().Unix()) {
		response.Code = 400
		response.Msg = "Time expend!"
	}

	if response.Code == 400 {
		ctx.JSON(response)
		return
	} else {
		ctx.Next()
	}

}
