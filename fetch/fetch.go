package fetch

import (
	"fmt"
	"io"
	"net/http"
)

/*
https://zenn.dev/feedのみをfetchする
*/

/*
zenn.devのRSSを取得する関数
@param url RSSフィードのURL
@return RSSフィードの内容を文字列で返す
*/

// 時間の範囲指定するかも
func getRSS(url string) string {
	rss := getRSS("https://zenn.dev/feed")
	fmt.Println(rss)
	resp, err := http.Get("https://zenn.dev/feed")
	if err != nil {
		// エラーハンドリングを書く
	}
	defer resp.Body.Close()

	// _を使うことでエラーを無視できる
	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

/*
これはmodelに移すかも
func formatTime(service string, publishedDate string, timeZone *time.Location, layout string) string {
	originalLayout := "Mon, 02 Jan 2006 15:04:05 GMT" // RFC1123
	if service == "note" {
		originalLayout = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123Z
	}
	formattedTime, err := time.ParseInLocation(originalLayout, publishedDate, timeZone)
	ErrorHandling(err)
	formattedTimeString := formattedTime.Format(layout)
	return formattedTimeString
}
*/
