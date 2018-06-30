package service

import (
	"reflect"
	"testing"

	"github.com/karl-jm-huang/golang-CI/entity"
)

func TestMenufoodRegister(t *testing.T) {
	type args struct {
		name          string
		price         float64
		restaurant_id int
		categorys     string
		detail        string
		src           string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MenufoodRegister(tt.args.name, tt.args.price, tt.args.restaurant_id, tt.args.categorys, tt.args.detail, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("MenufoodRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MenufoodRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllMenufoods(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Menufood
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllMenufoods(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllMenufoods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllMenufoodsThroughTags(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Menufood_ins
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllMenufoodsThroughTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllMenufoodsThroughTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMenufoodByName(t *testing.T) {
	type args struct {
		rname string
	}
	tests := []struct {
		name string
		args args
		want *entity.Menufood
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMenufoodByName(tt.args.rname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMenufoodByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMenufood(t *testing.T) {
	type args struct {
		id        int
		src       string
		name      string
		price     float64
		detail    string
		categorys string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMenufood(tt.args.id, tt.args.src, tt.args.name, tt.args.price, tt.args.detail, tt.args.categorys); got != tt.want {
				t.Errorf("UpdateMenufood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMenufood(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteMenufood(tt.args.id); got != tt.want {
				t.Errorf("DeleteMenufood() = %v, want %v", got, tt.want)
			}
		})
	}
}
