package controller

import (
	// "fmt"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/gin-gonic/gin/binding"
	// "github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"

	"fmt"

	"auto_slip_api/model"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

	// CSVリーダーを作成
	decoder := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())	// デコード
	reader := csv.NewReader(decoder)

	// ファイルを読み込む
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの読み込みに失敗しました"})
		return
	}
	fmt.Println("data:", records)

	// csvデータをバイトに変換
	byts, err := csvutil.Marshal(records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "csvの読み込みに失敗しました"})
		return
	}

	// csvデータをマッピング
	var csvUtilMagazines []model.Magazine
	if err := csvutil.Unmarshal(byts, &csvUtilMagazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "マッピングに失敗しました"})
		return
	}

	if err := magazineService.RegisterMagazines(csvUtilMagazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "雑誌情報の登録に失敗しました"})
		return
	}

}
