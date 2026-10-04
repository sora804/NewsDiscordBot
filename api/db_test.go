package api_test

import (
	"newsSite/api"
	"path/filepath"
	"testing"
	"time"
)

func TestDBinit_Success(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		{
			name:    "success_case",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := api.DBinit()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("DBinit() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("DBinit() succeeded unexpectedly")
			}
		})
	}
}

func TestInsertInfo(t *testing.T) {
	//Timeの設定
	now := time.Now()
	// DBのパスを設定
	api.DBPath = filepath.Join(t.TempDir(), "test.db")

	if err := api.DBinit(); err != nil {
		t.Fatalf("DBinit() failed: %v", err)
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		info    api.Article
		wantErr bool
	}{
		{
			name:    "success_case",
			info:    api.Article{Id: "1", Title: "Title1", Url: "https://example.com", UpdatedTime: now.Format(time.RFC3339)},
			wantErr: false,
		},
		{
			name:    "idが空",
			info:    api.Article{Id: "", Title: "title2", Url: "https://example2.com", UpdatedTime: now.Format(time.RFC3339)},
			wantErr: true,
		},
		{
			name:    "titleが空",
			info:    api.Article{Id: "2", Title: "", Url: "https://example2.com", UpdatedTime: now.Format(time.RFC3339)},
			wantErr: true,
		},
		{
			name:    "urlが空",
			info:    api.Article{Id: "3", Title: "Title3", Url: "", UpdatedTime: now.Format(time.RFC3339)},
			wantErr: true,
		},
		{
			name:    "updated_timeが空",
			info:    api.Article{Id: "4", Title: "Title4", Url: "https://example4.com", UpdatedTime: ""},
			wantErr: true,
		},
		{
			name:    "全て空",
			info:    api.Article{Id: "", Title: "", Url: "", UpdatedTime: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := api.InsertInfo(tt.info)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("InsertInfo() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("InsertInfo() succeeded unexpectedly")
			}
		})
	}
}

func TestGetData(t *testing.T) {
	//Timeの設定
	now := time.Now()
	// DBのパスを設定
	api.DBPath = filepath.Join(t.TempDir(), "test.db")

	// テーブル作成
	if err := api.DBinit(); err != nil {
		t.Fatal(err)
	}
	// テストデータ
	if err := api.InsertInfo(api.Article{Id: "1", Title: "Title1", Url: "https://example.com", UpdatedTime: now.Format(time.RFC3339)}); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string // description of this test case
		want    api.Article
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success_case",
			want:    api.Article{Id: "1", Title: "Title1", Url: "https://example.com", UpdatedTime: now.Format(time.RFC3339)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := api.GetData()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetData() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
