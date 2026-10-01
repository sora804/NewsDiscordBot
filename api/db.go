package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
type DBFunc interface {
	//funcまとめる
	//ここやばいかも
	CreateDB()
	RegisterInfo(info string)
	Get()
	HardDelete()
}
*/

type Article struct {
	Id          string
	Title       string
	Url         string
	UpdatedTime string
}

//「--------------------------------------------------------
/*
func DBInit() {
	os.Remove("./articles.db")

	db, err := sql.Open("sqlite3", "./articles.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table Articles (id integer not null primary key, articles text, updated_time text);
	delete from Articles;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// todo:ここから続きコードリーディング
	stmt, err := tx.Prepare("insert into Articles(id, articles, updated_time) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i), time.Now().Format(time.RFC3339))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, articles, updated_time from Articles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var articles string
		var updatedTime string
		err = rows.Scan(&id, &articles, &updatedTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, articles, updatedTime)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select articles from Articles where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var articles string
	err = stmt.QueryRow("3").Scan(&articles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(articles)

	_, err = db.Exec("delete from Articles")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into Articles(id, articles, updated_time) values(1, 'foo', ?), (2, 'bar', ?), (3, 'baz', ?)", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, articles, updated_time from Articles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var articles string
		var updatedTime string
		err = rows.Scan(&id, &articles, &updatedTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, articles, updatedTime)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*
type DB interface {
	RegisterInfo(s string) error
}

type sqliteDB struct{ conn *sql.DB }


// DB初期化処理を書く
// もしテーブルがなければ
// 引数要らなくね？
func DBinit(path string) (DB, error) {
	conn, err := sql.Open("sqlite", path) // ① 接続を開く
	if err != nil {
		return DB{}, err
	}
	if err := conn.Ping(); err != nil { // ② 疎通確認
		return DB{}, err
	}
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		url TEXT UNIQUE
	)`) // ③ テーブル作成
	if err != nil {
		return DB{}, err
	}
	return DB{conn: conn}, nil
}
*/

func DBinit() (string, error) {
	// DB接続
	db, err := sql.Open("sqlite3", "database/Articles.db")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// テーブル作成
	/*
		コードかSQLか
	*/
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Articles (
            id           INTEGER PRIMARY KEY AUTOINCREMENT,
            articles     TEXT NOT NULL,
            updated_time TEXT
        )
    `)
	if err != nil {
		return "", err
	}

	return "OK", nil //、旦これ
}

/*
５分おきに自動登録

insert db もらう、登録
*/
//引数変えた
func RegisterInfo(info string) {
	//ここはルーティングからのinfoとからむかな～

	db, err := sql.Open("sqlite3", "database/Articles.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*
		下は変数名にする
	*/
	_, err = db.Exec(
		"INSERT INTO Articles (id, articles, updated_time) VALUES (?, ?, ?)",
		"1", info, time.Now().Format(time.RFC3339),
	)
	if err != nil {
		log.Fatal(err)
	}
}

/*
httpでrequestがあればデータを取得する
*/
func GetData() (string, error) {
	//dbアクセス
	db, err := sql.Open("sqlite3", "database/Articles.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//今
	now := time.Now()
	log.Println("現在時刻:", now)

	// SELECT（24時間以内）
	rows, err := db.Query("SELECT id, title, url, updated_time FROM Articles WHERE updated_time >= ?", now.Add(-24*time.Hour))
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
	return "", err
}

/*
１日おきに自動で４８時間以前をを削除する
*/
func HardDelete() {
	db, err := sql.Open("sqlite3", "database/Articles.db")
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
