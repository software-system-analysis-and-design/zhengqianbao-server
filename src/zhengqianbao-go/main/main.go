package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// app.Use(cors.Default())

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods: []string{
			"GET", "OPTIONS", "POST",
			"PATCH", "PUT", "DELETE",
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	app.Use(crs)

	app.Use(CorsMiddleware)

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
	user.Get("/profile", jwtHandler.Serve, userTokenHandler, userProfileHandler)
	user.Options("/profile", jwtHandler.Serve, userTokenHandler, userProfileHandler)
	user.Post("/update", jwtHandler.Serve, userTokenHandler, userUpdateHandler)
	user.Options("/update", jwtHandler.Serve, userTokenHandler, userUpdateHandler)

	// questionnaire
	questionnaire := app.Party("/questionnaire")
	questionnaire.Get("/previews", QPreviewHandler)
	questionnaire.Post("/create", jwtHandler.Serve, QCreateHandler)
	questionnaire.Options("/create", jwtHandler.Serve, QCreateHandler)
	questionnaire.Post("/select", QSelectHandler)
	questionnaire.Post("/update", jwtHandler.Serve, QUpdateHandler)
	questionnaire.Options("/update", jwtHandler.Serve, QUpdateHandler)
	questionnaire.Post("/delete", jwtHandler.Serve, QDeleteHandler)
	questionnaire.Options("/delete", jwtHandler.Serve, QDeleteHandler)
	questionnaire.Post("/trash", jwtHandler.Serve, QTrashHandler)
	questionnaire.Options("/trash", jwtHandler.Serve, QTrashHandler)

	// update
	record := app.Party("/record")
	record.Post("/create", jwtHandler.Serve, RecordCreateHandler)
	record.Options("/create", jwtHandler.Serve, RecordCreateHandler)
	record.Post("/select", jwtHandler.Serve, RecordSelectHandler)
	record.Options("/select", jwtHandler.Serve, RecordSelectHandler)
	record.Post("/update", jwtHandler.Serve, RecordUpdateHandler)
	record.Options("/update", jwtHandler.Serve, RecordUpdateHandler)
	record.Post("/delete", jwtHandler.Serve, RecordDeleteHandler)
	record.Options("/delete", jwtHandler.Serve, RecordDeleteHandler)

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// iris.TLS(":8080", "1_littlefish33.cn_bundle.crt", "2_littlefish33.cn.key"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

func CorsMiddleware(ctx iris.Context) {
	method := ctx.Request().Method

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		ctx.WriteString("Options Request!")
		return
	}

	// 核心处理方式
	// ctx.Header("Access-Control-Allow-Origin", "*")
	// ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	// ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")

	ctx.Next()
}
