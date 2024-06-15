package controller

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"auto_slip_api/model"
	"auto_slip_api/service"
)

// 雑誌登録
func CreateMagazine(c *gin.Context) {
	// リクエストボディを読み取ってログに出力
	// body, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// エラー処理
	// 	fmt.Println("Failed to read request body:", err)
	// }
	// fmt.Println("Request body:", string(body))

	var magazines []model.Magazine
	if err := c.ShouldBindBodyWith(&magazines,binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7006, 
			"error": "リクエストデータが無効です"})
		return
	}

	// // // 乱数生成器のシードを設定する（一般的には現在時刻を使う）
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// // // 5桁の乱数を生成し、0で埋める
	// randomInt := rand.Intn(100000) // 0から99999までの乱数を生成
	// randomString := fmt.Sprintf("%05d", randomInt)
	// group.GroupJoinId = randomString // 乱数を代入

	// fmt.Println(group)

	if err := service.CreateMagazines(magazines); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7005, 
			"error": "グループの作成に失敗しました"})
		return
	}

	// fmt.Println("ロールいくぜ")
	// CreateRolls(c,randomString);
	// fmt.Println("おわった")
	// var rolls []model.Roll


	// 一旦格納
	type RequestBody struct {
		GroupName string   `json:"groupName"`
		Rolls     []string `json:"rolls"`
	}
	var rolls RequestBody
	// c.Bindでbodyを[]model.Rollに結びつける
	err := c.ShouldBindBodyWith(&rolls,binding.JSON)
	if err != nil {
		fmt.Println("ここまではきた")
		fmt.Println(rolls)
		// エラーが出た際の処理
		c.String(http.StatusBadRequest, "Invalid JSON")
	}

	// []model.Roll を []*model.Roll に変換する
	// for文で回す
	// var ptrRolls []*model.Roll
	// for i := range rolls.Rolls {
	// 	fmt.Println("ここまでもきた")

	// 	roll := model.Roll{
	// 		GroupId:  randomString, // groupIdをセット
	// 		RollName: rolls.Rolls[i],     // rollNameをセット
	// 	}
	// 	ptrRolls = append(ptrRolls, &roll)
	// }

	// // サービスの関数でポインタのスライスを渡す
	// err = service.CreateRoll(ptrRolls)
	// if err != nil {
	// 	// エラーが出た際の処理
	// 	c.String(http.StatusInternalServerError, "Internal Server Error")
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "グループが作成されました",
	// 	"group":   group,
	// })
}
