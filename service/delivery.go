package service

import (
	// "log"
	"auto_slip_api/model"
)

type DeliveryService struct{}

// 配達データ
type DeliveryData struct {
	CustomerName string             `json:"customerName"`
	CustomerUuid string             `json:"customerUuid"`
	Magazines    []DeliveryMagazine `json:"magazines"`
}

// 配達する本のデータ
type DeliveryMagazine struct {
	MagazineCode string `json:"magazineCode"`
	MagazineName string `json:"magazineName"`
	Quantity     int    `json:"quantity"`
	Number       string `json:"number"`
	Price        int    `json:"price"`
}

// 配達データを取得するよ
func (s *DeliveryService) GetDeliveryData(agencyList []model.Agency) ([]DeliveryData, error) {
	// 返すデータ
	deliveryData := []DeliveryData{}

	// 配達が必要な顧客を取得
	customers, err := model.FindCustomersNeedDelivery()
	if err != nil {
		return nil, err
	}

	// 顧客ごとに検索
	for _, customer := range customers {
		magazines := []DeliveryMagazine{}
		// 顧客の雑誌を取得
		for _, agency := range agencyList {
			quantity, err := model.FindRegularByCustomerAndMagazine(customer.CustomerUuid, agency.MagazineCode)
			if err != nil {
				return nil, err
			}
			deliveryMagazine := DeliveryMagazine{
				MagazineCode: agency.MagazineCode,
				MagazineName: agency.MagazineName,
				Quantity:     quantity,
				Number:       agency.Number,
				Price:        agency.Price,
			}
			// 配達データを追加
			if quantity != 0 {
				magazines = append(magazines, deliveryMagazine)
			}
		}

		// 顧客の配達データに整形
		delivery := DeliveryData{
			CustomerName: customer.CustomerName,
			CustomerUuid: customer.CustomerUuid,
			Magazines:    magazines,
		}
		// 配達データを追加

		if(len(magazines) != 0){
		deliveryData = append(deliveryData, delivery)
		}
	}

	return deliveryData, nil
}
