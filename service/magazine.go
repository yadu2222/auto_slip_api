package service

import (
	"log"
	"auto_slip_api/model"
)

// 新しいグループの作成
func CreateMagazines(magazines []model.Magazine) error {
	_, err := DbEngine.Insert(magazines)
	if err != nil {
		log.Println("グループの作成に失敗しました:", err)
		return err
	}
	return nil
}

// 取得
func GetGroupByID(id int64) (*model.Magazine, error) {
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

func UpdateGroup(magazine *model.Magazine) error {
	_, err := DbEngine.ID(magazine.MagazineUuid).Update(magazine)
	if err != nil {
		log.Println("グループの更新に失敗しました:", err)
		return err
	}
	return nil
}
