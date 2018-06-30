package service

import (
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
	//"fmt"
	//simplejson "github.com/bitly/go-simplejson"
	"github.com/moandy/canyonsysu/db"
)

func RestaurantRegister(name string, address string, certificates string, servertime string) (bool, error) {
	//fmt.Println("2.RestaurantRegisterTest:")
	var v entity.Restaurant
	//v.ID = 0
	v.Name = name
	v.Address = address
	v.Certificates = certificates
	v.Servertime = servertime
	if err := db.CreateRestaurant(&v); err != nil {
		loghelper.Error.Println("Restaurant Register: Already exist Restaurant")
		return false, nil
	}
	return true, nil
}

func ListAllRestaurants() []entity.Restaurant {
	return db.QueryRestaurant(func(u *entity.Restaurant) bool {
		return true
	})
}

func GetRestaurantByName(rname string) *entity.Restaurant {
	return db.QueryRestaurantByName(rname)
}

func UpdateRestaurant(name string, address string, servertime string, certificates string) int {
	filter := func(m *entity.Restaurant) bool {
		return m.Name == name
	}
	return db.UpdateRestaurant(filter, func(arg1 *entity.Restaurant) {
		arg1.Name = name
		arg1.Servertime = servertime
		arg1.Address = address
		arg1.Certificates = certificates
	})
}
