package service

import (
	"reflect"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/karl-jm-huang/golang-CI/entity"
)

func TestOrderfoodRegister(t *testing.T) {
	type args struct {
		customer_phone string
		table_id       int
		order_contain  *simplejson.Json
		total          float64
		order_num      int
		time           string
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
			got, err := OrderfoodRegister(tt.args.customer_phone, tt.args.table_id, tt.args.order_contain, tt.args.total, tt.args.order_num, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderfoodRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OrderfoodRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllOrders(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Order_ins
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllOrders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteOrderBy(t *testing.T) {
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
			if got := DeleteOrderBy(tt.args.id); got != tt.want {
				t.Errorf("DeleteOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrderByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *entity.Order_ins
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrderByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrderByPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want []entity.Order_ins
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrderByPhone(tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
