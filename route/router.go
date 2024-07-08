package route

import (
	"auto_slip_api/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() (*gin.Engine, error) {

	router := gin.Default() // gin の初期化エンジンを作成

	
		v1 := router.Group("/v1")
		{
			// リクエストを鯖側で確かめるテスト用エンドポイント
			v1.GET("/test/cfmreq", controller.CfmReq) // /v1/test

			// 雑誌
			magazine := v1.Group("/magazine")
			{
				magazine.GET("/", controller.CreateMagazinesHandler)    // 雑誌取得
				magazine.POST("/register", controller.CreateMagazinesHandler)   // 雑誌登録
				magazine.PUT("/", controller.CreateMagazinesHandler)    // 雑誌更新
				magazine.DELETE("/", controller.CreateMagazinesHandler) // 雑誌削除
			}

			// 顧客
			customer := v1.Group("/customer")
			{
				customer.GET("/customer", controller.CreateMagazinesHandler)
				customer.POST("/register", controller.RegisterCustomerHandler)
				customer.PUT("/update", controller.CreateMagazinesHandler)
				customer.DELETE("/delete", controller.CreateMagazinesHandler)
			}

			// 定期
			regular := v1.Group("/regular")
			{
				regular.GET("/", controller.CreateMagazinesHandler)
				regular.POST("/", controller.CreateMagazinesHandler)
				regular.PUT("/", controller.CreateMagazinesHandler)
				regular.DELETE("/", controller.CreateMagazinesHandler)
			}

			// 数取り
			counting := v1.Group("/counting")
			{
				counting.GET("/", controller.CreateMagazinesHandler)
				counting.POST("/show", controller.CreateMagazinesHandler)
				counting.PUT("/", controller.CreateMagazinesHandler)
				counting.DELETE("/", controller.CreateMagazinesHandler)
			}

			// 請求
			claim := v1.Group("/claim")
			{
				claim.GET("/", controller.CreateMagazinesHandler)
				claim.POST("/", controller.CreateMagazinesHandler)
				claim.PUT("/", controller.CreateMagazinesHandler)
				claim.DELETE("/", controller.CreateMagazinesHandler)
			}

			// 納品
			delivery := v1.Group("/delivery")
			{
				delivery.GET("/", controller.CreateMagazinesHandler)
				delivery.POST("/", controller.CreateMagazinesHandler)
				delivery.PUT("/", controller.CreateMagazinesHandler)
				delivery.DELETE("/", controller.CreateMagazinesHandler)
			}

			// csvファイルを受け取ってDBに登録
			// 使うの最初だけ
			csv := v1.Group("/csv")
			{
				csv.POST("/magazines", controller.CsvMagazinesRegister)
				csv.POST("/customers", controller.CsvCustomersRegister)
				csv.POST("/regulars", controller.CsvRegularRegister)
				csv.POST("/counting", controller.CSVCountingHandler)
			}
		}
	

	return router, nil // router設定されたengineを返す。

}
