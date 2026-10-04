package api

import (
	"fmt"
)

/*
dbテーブルを作成
これはmainが呼び出す、毎回やるのかな？
*/

/*
zenn.devのRSSを取得してDBに登録するハンドラー処理
url代入とfetch側の関数呼び出し
あと、この関数を呼び出すのをどっかで
*/
func RegisterInfo_Handler() error {
	//fetch部分
	url := "https://zenn.dev/feed"
	rss, err := FetchRSS(url)
	if err != nil {
		fmt.Println("failed to fetch RSS")
		return err
	}

	//パース部分
	parsedRSS := ParseRSS(rss)

	//dbに保存する関数呼び出し
	//DB, err := DBinit("database/Articles.db")
	err = DBinit()
	if err != nil {
		fmt.Println("DBinit handler failed:", err)
		return err
	}
	err = InsertInfo(parsedRSS)
	if err != nil {
		fmt.Println("InsertInfo handler failed:", err)
		return err
	}
	//上直す

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
