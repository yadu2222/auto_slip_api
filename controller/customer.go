package controller

import (
	// "fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"


	"auto_slip_api/model"
	"auto_slip_api/service"

	"encoding/json"
	"auto_slip_api/csv"
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

// お客様削除
func DeleteCustomerHandler(c *gin.Context) {
	// パラメータからお客様IDを取得
	customerID := c.Param("customer_uuid")
	// 投げる
	customer ,err := customerService.DeleteCustomer(customerID); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500, 
			"error": "お客様情報の削除に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "お客様情報の削除に成功しました",
		"srvResData": customer,
	})
}

// 顧客一覧を取得
func GetCustomersHandler(c *gin.Context) {
	// 投げる
	customers, err := customerService.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "顧客情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "顧客情報の取得に成功しました",
		"srvResData": customers,
	})
}

// 顧客情報を名前で検索して取得
func GetCustomerByNameHandler(c *gin.Context) {
	// パラメータからお客様名を取得
	customerName := c.Param("customer_name")
	// 投げる
	customer, err := customerService.FindCustomerByName(customerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "顧客情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "顧客情報の取得に成功しました",
		"srvResData": customer,
	})
}

// csvからの登録
func CsvCustomersRegister(c *gin.Context) {
	// ファイルを受け取る
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "無効なデータです"})
		return
	}
	defer file.Close()
	// CSVファイルを読み込む
	records, err := csv.ProcessUniCSVFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの読み込みに失敗しました"})
		return
	}
	// CSVをJSONに変換
	byte, err := service.CsvToCustomerJSON(records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの変換に失敗しました"})
		return
	}
	// 変換したJSONを構造体にマッピング
	var csvUtilCustomers []model.Customer
	if err := json.Unmarshal(byte, &csvUtilCustomers); err != nil {
		print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}
	// 顧客情報を登録
	if err := customerService.RegisterCustomers(csvUtilCustomers); err != nil {	// なげる
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "顧客情報の登録に失敗しました"})
		return
	}
}
