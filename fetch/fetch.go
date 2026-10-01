package fetch

import (
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
	//rss := getRSS("https://zenn.dev/feed")
	//fmt.Println(rss)
	resp, err := http.Get(url)
	if err != nil {
		// エラーハンドリングを書く
		return "", err
	}
	defer resp.Body.Close()

	// _を使うことでエラーを無視できる
	body, err := io.ReadAll(resp.Body)
	if err != nil {
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
