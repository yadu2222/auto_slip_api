package model

// import (
// 	"time"
// )

// 実質Viewみたいな構造体
// どんびきです
type RegularAgency struct {
	// 雑誌情報
	// これってあっちがもってるのでは？
	// MagazineCode string `json:"magazineCode"`	// 雑誌コード
	// MagazineName string `json:"magazineName"`	// 雑誌名
	// Number string `json:"number"`		// 号数
	// TotalQuantity int `json:"totalQuantity"`	// 冊数
	// 定期情報
	RegularUuid string `xorm:"varchar(36) not null" json:"RegularUUId"`
	// 顧客情報
	CustomerUuid string `xorm:"varchar(36) not null" json:"customerUUID"`	// 顧客コード
	CustomerName string `json:"customerName"`	// 顧客名
	Quantity int `json:"quantity"`		// 冊数
	MethodType int `json:"methodType"`		// 伝票or店取りというはなし
	
}

// magazine_codeをもとに定期情報を取得
func FindCountingMagazine(magazineCode string) ([]RegularAgency, error) {
	//クソデカ構造体のスライスを定義
	var regularAgencys []RegularAgency

	//クソデカ構造体をとるすごいやつだよ
	err := db.Table("regulars").
		Join("LEFT", "magazines", "magazines.magazine_code = regulars.magazine_code").
		Where("magazine_code = (?)", magazineCode).
		Join("LEFT", "customers", "customers.customer_uuid = regulars_uuid").
		Join("LEFT", "method_types", "method_types.method_id = customers.method_id").
		Select("regulars.regular_uuid, customers.customer_uuid, customers.customer_name, regulars.quantity, method_types.method_type").
		Find(&regularAgencys)
	if err != nil { //エラーハンドル ただエラー投げてるだけ
		return nil, err
	}

	//クソデカ構造体のスライスを返す
	return regularAgencys, nil
}