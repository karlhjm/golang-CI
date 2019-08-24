package service

import (
	"github.com/karlhjm/golang-CI/entity"
	"github.com/karlhjm/golang-CI/loghelper"
	//"fmt"
	//simplejson "github.com/bitly/go-simplejson"
	"github.com/karlhjm/golang-CI/db"
)

func CustomerRegister(name string, password string, restaurant_id int, phone string) (bool, error) {
	var v entity.Customer
	//v.ID = 0
	v.Name = name
	v.Password = password
	v.Restaurant_ID = restaurant_id
	v.Phone = phone
	if err := db.CreateCustomer(&v); err != nil {
		loghelper.Error.Println("User Register: Already exist username")
		return false, nil
	}
	return true, nil
}

func ListAllCustomers() []entity.Customer {
	return db.QueryCustomer(func(u *entity.Customer) bool {
		return true
	})
}

func GetCustomerByName(cname string) *entity.Customer {
	return db.QueryCustomerByName(cname)
}
