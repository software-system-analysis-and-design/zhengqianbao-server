package main

import (
	"strconv"
	"time"

	"./controllers"
	"./helper"
	"./models"
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte("My Secret"), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})

	// user login and register module
	user := app.Party("/user")
	user.Post("/register", userRegisterHandler)
	// user.Post("/login", jwtHandler.Serve, userLoginHandler)
	user.Post("/login", userLoginHandler)
	user.Get("/logout", userLogoutHandler)
	user.Get("/profile", jwtHandler.Serve, tokenHandler, userProfileHandler)
	user.Post("/update", jwtHandler.Serve, tokenHandler, userUpdateHandler)

	// test := app.Party("/test")
	// test.Post("/token", jwtHandler.Serve, testTokenHandler)

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

func userRegisterHandler(ctx iris.Context) {
	iscow, _ := strconv.ParseBool(ctx.FormValue("iscow"))
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")
	gender := ctx.FormValue("gender")
	age, _ := strconv.Atoi(ctx.FormValue("age"))
	university := ctx.FormValue("university")
	company := ctx.FormValue("company")
	description := ctx.FormValue("description")
	class, _ := strconv.Atoi(ctx.FormValue("class"))
	phone := ctx.FormValue("phone")

	userDB := controllers.GetUserDBInstance()
	var found = userDB.Query(phone)
	if found == true {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "Phone number already registered!",
		}
		ctx.JSON(response)
		return
	}

	registerUser := models.User{Phone: phone, Iscow: iscow, Name: name, Password: password, Gender: gender,
		Age: age, University: university, Company: company, Description: description, Class: class}

	var ok = userDB.Insert(&registerUser)
	if ok == true {
		response := helper.Register_Response{
			Code: 200,
			Msg:  "Register successfully!",
		}
		ctx.JSON(response)
	}

}

func userLoginHandler(ctx iris.Context) {
	phone := ctx.FormValue("phone")
	password := ctx.FormValue("password")

	userDB := controllers.GetUserDBInstance()
	var found = userDB.Query(phone)
	if found == false {
		response := helper.Login_Response{
			Code: 400,
			Msg:  "User dose not exist!",
		}
		ctx.JSON(response)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":    phone,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 添加过期时间
	})

	tokenString, _ := token.SignedString([]byte("My Secret"))
	tokenResponse := helper.Token_Response{
		Code: 200,
		Msg:  "Login successfully!",
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
	userDB := controllers.GetUserDBInstance()
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
	class, _ := strconv.Atoi(ctx.FormValue("class"))
	phone := ctx.FormValue("phone")

	if tokenPhone != phone {
		response := helper.Register_Response{
			Code: 400,
			Msg:  "Phone number is wrong!",
		}
		ctx.JSON(response)
		return
	}

	updatedUser := models.User{Phone: phone, Iscow: iscow, Name: name, Password: password, Gender: gender,
		Age: age, University: university, Company: company, Description: description, Class: class}

	var ok = userDB.Insert(&updatedUser)
	if ok == true {
		response := helper.Register_Response{
			Code: 200,
			Msg:  "Update successfully!",
		}
		ctx.JSON(response)
	}
}

func userProfileHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	phone := userMsg["phone"].(string)
	userDB := controllers.GetUserDBInstance()
	user, _ := userDB.Select(phone)
	ctx.JSON(user)
}

func tokenHandler(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userDB := controllers.GetUserDBInstance()
	user, _ := userDB.Select(userMsg["phone"].(string))

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
