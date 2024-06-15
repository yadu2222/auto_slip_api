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
			magazine := v1.Group("/magazines")
			{
				magazine.GET("/", controller.CreateMagazine)    // 雑誌取得
				magazine.POST("/", controller.CreateMagazine)   // 雑誌登録
				magazine.PUT("/", controller.CreateMagazine)    // 雑誌更新
				magazine.DELETE("/", controller.CreateMagazine) // 雑誌削除
			}

			// 顧客
			customer := v1.Group("/customer")
			{
				customer.GET("/", controller.CreateMagazine)
				customer.POST("/", controller.CreateMagazine)
				customer.PUT("/", controller.CreateMagazine)
				customer.DELETE("/", controller.CreateMagazine)
			}

			// 定期
			regular := v1.Group("/regular")
			{
				regular.GET("/", controller.CreateMagazine)
				regular.POST("/", controller.CreateMagazine)
				regular.PUT("/", controller.CreateMagazine)
				regular.DELETE("/", controller.CreateMagazine)
			}

			// 数取り
			counting := v1.Group("/counting")
			{
				counting.GET("/", controller.CreateMagazine)
				counting.POST("/", controller.CreateMagazine)
				counting.PUT("/", controller.CreateMagazine)
				counting.DELETE("/", controller.CreateMagazine)
			}

			// 請求
			claim := v1.Group("/claim")
			{
				claim.GET("/", controller.CreateMagazine)
				claim.POST("/", controller.CreateMagazine)
				claim.PUT("/", controller.CreateMagazine)
				claim.DELETE("/", controller.CreateMagazine)
			}

			// 納品
			delivery := v1.Group("/delivery")
			{
				delivery.GET("/", controller.CreateMagazine)
				delivery.POST("/", controller.CreateMagazine)
				delivery.PUT("/", controller.CreateMagazine)
				delivery.DELETE("/", controller.CreateMagazine)
			}

		}
	

	return router, nil // router設定されたengineを返す。

}
