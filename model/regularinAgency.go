package model

import (
	"fmt"
)

type RegularAgency struct {
	RegularUuid   string `xorm:"varchar(36) not null" json:"regularUUID"`
	CustomerUuid  string `xorm:"varchar(36) not null" json:"customerUUID"`
	CustomerName  string `json:"customerName"`
	Quantity      int    `json:"quantity"`
	MethodType    string    `json:"methodType"`
}

// magazine_codeをもとに定期情報を取得
func FindCountingMagazine(magazineCode string) ([]RegularAgency, error) {
	var regularAgencys []RegularAgency

	err := db.Table("regulars").
		Where("regulars.magazine_code = ?", magazineCode).
		Join("LEFT", "magazines", "magazines.magazine_code = regulars.magazine_code").
		Join("LEFT", "customers", "customers.customer_uuid = regulars.customer_uuid").
		Join("LEFT", "method_types", "method_types.method_id = customers.method_type").
		Select("regulars.regular_uuid, customers.customer_uuid, customers.customer_name, regulars.quantity, method_types.method_type").
		Find(&regularAgencys)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return regularAgencys, nil
}
