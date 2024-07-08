package service

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"strconv"
	"strings"
)

// csvデータをマップ>>JSONに変換する
// 取次データ
func CsvToAgencyJson(records [][]string) ([]byte, error) {
	// some code
	// データの変換と格納
	var convertedList []map[string]string // 変換後のリストを格納するスライス
	for _, record := range records {
		// UUIDを生成して追加
		uid, _ := uuid.NewRandom()
		// マップを作成してリストに追加
		magazineMap := map[string]string{
			"countingUUID": uid.String(),
			"magazineCode": record[5],
			"magazineName": record[10],
			"number":record[6],
			"quenity":    record[11],
		}
		convertedList = append(convertedList, magazineMap)
	}
	// マップをJSONに変換
	byte, err := json.Marshal(convertedList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(byte))
	return byte, nil

}


// csvデータをマップ>>JSONに変換する
// 雑誌データ
func CsvToMagazineJson(records [][]string) ([]byte, error) {
	// some code
	// データの変換と格納
	var convertedList []map[string]string // 変換後のリストを格納するスライス
	for _, record := range records {
		// UUIDを生成して追加
		// uid, _ := uuid.NewRandom()
		// マップを作成してリストに追加
		magazineMap := map[string]string{
			"magazineCode": record[0],
			"magazineName": record[1],
			"takerUUID":    "c99cb6c4-42b9-4d6b-9884-ae6664f9df00", // とりあえず自分のid
		}
		convertedList = append(convertedList, magazineMap)
	}
	// マップをJSONに変換
	byte, err := json.Marshal(convertedList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(byte))
	return byte, nil

}

// csvデータをマップ>>JSONに変換する
// 顧客データ
func CsvToCustomerJSON(records [][]string) ([]byte, error) {
	// some code
	// データの変換と格納
	var convertedList []map[string]interface{} // 変換後のリストを格納するスライス

	for _, record := range records {
		// methodTypeを整数に変換
		methodType, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			methodType = 0 // デフォルト値を使用
			fmt.Printf("methodType の変換に失敗しました: %v\n", record[2])
		}

		// tellTypeを整数に変換
		tellType, err := strconv.Atoi(strings.TrimSpace(record[4]))
		if err != nil {
			tellType = 0 // デフォルト値を使用
			fmt.Printf("tellType の変換に失敗しました: %v\n", record[4])
		}

		// マップを作成してリストに追加
		customerMap := map[string]interface{}{
			"customerUUID": record[0],
			"customerName": record[1],
			"methodType":   methodType,
			"tellAddress":  record[3],
			"tellType":     tellType,
			"note":         record[5],
		}
		convertedList = append(convertedList, customerMap)
	}
	// マップをJSONに変換
	byte, err := json.Marshal(convertedList)
	if err != nil {
		fmt.Println(err)
		
	}
	fmt.Println(string(byte))
	return byte, nil

}


// csvデータをマップ>>JSONに変換する
// 定期データ
func CsvToRegularJSON(records [][]string) ([]byte, error) {
	// some code
	// データの変換と格納
	var convertedList []map[string]interface{} // 変換後のリストを格納するスライス

	for _, record := range records {
		// methodTypeを整数に変換
		quantity, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			quantity = 0 // デフォルト値を使用
			fmt.Printf("methodType の変換に失敗しました: %v\n", record[2])
		}

		// マップを作成してリストに追加
		customerMap := map[string]interface{}{
			"customerUUID": record[0],
			"magazineCode": record[1],
			"quantity": quantity,
		}
		convertedList = append(convertedList, customerMap)
	}
	// マップをJSONに変換
	byte, err := json.Marshal(convertedList)
	if err != nil {
		fmt.Println(err)
		
	}
	fmt.Println(string(byte))
	return byte, nil

}