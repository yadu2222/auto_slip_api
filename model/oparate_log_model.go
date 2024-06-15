package model

import (
	"time"
)

type OparateLog struct {
	OparateLogId int `xorm:"pk autoincr" json:"oparateLogId"`
	TakerUuid string  `xorm:"varchar(20) not null" json:"takerUUID"`
	OparateLog string `xorm:"not null" json:"oparateLog"`
	OparateDate time.Time `xorm:"not null" json:"oparateDate"`
}

func (OparateLog) TableName() string {
	return "oparate_logs"
}

func CreateOparateLogTestData() {
	oparateLog1 := &OparateLog{
		TakerUuid: "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
		OparateLog: "雑誌を登録しました",
		OparateDate: time.Now(),
	}
	db.Insert(oparateLog1)
}

func InitOparateLogFK() error {
	_, err := db.Exec("ALTER TABLE oparate_logs ADD FOREIGN KEY (taker_uuid) REFERENCES employees(employee_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}