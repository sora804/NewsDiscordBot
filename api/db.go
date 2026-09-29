package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	//funcまとめる
	//ここやばいかも
	CreateDB()
	RegisterInfo()
	Get()
	HardDelete()
}

type Article struct {
	Id          string
	Title       string
	Url         string
	UpdatedTime string
}

/*
５分おきに自動登録
*/
func (a Article) CreateDB() {
	// DB接続
	db, err := sql.Open("sqlite3", "database/article.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成
	/*
		コードかSQLか
	*/
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS articles (
            id           TEXT PRIMARY KEY,
            title        TEXT NOT NULL,
            url          TEXT NOT NULL,
            updated_time TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

}

/*
insert db もらう、登録
*/
func RegisterInfo() {
	//ここはルーティングからのinfoとからむかな～

	db, err := sql.Open("sqlite3", "database/article.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*
		下は変数名にする
	*/
	_, err = db.Exec(
		"INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
		"田中太郎", "taro@example.com", 27,
	)
	if err != nil {
		log.Fatal(err)
	}
}

/*
httpでrequestがあればデータを取得する
*/
func Get() {
	//dbアクセス
	db, err := sql.Open("sqlite3", "database/article.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//今
	now := time.Now()
	log.Println("現在時刻:", now)

	// SELECT（24時間以内）
	rows, err := db.Query("SELECT id, title, url, updated_time FROM articles WHERE updated_time >= ?", now.Add(-24*time.Hour))
	/*
		上やばいよ、
		コードかSQLか
	*/
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var title, url, updatedTime string
		rows.Scan(&id, &title, &url, &updatedTime)
		fmt.Printf("ID:%s Title:%s URL:%s UpdatedTime:%s\n", id, title, url, updatedTime)
	}
}

/*
１日おきに自動で４８時間以前をを削除する
*/
func HardDelete() {
	db, err := sql.Open("sqlite3", "database/article.db")
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

	db.Exec("DELETE FROM articles WHERE updated_time <= ?", now.Add(-72*time.Hour))

	fmt.Println("完了")
}
