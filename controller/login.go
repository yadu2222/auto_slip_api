package controller

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"auto_slip_api/model"

    "auto_slip_api/service"
)

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
//     username := r.FormValue("username")
//     password := r.FormValue("password")

//     token, err := service.Login(username, password)
//     if err != nil {
//         w.WriteHeader(http.StatusUnauthorized)
//         w.Write([]byte("Unauthorized"))
//         return
//     }

//     http.SetCookie(w, &http.Cookie{
//         Name:  "token",
//         Value: token,
//         Path:  "/",
//     })
//     w.Write([]byte("Login successful"))
// }

var loginService = service.LoginService{} // サービスの実体を作る。

func LoginHandler(c *gin.Context) {
	// マッピング
	var user model.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 400,
			"error":      "リクエストデータが無効です",
			"srvResData": gin.H{}})
		return
	}
	// 投げる
	token,err := loginService.Login(user); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "ログインに失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "ログインに成功しました",
		"srvResData": gin.H{
			"token": token,
		},
	})
}

// user作成
func CreateUserHandler(c *gin.Context) {
	
	// 投げる
	if err := loginService.CreateUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 500,
			"error":      "ユーザー情報の登録に失敗しました"})
		return
	}
	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 200,
		"srvResMsg":  "ユーザー情報の登録に成功しました",
		"srvResData": gin.H{},
	})
}
