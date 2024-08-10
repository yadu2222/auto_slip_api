package controller

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"auto_slip_api/model"
	"auto_slip_api/service"

	"auto_slip_api/csv"
	"encoding/json"
)

var countingService = service.CountingService{} // サービスの実体を作る。

// 数取り
func CSVCountingHandler(c *gin.Context) {

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
	records, err := csv.ProcessJapCSVFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの読み込みに失敗しました"})
		return
	}
	// CSVをJSONに変換
	byte, err := service.CsvToAgencyJson(records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの変換に失敗しました"})
		return
	}
	// 変換したJSONを構造体にマッピング
	var csvUtilCustomers []model.Agency
	if err := json.Unmarshal(byte, &csvUtilCustomers); err != nil {
		print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}
	// 投げる
	countingList, err := countingService.MagazineCounting(csvUtilCustomers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "数取りに失敗しました",
		})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "数取りに成功しました",
		"srvResData": countingList,
	})
}
