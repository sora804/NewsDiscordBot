package fetch

import (
	"fmt"
	"io"
	"net/http"
)

/*
あとで時間の範囲指定する
handler側でURL＝https://zenn.dev/feed
*/

/*
引数のurlのRSSを取得する関数
@param url RSSフィードのURL
@return RSSフィードの内容を文字列で返す
*/
func FetchRSS(url string) (string, error) {
	//RSSで取ってくる
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch RSS:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 読み込んだRSSのbodyを文字列として返す
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read RSS body:", err)
		return "", err
	}

	return string(body), nil
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
