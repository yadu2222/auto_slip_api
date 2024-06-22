package controller

import (
	
	// "fmt"
	// "encoding/csv"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// "github.com/gin-gonic/gin/binding"
	// "github.com/gocarina/gocsv"
	// "github.com/jszwec/csvutil"

	"fmt"
	"auto_slip_api/model"
	"auto_slip_api/csv"

	// "golang.org/x/text/encoding/japanese"
	// "golang.org/x/text/encoding/unicode"
	// "golang.org/x/text/transform"
)

// var magazineService = service.MagazineService{} // サービスの実体を作る。

func CsvMagazinesRegister(c *gin.Context) {
	// ファイルを受け取る
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	records, err := csv.ProcessUniCSVFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの読み込みに失敗しました"})
		return
	}
	// データの変換と格納
	var convertedList []map[string]string // 変換後のリストを格納するスライス
	for _, record := range records {
		// UUIDを生成して追加
		uid, _ := uuid.NewRandom()
		// マップを作成してリストに追加
		magazineMap := map[string]string{
			"magazineUUId": uid.String(),
			"magazineCode": record[0],
			"magazineName": record[1],
			"takerUUID":"c99cb6c4-42b9-4d6b-9884-ae6664f9df00",	// とりあえず自分のid	
		}
		convertedList = append(convertedList, magazineMap)
	}

	// マップをJSONに変換
	byte, err := json.Marshal(convertedList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byte))

	// 変換したJSONを構造体にマッピング
	var csvUtilMagazines []model.Magazine
	if err := json.Unmarshal(byte, &csvUtilMagazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}
	// 雑誌情報を登録
	if err := magazineService.RegisterMagazines(csvUtilMagazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の登録に失敗しました"})
		return
	}
}
