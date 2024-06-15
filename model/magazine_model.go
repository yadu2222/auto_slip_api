package model

// "fmt"
// "math/rand"
// "time"

// groupeテーブル
// typeで型定義
type Magazine struct {
	// カラムを指定しない場合は、変数名がそのままカラム名になる
	// pk primaryKey
	// autoincr 自動インクリメント
	// json json化する際のキー名
	MagazineUuid string  `xorm:"varchar(36) pk" json:"magazineUUId"`
	MagazineCode string  `xorm:"varchar(10) not null unique" json:"magazineCode"`
	MagazineName string `json:"magazineName"`
	TakerUuid string `xorm:"varchar(36)" json:"takerUUID"`
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
func TestMagazine() *Magazine {

	// // 乱数生成器のシードを設定する（一般的には現在時刻を使う）
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// // 5桁の乱数を生成し、0で埋める
	// randomInt := rand.Intn(100000) // 0から99999までの乱数を生成
	// randomString := fmt.Sprintf("%05d", randomInt)

	// ロールデータを作成
	magazine := &Magazine{
		MagazineCode: "1",
		MagazineName: "test",
	}
	println(magazine)
	return magazine
}
