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

func InitOparateLogFK() error {
	_, err := db.Exec("ALTER TABLE oparate_logs ADD FOREIGN KEY (taker_uuid) REFERENCES employees(employee_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}