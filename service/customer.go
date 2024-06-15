package service

import (
	"log"
	"auto_slip_api/model"
)

// 新しいグループの作成
func CreateCustomer(group *model.Magazine) error {
	_, err := DbEngine.Insert(group)
	if err != nil {
		log.Println("グループの作成に失敗しました:", err)
		return err
	}
	return nil
}

// 取得
func GetCustomerByID(id int64) (*model.Magazine, error) {
	group := new(model.Magazine)
	has, err := DbEngine.ID(id).Get(group)
	if err != nil {
		log.Println("グループの取得に失敗しました:", err)
		return nil, err
	}
	if !has {
		log.Println("指定されたIDのグループは存在しません")
		return nil, nil
	}
	return group, nil
}

func UpdateCustomer(magazine *model.Magazine) error {
	_, err := DbEngine.ID(magazine.MagazineUuid).Update(magazine)
	if err != nil {
		log.Println("グループの更新に失敗しました:", err)
		return err
	}
	return nil
}
