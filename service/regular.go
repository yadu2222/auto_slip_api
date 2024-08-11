package service

import (
	"auto_slip_api/model"
	"log"

	"github.com/google/uuid"
)

type RegularService struct{}

// 作成
func(s *RegularService) RegisterRegular(regular model.Regular) error {


	// UUIDを生成して追加
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	regular.RegularUuid = uid.String() // UUIDを文字列に変換して代入

	// 投げる
	err = model.RegisterRegular(regular)
	if err != nil {
		log.Println("定期の登録に失敗しました:", err)
		return err
	}
	return nil
}


// 定期の一括登録
func (s *RegularService) RegisterRegulars(regulars []model.Regular) error {
	for i := 0; i < len(regulars); i++ {
		// UUIDを生成して追加
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		regulars[i].RegularUuid = uid.String() // UUIDを文字列に変換して代入

		//　csv_idからcustomer_idを取得
		customer, err := model.FindCustomerByCsvID(regulars[i].CustomerUuid)
		if err != nil {
			log.Println("顧客IDの取得に失敗しました:", err)
			return err
		}
		regulars[i].CustomerUuid = customer.CustomerUuid
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

// 定期情報の構造体
type RegularCustomerInfo struct {
	RegularUuid string	`json:"regularUUID"`        // 定期情報ID
	Quantity    int    	`json:"quantity"`        // 冊数
	Customer    model.Customer `json:"customer"` // 顧客情報
}
type RegularMagazineInfo struct {
	RegularUuid string   `json:"regularUUID"`      // 定期情報ID
	Quantity    int      `json:"quantity"`      // 冊数
	Magaine     model.Magazine	`json:"magazine"` // 雑誌情報
}

// viewみたいな構造体
// 雑誌を主キーにした定期情報
type FindMagazineRegular struct {
	Magazine model.Magazine	`json:"magazine"` // 雑誌情報
	Regulars []RegularCustomerInfo	`json:"regulars"` // 定期情報
}

// 顧客を主キーにした定期情報
type FindCustomerRegular struct {
	Customer model.Customer	`json:"customer"` // 顧客情報
	Regulars []RegularMagazineInfo	`json:"regulars"` // 定期情報
}

// 雑誌を主キーに定期を一覧取得
func (s *RegularService) FindMagazineRegulars() ([]FindMagazineRegular, error) {
	var results []FindMagazineRegular
	// 雑誌情報を一覧取得
	magazines, err := model.GetMagazines()
	if err != nil {
		log.Println("定期情報の取得に失敗しました:", err)
		return nil, err
	}
	// 雑誌情報に合わせて定期情報を取得
	for _, magazine := range magazines {
		// 定期情報を取得
		regulars, err := model.FindRegularByMagazine(magazine.MagazineCode)
		if err != nil {
			log.Println("定期情報の取得に失敗しました:", err)
			return nil, err
		}

		// 定期情報に合わせてお客様情報を取得
		var regularInfos []RegularCustomerInfo
		for _, regular := range regulars {
			customer, err := model.FindCustomerByID(regular.CustomerUuid)
			if err != nil {
				log.Println("顧客情報の取得に失敗しました:", err)
				return nil, err
			}

			regularInfos = append(regularInfos, RegularCustomerInfo{
				RegularUuid: regular.RegularUuid,
				Quantity:    regular.Quantity,
				Customer:    customer,
			})

		}

		results = append(results, FindMagazineRegular{
			Magazine: magazine,
			Regulars: regularInfos,
		})

	}
	return results, nil
}

// 顧客を主キーに定期を一覧取得
func (s *RegularService) FindCustomerRegulars() ([]FindCustomerRegular, error) {
	var results []FindCustomerRegular
	// 顧客情報を一覧取得
	customers, err := model.GetCustomers()
	if err != nil {
		log.Println("定期情報の取得に失敗しました:", err)
		return nil, err
	}
	// 顧客情報に合わせて定期情報を取得
	for _, customer := range customers {
		// 定期情報を取得
		regulars, err := model.FindRegularByCustomer(customer.CustomerUuid)
		if err != nil {
			log.Println("定期情報の取得に失敗しました:", err)
			return nil, err
		}

		// 定期情報に合わせて雑誌情報を取得
		var regularInfos []RegularMagazineInfo
		for _, regular := range regulars {
			magazine, err := model.FindMagazineCode(regular.MagazineCode)
			if err != nil {
				log.Println("雑誌情報の取得に失敗しました:", err)
				return nil, err
			}

			regularInfos = append(regularInfos, RegularMagazineInfo{
				RegularUuid: regular.RegularUuid,
				Quantity:    regular.Quantity,
				Magaine:	magazine[0],
			})

		}

		results = append(results, FindCustomerRegular{
			Customer: customer,
			Regulars: regularInfos,
		})

	}
	return results, nil
}

// 削除
func(s *RegularService) DeleteRegular(regularUuid string) (*model.Regular, error) {
	regular, err := model.DeleteRegular(regularUuid)
	if err != nil {
		log.Println("定期情報の削除に失敗しました:", err)
		return nil, err
	}
	return regular, nil
}
