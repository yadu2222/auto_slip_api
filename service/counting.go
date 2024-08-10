package service

import (
	// "log"
	"auto_slip_api/model"
)

type CountingService struct{}

// 本の情報を見出しに、顧客データをまとめた構造体
type Counting struct {
	Agency model.Agency
	RegularAgencys []model.RegularAgency
	CountFlag bool
}

// かずをとるよ
func(s *CountingService) MagazineCounting(agencyList []model.Agency) ([]Counting, error) {
	//期限をキー、バリューを課題データのマップにする
	countings := []Counting{}
	for _,agency := range agencyList {
		countingList, err := model.FindCountingMagazine(agency.MagazineCode)
		if err != nil { //エラーハンドル エラーを上に投げるだけ
			return nil, err
		}

		// agencyをキーにして、課題データのスライスをバリューにする
		// counting構造体を初期化
        counting := Counting{
            Agency:        agency,
            RegularAgencys: countingList,
			CountFlag:     agency.Quenity <= len(countingList),
        }
		// スライスの追加

		if(countingList != nil){
			countings = append(countings, counting)	
		}
	}
	return countings, nil
}