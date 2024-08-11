package service

import (
	"log"
	// "github.com/google/uuid"
	"auto_slip_api/model"
)

type MagazineService struct{}

// 雑誌登録
func(s *MagazineService) RegisterMagazines(magazines []model.Magazine) error {
	// 雑誌の数だけループ
	// UUIDを生成して追加
	// for i := 0; i < len(magazines); i++{
	// // UUIDを生成して追加
	// 	uid, err := uuid.NewRandom()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	magazines[i].MagazineUuid = uid.String() // UUIDを文字列に変換して代入
	// }
	err := model.RegisterMagazines(magazines)
	if err != nil {
		log.Println("雑誌の登録に失敗しました:", err)
		return err
	}
	return nil
}

// 雑誌一覧を取得
func(s *MagazineService) GetMagazines() ([]model.Magazine, error) {
	results, err := model.GetMagazines()
	if err != nil {
		log.Println("雑誌の取得に失敗しました:", err)
		return nil, err
	}
	return results, nil
}

// 雑誌を削除
func(s *MagazineService) DeleteMagazine(magazineCode string) (*model.Magazine, error) {
	magazine, err := model.DeleteMagazine(magazineCode)
	if err != nil {
		log.Println("雑誌の削除に失敗しました:", err)
		return nil, err
	}
	return magazine, nil
}
// 雑誌コードで検索
func(s *MagazineService) FindMagazineByCode(magazineCode string) ([]model.Magazine, error) {
	
	magazine, err := model.FindMagazineCode(magazineCode)
	if err != nil {
		log.Println("雑誌情報の取得に失敗しました:", err)
		return nil, err
	}
	return magazine, nil
}

// 雑誌名で検索
func(s *MagazineService) FindMagazineByName(magazineName string) ([]model.Magazine, error) {
	
	magazine, err := model.FindMagazineName(magazineName)
	if err != nil {
		log.Println("雑誌情報の取得に失敗しました:", err)
		return nil, err
	}
	return magazine, nil
}
