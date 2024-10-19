package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"auto_slip_api/model"
)

func GenerateToken(userId string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 下記追加部分
func ParseToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}
	// トークンを検証
	// ユーザが存在するか
	isUser,err := model.GetUserById(token.Claims.(jwt.MapClaims)["user_id"].(string))
	if isUser.UserId == "" {
		return false, err
	}
	// トークンの有効期限が切れていないか
	if time.Unix(int64(token.Claims.(jwt.MapClaims)["exp"].(float64)), 0).Before(time.Now()) {
		return false, err
	}
	return true, nil
}
