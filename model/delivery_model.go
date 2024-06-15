package model

import (
	"time"
)

// 納品情報の保存
type DeliveryLog struct {
	DeliveryUuid string    `xorm:"varchar(36) pk" json:"deliveryUUID"`
	RegularUuid  string    `xorm:"varchar(36)" json:"regularUUID"`
	Price        int       `json:"price"`
	DeliveryDate time.Time `xorm:"DATETIME" json:"deliveryDate"`
	TakerUuid    string    `xorm:"varchar(36)" json:"takerUUID"`
}

func (DeliveryLog) TableName() string {
	return "delivery_logs"
}

// テストデータ
func CreateDeliveryLogTestData() {
	deliveryLog1 := &DeliveryLog{
		DeliveryUuid: "ac62957c-f86d-4814-95e0-ae8f86a126cd",
		RegularUuid:  "903e3147-1b8c-4e26-a5ee-f525a246e2df",
		Price:        1000,
		DeliveryDate: time.Now(),
		TakerUuid:    "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
	}
	db.Insert(deliveryLog1)
}

// FK制約の追加
func InitDeliveryLogFK() error {
	// regularid
	_, err := db.Exec("ALTER TABLE delivery_logs ADD FOREIGN KEY (regular_uuid) REFERENCES regulars(regular_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	// OuchiUuid
	_, err = db.Exec("ALTER TABLE delivery_logs ADD FOREIGN KEY (taker_uuid) REFERENCES employees(employee_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}
