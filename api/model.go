package api

type Article struct {
	Id          string
	Title       string
	Url         string
	UpdatedTime string
}

func ParseRSS(rss string) Article {
	// RSSの解析処理を書く
	// ここはあとでやる
	return Article{}
}

//データの型、jsonとか
//RSSとSQLiteのデータ型
