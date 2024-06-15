package route

import (
	"auto_slip_api/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() (*gin.Engine, error) {

	router := gin.Default() // gin の初期化エンジンを作成

	
		v1 := router.Group("/v1")
		{

			// 雑誌
			magazine := v1.Group("/magazine")
			{
				magazine.GET("/", controller.CreateMagazines)    // 雑誌取得
				magazine.POST("/register", controller.CreateMagazines)   // 雑誌登録
				magazine.PUT("/", controller.CreateMagazines)    // 雑誌更新
				magazine.DELETE("/", controller.CreateMagazines) // 雑誌削除
			}

			// 顧客
			customer := v1.Group("/customer")
			{
				customer.GET("/customer", controller.CreateMagazines)
				customer.POST("/register", controller.RegisterCustomerHandler)
				customer.PUT("/update", controller.CreateMagazines)
				customer.DELETE("/delete", controller.CreateMagazines)
			}

			// 定期
			regular := v1.Group("/regular")
			{
				regular.GET("/", controller.CreateMagazines)
				regular.POST("/", controller.CreateMagazines)
				regular.PUT("/", controller.CreateMagazines)
				regular.DELETE("/", controller.CreateMagazines)
			}

			// 数取り
			counting := v1.Group("/counting")
			{
				counting.GET("/", controller.CreateMagazines)
				counting.POST("/", controller.CreateMagazines)
				counting.PUT("/", controller.CreateMagazines)
				counting.DELETE("/", controller.CreateMagazines)
			}

			// 請求
			claim := v1.Group("/claim")
			{
				claim.GET("/", controller.CreateMagazines)
				claim.POST("/", controller.CreateMagazines)
				claim.PUT("/", controller.CreateMagazines)
				claim.DELETE("/", controller.CreateMagazines)
			}

			// 納品
			delivery := v1.Group("/delivery")
			{
				delivery.GET("/", controller.CreateMagazines)
				delivery.POST("/", controller.CreateMagazines)
				delivery.PUT("/", controller.CreateMagazines)
				delivery.DELETE("/", controller.CreateMagazines)
			}

			csv := v1.Group("/csv")
			{
				csv.POST("/magazines", controller.CsvMagazinesRegister)
				csv.POST("/customers", controller.CreateMagazines)
				csv.POST("/regulars", controller.CreateMagazines)
			}

		}
	

	return router, nil // router設定されたengineを返す。

}
