package model

import (
    "golang.org/x/crypto/bcrypt"
    "os"
    "log"
)

// 納品情報の保存
type User struct {
	UserId string    `xorm:"varchar(36) pk" json:"userId"`
	Password string    `xorm:"varchar(60) not null" json:"password"`
}

func (User) TableName() string {
	return "users"
}

// idからユーザー情報を取得する
func GetUserById(id string) (User, error) {
    var user User
    has, err := db.Where("user_id = ?", id).Get(&user)
    if err != nil {
        return user, err
    }
    if !has{
        return user, err
    }
    return user,nil
}

// パスワードをハッシュ化して保存する関数
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// 入力されたパスワードとハッシュと一致するか検証する関数
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func RegisterUser() error{
    // ユーザー情報を登録する
    userId := os.Getenv("USER_ID")
    password := os.Getenv("USER_PASSWORD")
    user := User{
        UserId: userId,
        Password: password,
    }

    // パスワードをハッシュ化
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        log.Println("ユーザー情報の登録に失敗しました:", err)
        return err
    }
    user.Password = hashedPassword

    // ユーザー情報を登録
    _, err = db.Insert(&user)
    if err != nil {
        log.Println("ユーザー情報の登録に失敗しました:", err)
        return err
    }
    log.Println("ユーザー情報を登録しました")
    return err

}