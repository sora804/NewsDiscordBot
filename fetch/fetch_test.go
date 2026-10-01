package fetch

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 正常系: サーバーが返した本文がそのまま文字列で返ること
func TestGetRSS_Success(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "RSSのXMLを返す",
			body: `<?xml version="1.0"?><rss version="2.0"><channel><title>test</title></channel></rss>`,
		},
		{
			name: "日本語を含む本文",
			body: `<rss><channel><title>テスト記事</title></channel></rss>`,
		},
		{
			name: "空の本文",
			body: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer srv.Close()

			got := getRSS(srv.URL)
			if got != tt.body {
				t.Errorf("getRSS() = %q, want %q", got, tt.body)
			}
		})
	}
}

// 引数のURLが実際に使われていること（固定URLになっていないか）
func TestGetRSS_UsesGivenURL(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	_ = getRSS(srv.URL)

	if !called {
		t.Error("渡したURLのサーバーにリクエストが届いていない（URLがハードコードされている可能性）")
	}
}

// 異常系: 接続できないURLでもパニックせず、空文字を返すこと
func TestGetRSS_ConnectionError(t *testing.T) {
	// 起動してすぐ閉じることで、確実に接続できないURLを作る
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := srv.URL
	srv.Close()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("パニックした: %v", r)
		}
	}()

	got := getRSS(url)
	if got != "" {
		t.Errorf("getRSS() = %q, want empty string", got)
	}
}

/*
URLを取れて、
読めるか（logに吐く）どうか
*/
