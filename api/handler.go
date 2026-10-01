package api

import (
	"fmt"
	"log"
	"newsSite/fetch"
)

/*
dbテーブルを作成
これはmainが呼び出す、毎回やるのかな？
*/

/*
Todo
・エラーハンドリングの追加
*/

/*
zenn.devのRSSを取得してDBに登録するハンドラー処理
url代入とfetch側の関数呼び出し
あと、この関数を呼び出すのをどっかで
*/
func RegisterRSS_Handler() string {
	//fetch部分
	url := "https://zenn.dev/feed"
	rss, err := fetch.FetchRSS(url)
	if err != nil {
		log.Println("failed to fetch RSS")
		log.Fatal(err)
	}

	//パース部分
	parsedRSS := ParseRSS(rss)

	//dbに保存する関数呼び出し
	//DB, err := DBinit("database/Articles.db")
	DB, err := DBinit()
	if DB != "OK" || err != nil {
		log.Fatal(err)
	}
	RegisterInfo(parsedRSS)
	//上直す

	return parsedRSS
}

/*
データ取得
@return string, error
*/
func GetRSS_Handler() (string, error) {
	info, err := GetData()

	//登録されてなければ登録せよ
	if info == "" {
		// zennから登録
		RegisterRSS_Handler()
	}

	if err != nil {
		fmt.Errorf("failed to get RSS: %v", err)
		return "", err
	}
	return info, nil
}
