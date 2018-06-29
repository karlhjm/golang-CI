package service

import (
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
	//"fmt"
	//simplejson "github.com/bitly/go-simplejson"
)

func CustomerRegister(name string, password string, restaurant_id int, phone string) (bool, error) {
	var v entity.Customer
	//v.ID = 0
	v.Name = name
	v.Password = password
	v.Restaurant_ID = restaurant_id
	v.Phone = phone
	if err := entity.CreateCustomer(&v); err != nil {
		loghelper.Error.Println("User Register: Already exist username")
		return false, nil
	}
	return true, nil
}

func ListAllCustomers() []entity.Customer {
	return entity.QueryCustomer(func(u *entity.Customer) bool {
		return true
	})
}

func GetCustomerByName(cname string) *entity.Customer {
	return entity.QueryCustomerByName(cname)
}