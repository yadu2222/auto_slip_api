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
		magazine := v1.Group("/magazines")
		{
			magazine.GET("/magazines", controller.GetMagazinesHandler)    // 雑誌取得
			magazine.GET("/magazines/:magazine_name", controller.GetMagazineByNameHandler) // 雑誌名で検索
			magazine.GET("/magazine/:magazine_code", controller.GetMagazineByCodeHandler) // 雑誌コードで検索
			magazine.POST("/register", controller.CreateMagazinesHandler) // 雑誌登録
			magazine.PUT("/", controller.CreateMagazinesHandler)          // 雑誌更新
			magazine.DELETE("/delete/:magazine_code", controller.DeleteMagazineHandler)       // 雑誌削除
		}

		// 顧客
		customer := v1.Group("/customers")
		{
			customer.GET("/customers", controller.GetCustomersHandler)
			customer.GET("/customers/:customer_name", controller.GetCustomerByNameHandler)
			customer.POST("/register", controller.RegisterCustomerHandler)
			customer.PUT("/update", controller.CreateMagazinesHandler)
			customer.DELETE("/delete/:customer_uuid", controller.DeleteCustomerHandler)
		}

		// 定期
		regular := v1.Group("/regulars")
		{
			regular.GET("/regulars", controller.GetMagazineRegularsHandler)	// 雑誌を主キーに定期情報一覧取得
			regular.GET("regulars/customer/:customer_name",controller.GetRegularsByCustomerNameHandler)	// 顧客名で検索
			regular.GET("regulars/magazine/name/:magazine_name",controller.GetRegularsByMagazineNameHandler)	// 雑誌名で検索
			regular.GET("regulars/magazine/code/:magazine_code",controller.GetRegularsByMagazineCodeHandler)	// 雑誌コードで検索
			regular.GET("/", controller.CreateMagazinesHandler)
			regular.POST("/register", controller.CreateRegularHandler)
			regular.PUT("/", controller.CreateMagazinesHandler)
			regular.DELETE("/delete/:regular_uuid", controller.DeleteRegularHandler)	// 定期削除
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
