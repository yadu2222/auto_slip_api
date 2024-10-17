package service

import (
	// "log"
	"auto_slip_api/model"
	
)

type CountingService struct{}

// 本の情報を見出しに、顧客データをまとめた構造体
type Counting struct {
	Agency         model.Agency	`json:"agency"`
	Note 		 string		`json:"note"`
	RegularAgencys []model.RegularAgency	`json:"regularAgencys"`
	CountFlag      bool		`json:"countFlag"`
	LibraryCount	int `json:"livraryCount"`	
	DeliveryCount     int `json:"deliveryCount"`
	StoreCount        int `json:"storeCount"`
	StoreSlipCount int `json:"storeSlipCount"`
	HaulerCount int `json:"haulerCount"`
}

// かずをとるよ
func (s *CountingService) MagazineCounting(agencyList []model.Agency) ([]Counting, error) {
	//期限をキー、バリューを課題データのマップにする
	countings := []Counting{}
	for _, agency := range agencyList {
		countingList, err := model.FindCountingMagazine(agency.MagazineCode)
		if err != nil { //エラーハンドル エラーを上に投げるだけ
			return nil, err
		}

		// 定期で必要な冊数をカウント
		var count = 0
		// タイプごとにカウント
		// 図書館と配達と暁光高校は同じ
		// それ以外は違う?
		// それぞれのタイプごとにカウント
		var deliveryCount = 0
		var storeCount = 0
		var libraryCount = 0
		var storeSlipCount = 0
		var haulerCount = 0
		for _, counting := range countingList {
			count += counting.Quantity
			switch counting.MethodType {
			case 1, 5:
				deliveryCount += counting.Quantity
			case 2:
				storeCount += counting.Quantity
			case 3:
				storeSlipCount += counting.Quantity
			case 4:
				libraryCount += counting.Quantity
			case 6:
				haulerCount += counting.Quantity
			}
		}
		
		magazine,err := model.FindMagazineByCode(agency.MagazineCode)
		if err != nil {
			return nil, err
		}

		// agencyをキーにして、課題データのスライスをバリューにする
		// counting構造体を初期化
		counting := Counting{
			Agency:         agency,
			Note: magazine.Note,
			RegularAgencys: countingList,
			CountFlag:      agency.Quenity >= count,
			DeliveryCount: deliveryCount,
			StoreCount: storeCount,
			LibraryCount: libraryCount,
			StoreSlipCount: storeSlipCount,
			HaulerCount: haulerCount,
		}
		// スライスの追加

		if countingList != nil {
			countings = append(countings, counting)
		}
	}
	return countings, nil
}
