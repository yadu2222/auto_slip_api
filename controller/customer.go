package controller

import (
	// "fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"


	"auto_slip_api/model"
	"auto_slip_api/service"
)

var customerService = service.CustomerService{} // サービスの実体を作る。


// お客様登録
func RegisterCustomerHandler(c *gin.Context) {
	// マッピング
	var customer model.Customer
	if err := c.ShouldBindBodyWith(&customer,binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400, 
			"error": "リクエストデータが無効です",
			"srvResData": gin.H{},})
		return
	}
	// 投げる
	if err := customerService.RegisterCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500, 
			"error": "お客様情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "お客様情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}
