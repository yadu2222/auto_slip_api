package model

// "fmt"
// "math/rand"
// "time"
import (
	"log"
)

// groupeテーブル
// typeで型定義
type Magazine struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	// MagazineUuid string  `xorm:"varchar(36) pk" json:"magazineUUId"`
	MagazineCode string `xorm:"varchar(10) pk" json:"magazineCode"`
	MagazineName string `json:"magazineName"`
	TakerUuid    string `xorm:"varchar(36)" json:"takerUUID"`
	TakerName    string `xorm:"-" json:"takerName"`
	Note string `xorm:"varchar(100)" json:"note"`
}

func (Magazine) TableName() string {
	return "magazines"
}

// FK制約の追加
func InitMagazineFK() error {
	// UserTypeId
	_, err := db.Exec("ALTER TABLE magazines ADD FOREIGN KEY (taker_uuid) REFERENCES employees(employee_uuid) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		return err
	}
	return nil
}

// TODO: テストデータ作成
func CreateMagazineTestData() {

	// // 乱数生成器のシードを設定する（一般的には現在時刻を使う）
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// // 5桁の乱数を生成し、0で埋める
	// randomInt := rand.Intn(100000) // 0から99999までの乱数を生成
	// randomString := fmt.Sprintf("%05d", randomInt)

	// ロールデータを作成

	magazine1 := &Magazine{
		// MagazineUuid: "8f1a4b8f-29ec-4704-b364-1d2d55532673",
		MagazineCode: "29934",
		MagazineName: "少年ジャンプ",
		TakerUuid:    "c99cb6c4-42b9-4d6b-9884-ae6664f9df00",
	}
	db.Insert(magazine1)
}

// 重複チェックを行い、雑誌を登録する関数
func RegisterMagazine(magazine Magazine) error {

	exists, err := isMagazineExists(magazine)
	if err != nil {
		// エラーが発生した場合、ログを出力して処理を継続
		log.Printf("雑誌 %s の重複チェック中にエラーが発生しました: %v", magazine.MagazineName, err)
		return err
	}
	if exists {
		// 重複がある場合はログを出力して処理を継続
		log.Printf("雑誌 %s はすでに存在します", magazine.MagazineName)
		return err
	}

	// 雑誌を登録
	_, err = db.Insert(&magazine)
	if err != nil {
		// エラーが発生した場合、ログを出力して処理を継続します
		log.Printf("雑誌 %s の登録中にエラーが発生しました: %v", magazine.MagazineName, err)
		return err
	}
	log.Printf("雑誌 %s を登録しました", magazine.MagazineName)

	return nil
}

// 重複チェックを行い、雑誌を登録する関数
func RegisterMagazines(magazines []Magazine) error {
	for _, magazine := range magazines {
		exists, err := isMagazineExists(magazine)
		if err != nil {
			// エラーが発生した場合、ログを出力して処理を継続
			log.Printf("雑誌 %s の重複チェック中にエラーが発生しました: %v", magazine.MagazineName, err)
			return err
		}
		if exists {
			// 重複がある場合はログを出力して処理を継続
			log.Printf("雑誌 %s はすでに存在します", magazine.MagazineName)
			continue
		}

		// 雑誌を登録
		_, err = db.Insert(&magazine)
		if err != nil {
			// エラーが発生した場合、ログを出力して処理を継続します
			log.Printf("雑誌 %s の登録中にエラーが発生しました: %v", magazine.MagazineName, err)
			return err
		}
		log.Printf("雑誌 %s を登録しました", magazine.MagazineName)
	}
	return nil
}

// 指定された雑誌がすでに存在するかをチェックする関数
func isMagazineExists(magazine Magazine) (bool, error) {
	// ここで具体的に雑誌の重複チェックを実装します
	var count int64
	session := db.Where("magazine_code = ?", magazine.MagazineCode)
	count, err := session.Count(&Magazine{})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 雑誌コードから雑誌を取得
func FindMagazineByCode(magazineCode string) (Magazine, error) {
	var magazine Magazine
	has, err := db.Where("magazine_code = ?", magazineCode).Get(&magazine)
	if err != nil {
		return magazine, err
	}
	if !has {
		return magazine, err
	}
	return magazine, nil
}

// 雑誌コードから雑誌雑誌を取得
func FindMagazineCode(code string) ([]Magazine, error) {
	var magazine []Magazine
	session := db.Table("magazines")
	err := session.Where("magazine_code like ?", "%"+code+"%").Find(&magazine)
	if err != nil {
		print(err.Error())
		return magazine, err
	}
	// log.Printf("雑誌 %s を取得しました", magazine.MagazineName)
	return magazine, nil
}

// 雑誌名で部分一致検索
func FindMagazineName(name string) ([]Magazine, error) {
	var magazines []Magazine
	session := db.Table("magazines")
	err := session.Where("magazine_name like ?", "%"+name+"%").Find(&magazines)
	if err != nil {
		print(err.Error())
		return magazines, err
	}
	log.Printf("雑誌 %s を取得しました", name)
	return magazines, nil
}

// 雑誌一覧を取得
func GetMagazines() ([]Magazine, error) {
	var magazines []Magazine

	// なぜか作業者の名前がバインドできないのでここで再定義する
	// TODO:なおす
	type MagazineInfo struct {
		MagazineName string
		MagazineCode string
		TakerUuid    string
		TakerName    string
		Note string
	}
	var magazineInfos []MagazineInfo

	// dbに投げる
	err := db.Table("magazines").
		Join("left", "employees", "magazines.taker_uuid = employees.employee_uuid").
		Select("magazines.magazine_code,magazines.magazine_name,magazines.taker_uuid, employees.employee_name as taker_name,magazines.note").
		Find(&magazineInfos)
	if err != nil {
		log.Println("雑誌の取得に失敗しました:", err)
		return nil, err
	}
	// 再バインド
	for _, magazineInfo := range magazineInfos {
		magazine := Magazine{
			MagazineCode: magazineInfo.MagazineCode,
			MagazineName: magazineInfo.MagazineName,
			TakerUuid:    magazineInfo.TakerUuid,
			TakerName:    magazineInfo.TakerName,
			Note: magazineInfo.Note,
		}
		magazines = append(magazines, magazine)
	}
	// 返す
	return magazines, nil
}

// 雑誌を削除
func DeleteMagazine(magazineCode string) (*Magazine, error) {
	var magazine Magazine

	// 削除する前に雑誌の情報を取得
	has, err := db.Where("magazine_code = ?", magazineCode).Get(&magazine)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, err
	}

	// 雑誌を削除
	_, err = db.Where("magazine_code = ?", magazineCode).Delete(&Magazine{})
	if err != nil {
		return nil, err
	}

	// 削除した雑誌の情報を返す
	return &magazine, nil
}

// 雑誌情報を更新
func UpdateMagazine(magazine Magazine,oldMagazineCode string) error {
	_, err := db.Where("magazine_code = ?", oldMagazineCode).Update(&magazine)
	if err != nil {
		log.Println("雑誌情報の更新に失敗しました:", err)
		return err
	}
	return nil
}
