package api_test

import (
	"newsSite/api"
	"testing"
)

func TestRegisterInfo_Handler(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		// TODO: Add test cases.
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
