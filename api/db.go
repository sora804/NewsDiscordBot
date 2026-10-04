package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

// databaseのパス設定
// テストと本番でカレントディレクトリが違うのを解決するため
var DBPath string = "database/Articles.db"

func DBinit() error {
	// DB接続
	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		fmt.Println("Failed to open DB:", err)
		return err
	}
	defer db.Close()

	// テーブル作成
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Articles (
            id           INTEGER PRIMARY KEY AUTOINCREMENT,
            title     TEXT NOT NULL,
            url         TEXT NOT NULL,
            updated_time TEXT
        )
    `)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	return nil
}

/*
５分おきに自動登録

insert db もらう、登録
*/
//引数変えた
func InsertInfo(info Article) error {
	//DB接続
	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return err
	}
	defer db.Close()

	// INSERTする値チェック
	if info.Id == "" || info.Title == "" || info.Url == "" || info.UpdatedTime == "" {
		return errors.New("empty field error")
	}

	//INSERT
	s, err := db.Exec(
		"INSERT INTO Articles (title, url, updated_time) VALUES (?, ?, ?)",
		info.Title, info.Url, time.Now().Format(time.RFC3339),
	)

	if err != nil {
		fmt.Println("Failed to execute insert:", err)
		return err
	}
	fmt.Println("Insert success:", s)

	return nil
}

/*
httpでrequestがあればデータを取得する
@param none
@return Article, error
*/
func GetData() (Article, error) {
	//dbアクセス
	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return Article{}, err
	}
	defer db.Close()

	//時間設定
	now := time.Now()
	log.Println("現在時刻:", now)

	var data Article

	// SELECT（24時間以内）
	rows, err := db.Query("SELECT id, title, url, updated_time FROM Articles WHERE updated_time >= ?", now.Add(-24*time.Hour))
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return Article{}, err
	}
	defer rows.Close()

	//INSERTしたのを読み込み
	for rows.Next() {
		rows.Scan(&data.Id, &data.Title, &data.Url, &data.UpdatedTime)
		fmt.Printf("ID:%s Title:%s URL:%s UpdatedTime:%s\n", data.Id, data.Title, data.Url, data.UpdatedTime)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("After Insert scan error:")
		return Article{}, err
	}
	return data, nil
}

/*
//１日おきに自動で４８時間以前をを削除する
func HardDelete() {
	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// DELETE
	//毎日１２時に発動
	//今の変数
	// ７２時間前に絞って消去

	now := time.Now()
	fmt.Println("現在時刻:", now)

	db.Exec("DELETE FROM Articles WHERE updated_time <= ?", now.Add(-72*time.Hour))

	fmt.Println("完了")
}
*/
