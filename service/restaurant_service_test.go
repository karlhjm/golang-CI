package service

import (
	"reflect"
	"testing"

	"github.com/karl-jm-huang/golang-CI/entity"
)

func TestRestaurantRegister(t *testing.T) {
	type args struct {
		name         string
		address      string
		certificates string
		servertime   string
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
			got, err := RestaurantRegister(tt.args.name, tt.args.address, tt.args.certificates, tt.args.servertime)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestaurantRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RestaurantRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllRestaurants(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Restaurant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllRestaurants(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllRestaurants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRestaurantByName(t *testing.T) {
	type args struct {
		rname string
	}
	tests := []struct {
		name string
		args args
		want *entity.Restaurant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRestaurantByName(tt.args.rname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRestaurantByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRestaurant(t *testing.T) {
	type args struct {
		name         string
		address      string
		servertime   string
		certificates string
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
			if got := UpdateRestaurant(tt.args.name, tt.args.address, tt.args.servertime, tt.args.certificates); got != tt.want {
				t.Errorf("UpdateRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
