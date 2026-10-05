package api_test

import (
	"newsSite/api"
	"path/filepath"
	"testing"
)

func TestRegisterInfo_Handler(t *testing.T) {
	//databaseのパスをテスト用に設定
	api.DBPath = filepath.Join(t.TempDir(), "test.db")

	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := api.RegisterInfo_Handler()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("RegisterInfo_Handler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("RegisterInfo_Handler() succeeded unexpectedly")
			}
		})
	}
}

func TestGetData_Handler(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    api.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := api.GetData_Handler()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetData_Handler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetData_Handler() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetData_Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
