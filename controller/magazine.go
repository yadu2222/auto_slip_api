package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"auto_slip_api/model"
	"auto_slip_api/service"

	"auto_slip_api/csv"
	"encoding/json"
)

var magazineService = service.MagazineService{} // サービスの実体を作る。

// 雑誌登録
func CreateMagazineHandler(c *gin.Context) {
	// マッピング
	var magazine model.Magazine
	if err := c.ShouldBindBodyWith(&magazine, binding.JSON); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "リクエストデータが無効です",
			"srvResData": gin.H{}})
		return
	}
	// 投げる
	if err := magazineService.RegisterMagazine(magazine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "お客様情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}

// 雑誌登録
func CreateMagazinesHandler(c *gin.Context) {
	// マッピング
	var magazines []model.Magazine
	if err := c.ShouldBindBodyWith(&magazines, binding.JSON); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "リクエストデータが無効です",
			"srvResData": gin.H{}})
		return
	}
	// 投げる
	if err := magazineService.RegisterMagazines(magazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "お客様情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}

// 雑誌を削除
func DeleteMagazineHandler(c *gin.Context) {
	// パラメータから雑誌コードを取得
	magazineCode := c.Param("magazine_code")

	// 雑誌コードから雑誌を削除
	// 投げる
	magazine, err := magazineService.DeleteMagazine(magazineCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の削除に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "雑誌情報の削除に成功しました",
		"srvResData": gin.H{
			"magazine": magazine,
		},
	})
}

// 雑誌一覧を取得
func GetMagazinesHandler(c *gin.Context) {
	// 投げる
	magazines, err := magazineService.GetMagazines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResData": magazines,
	})
}

// 雑誌コードで検索
func GetMagazineByCodeHandler(c *gin.Context) {
	// パラメータから雑誌コードを取得
	magazineCode := c.Param("magazine_code")
	// 投げる
	magazine, err := magazineService.FindMagazineByCode(magazineCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResData": magazine,
	})
}

// 雑誌名で検索
func GetMagazineByNameHandler(c *gin.Context) {
	// パラメータから雑誌名を取得
	magazineName := c.Param("magazine_name")
	// 投げる
	magazine, err := magazineService.FindMagazineByName(magazineName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の取得に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResData": magazine,
	})
}

// csvからの登録
func CsvMagazinesRegister(c *gin.Context) {
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
	byte, err := service.CsvToMagazineJson(records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの変換に失敗しました"})
		return
	}
	// 変換したJSONを構造体にマッピング
	var csvUtilMagazines []model.Magazine
	if err := json.Unmarshal(byte, &csvUtilMagazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}
	// 雑誌情報を登録
	if err := magazineService.RegisterMagazines(csvUtilMagazines); err != nil { // なげる
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の登録に失敗しました"})
		return
	}
}

// 雑誌情報を更新
func UpdateMagazineHandler(c *gin.Context) {
	// マッピング
	var magazine model.Magazine
	if err := c.ShouldBindBodyWith(&magazine, binding.JSON); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "リクエストデータが無効です",
			"srvResData": gin.H{}})
		return
	}

	// パラメータから雑誌コードを取得
	oldMagazineCode := c.Param("old_magazine_code")


	// 投げる
	if err := magazineService.UpdateMagazine(magazine,oldMagazineCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の更新に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "雑誌情報の更新に成功しました",
		"srvResData": gin.H{},
	})
}
