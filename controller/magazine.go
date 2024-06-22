package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"auto_slip_api/model"
	"auto_slip_api/service"
)

var magazineService = service.MagazineService{} // サービスの実体を作る。

// 雑誌登録
func CreateMagazinesHandler(c *gin.Context) {
	// マッピング
	var magazines []model.Magazine
	if err := c.ShouldBindBodyWith(&magazines,binding.JSON); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400, 
			"error": "リクエストデータが無効です",
			"srvResData": gin.H{},})
		return
	}
	// 投げる
	if err := magazineService.RegisterMagazines(magazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500, 
			"error": "雑誌情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "お客様情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}
