package service

import (
	"reflect"
	"testing"

	"github.com/moandy/canyonsysu/entity"
)

func TestCustomerRegister(t *testing.T) {
	type args struct {
		name          string
		password      string
		restaurant_id int
		phone         string
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
			got, err := CustomerRegister(tt.args.name, tt.args.password, tt.args.restaurant_id, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllCustomers(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Customer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllCustomers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCustomerByName(t *testing.T) {
	type args struct {
		cname string
	}
	tests := []struct {
		name string
		args args
		want *entity.Customer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCustomerByName(tt.args.cname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
