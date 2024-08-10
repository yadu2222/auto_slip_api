package service

import (
	"log"
	"github.com/google/uuid"
	"auto_slip_api/model"
)

type CustomerService struct{}

// 新しいお客様の登録
func(s *CustomerService) RegisterCustomer(customer model.Customer) error {
	// UUIDを生成して追加
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	customer.CustomerUuid = uid.String() // UUIDを文字列に変換して代入

	err = model.RegisterCustomer(customer)
	if err != nil {
		log.Println("お客様情報の登録に失敗しました:", err)
		return err
	}
	return nil
}

// お客様の一括登録
func(s *CustomerService) RegisterCustomers(customers []model.Customer) error {
	// for i := 0; i < len(customers); i++{
	// // // UUIDを生成して追加
	// // 	uid, err := uuid.NewRandom()
	// // 	if err != nil {
	// // 		return err
	// // 	}
	// // 	customers[i].CustomerUuid = uid.String() // UUIDを文字列に変換して代入
	// }

	err := model.RegisterCustomers(customers)
	if err != nil {
		log.Println("お客様情報の登録に失敗しました:", err)
		return err
	}
	return nil
}

// 取得
func FindCustomerByID(id int64) (*model.Magazine, error) {
	group := new(model.Magazine)
	has, err := DbEngine.ID(id).Get(group)
	if err != nil {
		log.Println("お客様情報の取得に失敗しました:", err)
		return nil, err
	}
	if !has {
		log.Println("指定されたIDのグループは存在しません")
		return nil, nil
	}
	return group, nil
}

func UpdateCustomer(magazine *model.Magazine) error {
	_, err := DbEngine.ID(magazine.MagazineCode).Update(magazine)
	if err != nil {
		log.Println("お客様情報の更新に失敗しました:", err)
		return err
	}
	return nil
}

// 顧客一覧を取得
func(s *CustomerService) GetCustomers() ([]model.Customer, error) {

	var customers []model.Customer

	customers, err := model.GetCustomers()
	if err != nil {
		log.Println("お客様情報の取得に失敗しました:", err)
		return nil, err
	}
	return customers, nil
}