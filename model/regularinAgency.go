package model

import (
	"fmt"
)

type RegularAgency struct {
	RegularUuid  string `xorm:"varchar(36) not null" json:"regularUUID"`
	CustomerUuid string `xorm:"varchar(36) not null" json:"customerUUID"`
	CustomerName string `json:"customerName"`
	Quantity     int    `json:"quantity"`
	MethodType   int    `json:"methodType"`
}

// magazine_codeをもとに定期情報を取得
// func FindCountingMagazine(magazineCode string) ([]RegularAgency, error) {
// 	var regularAgencys []RegularAgency

// 	err := db.Table("regulars").
// 		Where("regulars.magazine_code = ?", magazineCode).
// 		Join("LEFT", "magazines", "magazines.magazine_code = regulars.magazine_code").
// 		Join("LEFT", "customers", "customers.customer_uuid = regulars.customer_uuid").
// 		Select("regulars.regular_uuid, customers.customer_uuid, customers.customer_name, regulars.quantity, customers.method_type").
// 		Find(&regularAgencys)
// 	if err != nil {
// 		fmt.Print(err)
// 		return nil, err
// 	}

// 	return regularAgencys, nil
// }
func FindCountingMagazine(magazineCode string) ([]RegularAgency, error) {
	var regularAgencys []RegularAgency

	query := db.Table("regulars").
		Join("LEFT", "magazines", "magazines.magazine_code = regulars.magazine_code").
		Join("LEFT", "customers", "customers.customer_uuid = regulars.customer_uuid").
		Select("regulars.regular_uuid, customers.customer_uuid, customers.customer_name, regulars.quantity, customers.method_type")

	// magazine_codeが60000以上の場合、上4桁で検索
	if len(magazineCode) >= 5 && magazineCode >= "20000" {
		likePattern := magazineCode[:4] + "%"
		query = query.Where("substring(regulars.magazine_code,1,4) LIKE ?", likePattern)
	} else {
		// それ以外の場合は完全一致で検索
		query = query.Where("regulars.magazine_code = ?", magazineCode)
	}

	err := query.Find(&regularAgencys)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return regularAgencys, nil
}

