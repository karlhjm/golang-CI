package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/karl-jm-huang/golang-CI/entity"
	"github.com/karl-jm-huang/golang-CI/loghelper"
	//"github.com/go-xorm/xorm"
)

func insertCustomer(v *entity.Customer) error {
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertCustomer Error:", affected, err)
		return err
	}
	return nil
}

func findAllCustomers() []entity.Customer {
	vec := make([]entity.Customer, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllCustomers Error:", err)
	}
	return vec
}

func findCustomerByName(name string) *entity.Customer {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &entity.Customer{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		loghelper.Error.Println("findCustomerByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}
