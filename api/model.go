package api

import (
	"encoding/xml"
	"fmt"
)

type RSSfeed struct {
	Items []struct {
		Title   string `xml:"title"`
		Link    string `xml:"link"`
		Guid    string `xml:"guid"`
		PubDate string `xml:"pubDate"`
	} `xml:"channel>item"`
}

type Article struct {
	Id          string `xml:"guid"`
	Title       string `xml:"title"`
	Url         string `xml:"link"`
	UpdatedTime string `xml:"pubDate"`
}

/*
xmlのbyte型をarticle型に変換
@param data XMLのbyte型データ
@return Article型, error
*/
func (article *Article) ParseRSS(data []byte) (Article, error) {
	var feed RSSfeed
	//デコードはRSSfeedへ
	if err := xml.Unmarshal(data, &feed); err != nil {
		fmt.Println("model:to Article Error")
		return Article{}, err
	}

	// RSSfeedからArticleへ変換
	article.Id = feed.Items[0].Guid
	article.Title = feed.Items[0].Title
	article.Url = feed.Items[0].Link
	article.UpdatedTime = feed.Items[0].PubDate
	/*
		articles := make([]Article, 0, len(feed.Items))
		for _, item := range feed.Items {
			if item.Guid == "" {
				return Article{}, fmt.Errorf("RSSfeed:guid is empty")
			}
			articles = append(articles, Article{
				Id:          item.Guid, //失敗のよう
				Title:       item.Title,
				Url:         item.Link,
				UpdatedTime: item.PubDate,
			})
		}
	*/

	//空フィールドは禁止
	if article.Id == "" {
		return Article{}, fmt.Errorf("model:Id is empty")
	}
	if article.Title == "" {
		return Article{}, fmt.Errorf("model:Title is empty")
	}
	if article.Url == "" {
		return Article{}, fmt.Errorf("model:Url is empty")
	}
	if article.UpdatedTime == "" {
		return Article{}, fmt.Errorf("model:UpdatedTime is empty")
	}

	return *article, nil
}
