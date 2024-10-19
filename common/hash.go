package common

import (
    "golang.org/x/crypto/bcrypt"
)

// パスワードをハッシュ化して保存する関数
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}