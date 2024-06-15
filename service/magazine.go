package service

import (
	"log"
	"github.com/google/uuid"
	"auto_slip_api/model"
)

type MagazineService struct{}

// 雑誌登録
func(s *MagazineService) RegisterMagazines(magazines []model.Magazine) error {
	// 雑誌の数だけループ
	// UUIDを生成して追加
	for i := 0; i < len(magazines); i++{
	// UUIDを生成して追加
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		magazines[i].MagazineUuid = uid.String() // UUIDを文字列に変換して代入
	}
	err := model.RegisterMagazines(magazines)
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
