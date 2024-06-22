package csv

import (
	"fmt"
	"io"
	// CSVファイルを読み込むためのパッケージ
	"encoding/csv"
	"golang.org/x/text/transform"
	// 各文字コードに変換するためのパッケージ
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/japanese"
)


// TODO:エラーハンドリング

// 初期設定ファイル用
func ProcessUniCSVFile(file io.Reader) ([][]string, error) {
	reader := unicodeCsvReader(file)
	result,err := ProcessCSVFile(reader);
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 取次さんのファイル用
func ProcessJapCSVFile(file io.Reader) ([][]string, error) {
	reader := japaneseCsvReader(file)
	return ProcessCSVFile(reader)
}

// UTF8のリーダーを作成
func unicodeCsvReader(file io.Reader) *csv.Reader {
	decoder := transform.NewReader(file, unicode.UTF8.NewDecoder())
	reader := csv.NewReader(decoder)
	return reader
}
// ShiftJISのリーダーを作成
func japaneseCsvReader(file io.Reader) *csv.Reader {
	decoder := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	reader := csv.NewReader(decoder)
	return reader
}

// 読み込んで二次元配列に変換する
// リーダーを受け取る
func ProcessCSVFile(reader *csv.Reader) ([][]string, error) {
	// 2次元配列に変換
	var records [][]string	// 保存用配列
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	fmt.Println("data:", records)
	return records, nil
}