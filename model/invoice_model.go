package model

import (
	"time"
)

// 請求を保存する
type InvoiceLog struct {
	InvoiceUuid      string `xorm:"varchar(36) pk" json:"invoiceUUID"`
	DeliveryUuid      string  `xorm:"varchar(36) not null unique" json:"deliveryUUId"`
	InvoiceDate	time.Time `xorm:"DATETIME not null" json:"invoiceDate"`
}

// テストデータ
func CreateInvoiceLogTestData() {
	invoiceLog1 := &InvoiceLog{
		InvoiceUuid: "ff934ac7-ab6c-4dc9-8449-e2bcb4b69d28",
		DeliveryUuid: "ac62957c-f86d-4814-95e0-ae8f86a126cd",
		InvoiceDate: time.Now(),
	}
	db.Insert(invoiceLog1)
}

// テーブル名
func (InvoiceLog) TableName() string {
	return "invoice_logs"
}

// FK制約の追加
func InitInvoiceLogFK() error {
	// 納品情報の外部キーとリレーション
	_, err := db.Exec("ALTER TABLE invoice_logs ADD FOREIGN KEY (delivery_uuid) REFERENCES delivery_logs(delivery_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}
