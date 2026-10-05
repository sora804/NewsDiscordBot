package api

import (
	"fmt"
)

/*
zenn.devのRSSを取得してDBに登録するハンドラー処理
*/
func RegisterInfo_Handler() error {
	//fetch部分
	url := "https://zenn.dev/feed"
	byteData, err := FetchRSS(url)
	if err != nil {
		fmt.Println("failed to fetch RSS")
		return err
	}

	//パース部分
	parsedData, err := (&Article{}).ParseRSS(byteData)
	if err != nil {
		fmt.Println("failed to parse RSS")
		return err
	}

	//dbに保存する関数呼び出し
	//DB, err := DBinit("database/Articles.db")
	//初めだけ
	err = DBinit()
	if err != nil {
		fmt.Println("DBinit handler failed:", err)
		return err
	}

	//データ登録
	err = InsertInfo(parsedData)
	if err != nil {
		fmt.Println("InsertInfo handler failed:", err)
		return err
	}

	return nil
}

/*
データ取得
@return Article, error
*/
func GetData_Handler() (Article, error) {
	//DBからデータ取得
	info, err := GetData()
	if err != nil {
		fmt.Printf("failed to get Data: %v", err)
		return Article{}, err
	}

	//登録されてなければ登録せよ
	if info.Id == "" {
		// zennから登録
		err = RegisterInfo_Handler()
		if err != nil {
			fmt.Printf("failed to register Info: %v", err)
			return Article{}, err
		}
		// 再度データ取得
		info, err = GetData()
		if err != nil {
			fmt.Printf("failed to get Data after Re-registering RSS: %v", err)
			return Article{}, err
		}
	}

	return info, nil
}
