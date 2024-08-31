package model

import (
	"time"
)

// 納品情報の保存
type DeliveryLog struct {
	DeliveryUuid string    `xorm:"varchar(36) pk" json:"deliveryUUID"`
	CustomerUuid string    `xorm:"varchar(36) not null" json:"customerUUID"`
	MagazineName string    `xorm:"varchar(36) not null" json:"magazineName"`
	Quantity	 int       `json:"quantity"`
	Number       string    `json:"number"` // 号数
	Price        int       `json:"price"`
	DeliveryDate time.Time `xorm:"DATETIME" json:"deliveryDate"` // 納品日
	TakerUuid    string    `xorm:"varchar(36)" json:"takerUUID"`
}

func (DeliveryLog) TableName() string {
	return "delivery_logs"
}

// テストデータ
func CreateDeliveryLogTestData() {
	deliveryLog1 := &DeliveryLog{
		DeliveryUuid: "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d28",
		CustomerUuid: "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d29",
		MagazineName: "週刊少年ジャンプ",
		Quantity:     100,
		Number:       "2/15",
		Price:        1000,
		DeliveryDate: time.Now(),
		TakerUuid:    "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d30",

	}
	db.Insert(deliveryLog1)
}

// FK制約の追加
func InitDeliveryLogFK() error {
	// regularid
	_, err := db.Exec("ALTER TABLE delivery_logs ADD FOREIGN KEY (customer_uuid) REFERENCES customers(customer_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	// takeruuid
	_, err = db.Exec("ALTER TABLE delivery_logs ADD FOREIGN KEY (taker_uuid) REFERENCES employees(employee_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}
