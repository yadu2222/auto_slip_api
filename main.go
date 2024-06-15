package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"auto_slip_api/middleware"
	"auto_slip_api/route"
	"auto_slip_api/service"
)

func main() {

	// db初期化
	service.Init()
	// service.DBInit()

	router := gin.Default()                // ginの初期化
	router.Use(cors.Default())             // corsを有効にする
	router.Use(middleware.RecordUaAndTime) // middleware.RecordUaAndTimeを適用する

	// router設定されたengineを受け取る。
	router, err := route.GetRouter()
	if err != nil {
		fmt.Println("Couldnt receive router engine.", err) // エラー内容を出力し早期リターン
		return
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello お元気？")
	})

	router.Run() // サーバーを起動する
}
