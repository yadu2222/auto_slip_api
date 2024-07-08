package service

import (
	"log"
	"auto_slip_api/model"
	"github.com/google/uuid"
)

type RegularService struct{}

// 作成
func CreateRegular(group *model.Magazine) error {
	_, err := DbEngine.Insert(group)
	if err != nil {
		log.Println("グループの作成に失敗しました:", err)
		return err
	}
	return nil
}

// 定期の一括登録
func(s *RegularService) RegisterRegulars(regulars []model.Regular) error {
	for i := 0; i < len(regulars); i++{
	// UUIDを生成して追加
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		regulars[i].RegularUuid = uid.String() // UUIDを文字列に変換して代入
	}
	// // 雑誌コードから雑誌IDを取得
	// for i := 0; i < len(regulars); i++ {
	// 	magazine, err := model.FindMagazineCode(regulars[i].MagazineUuid)
	// 	if err != nil {
	// 		log.Println("雑誌IDの取得に失敗しました:", err)
	// 		return err
	// 	}
	// 	regulars[i].MagazineUuid = magazine.MagazineUuid
	// }
	err := model.RegisterRegulars(regulars)
	if err != nil {
		log.Println("定期情報の登録に失敗しました:", err)
		return err
	}
	return nil
}

// 取得
func GetRegularByID(id int64) (*model.Magazine, error) {
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

// 更新
// func UpdateRegular(magazine *model.Magazine) error {
// 	_, err := DbEngine.ID(magazine.MagazineUuid).Update(magazine)
// 	if err != nil {
// 		log.Println("グループの更新に失敗しました:", err)
// 		return err
// 	}
// 	return nil
// }
