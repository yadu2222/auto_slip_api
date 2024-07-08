package model

// import (
// 	"time"
// )

// 取次データの構造体
type Agency struct {
	CountingUuid string  `xorm:"pk varchar(36)" json:"countingUUId"`		// 数取りID
	MagazineName string `xorm:"varchar(36) not null" json:"magazineName"`	// 雑誌名
	MagazineCode string `xorm:"varchar(36) not null" json:"magazineCode"`	// 雑誌コード
	Number string `json:"number"`	// 号数
	Quenity int `json:"quantity"`	// 冊数
}
