package api_test

import (
	"newsSite/api"
	"testing"
)

func TestArticle_ParseRSS(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		want    api.Article
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "全項目あり",
			data: []byte(`<item>
				<title>テスト記事</title>
				<link>https://zenn.dev/kyami/articles/97231d0de16f31</link>
				<guid>https://zenn.dev/kyami/articles/97231d0de16f31</guid>
				<pubDate>Sat, 03 Oct 2026 01:46:00 GMT</pubDate>
			</item>`),
			want: api.Article{
				Id:          "https://zenn.dev/kyami/articles/97231d0de16f31",
				Title:       "テスト記事",
				Url:         "https://zenn.dev/kyami/articles/97231d0de16f31",
				UpdatedTime: "Sat, 03 Oct 2026 01:46:00 GMT",
			},
			wantErr: false,
		},
		{
			name: "titleがCDATA",
			data: []byte(`<item>
				<title><![CDATA[CDATAのタイトル]]></title>
				<link>https://example.com/a</link>
				<guid>https://example.com/a</guid>
				<pubDate>Fri, 02 Oct 2026 13:37:55 GMT</pubDate>
			</item>`),
			want: api.Article{
				Id:          "https://example.com/a",
				Title:       "CDATAのタイトル",
				Url:         "https://example.com/a",
				UpdatedTime: "Fri, 02 Oct 2026 13:37:55 GMT",
			},
			wantErr: false,
		},
		{
			name: "未定義の要素は無視される",
			data: []byte(`<item>
				<title>追加要素あり</title>
				<link>https://example.com/b</link>
				<guid>https://example.com/b</guid>
				<pubDate>Thu, 01 Oct 2026 10:48:47 GMT</pubDate>
				<description>無視される</description>
			</item>`),
			want: api.Article{
				Id:          "https://example.com/b",
				Title:       "追加要素あり",
				Url:         "https://example.com/b",
				UpdatedTime: "Thu, 01 Oct 2026 10:48:47 GMT",
			},
			wantErr: false,
		},
		{
			name:    "要素が欠けるとゼロ値になる",
			data:    []byte(`<item><title>タイトルのみ</title></item>`),
			want:    api.Article{Title: "タイトルのみ"},
			wantErr: false,
		},
		{
			name:    "不正なXML",
			data:    []byte(`<item><title>閉じていない</item>`),
			want:    api.Article{},
			wantErr: true,
		},
		{
			name:    "空のデータ",
			data:    []byte(``),
			want:    api.Article{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article := &api.Article{}
			got, gotErr := article.ParseRSS(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseRSS() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseRSS() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("ParseRSS() = %v, want %v", got, tt.want)
			}
		})
	}
}
