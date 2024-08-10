package controller

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"

	"auto_slip_api/model"
	"auto_slip_api/service"

	"auto_slip_api/csv"
	"encoding/json"
)

var regularService = service.RegularService{} // サービスの実体を作る。

// 定期情報登録
func CreateRegularHandler(c *gin.Context) {
	// マッピング
	var regular model.Regular
	if err := c.ShouldBindJSON(&regular); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "リクエストデータが無効です",
			"srvResData": gin.H{},
		})
		return
	}
	// 投げる
	if err := regularService.RegisterRegular(regular); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "定期情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "定期情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}


// 雑誌を主キーに定期情報一覧取得
func GetMagazineRegularsHandler(c *gin.Context) {
	// 投げる
	regulars, err := regularService.FindMagazineRegulars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "定期情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "定期情報の取得に成功しました",
		"srvResData": regulars,
	})
}

// 顧客を主キーに定期情報一覧取得
func GetCustomerRegularsHandler(c *gin.Context) {

	// 投げる
	regulars, err := regularService.FindCustomerRegulars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "定期情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "定期情報の取得に成功しました",
		"srvResData": regulars,
	})
}

// csvからの登録
func CsvRegularRegister(c *gin.Context) {
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
	byte, err := service.CsvToRegularJSON(records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの変換に失敗しました"})
		return
	}
	// 変換したJSONを構造体にマッピング
	var csvUtilRegulars []model.Regular
	if err := json.Unmarshal(byte, &csvUtilRegulars); err != nil {
		print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}
	// 雑誌情報を登録
	if err := regularService.RegisterRegulars(csvUtilRegulars); err != nil { // なげる
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "定期情報の登録に失敗しました"})
		return
	}
}
