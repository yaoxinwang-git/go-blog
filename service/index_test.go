package service

import (
	"go-blog/models"
	"reflect"
	"testing"
)

func TestGetAllIndexInfo(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		args    args
		want    *models.HomeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllIndexInfo(tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllIndexInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllIndexInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
