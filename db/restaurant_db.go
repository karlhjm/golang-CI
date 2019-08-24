package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/karlhjm/golang-CI/entity"
	"github.com/karlhjm/golang-CI/loghelper"
	//"github.com/go-xorm/xorm"
)

func insertRestaurant(v *entity.Restaurant) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		//fmt.Println("insertRestaurant Error:", affected, err)
		loghelper.Error.Println("insertRestaurant Error:", affected, err)
		return err
	}
	return nil
}

func findAllRestaurants() []entity.Restaurant {
	vec := make([]entity.Restaurant, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllRestaurants Error:", err)
	}
	return vec
}

func findRestaurantByName(name string) *entity.Restaurant {
	u := &entity.Restaurant{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		loghelper.Error.Println("findRestaurantByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func updateRestaurant(origin, modify *entity.Restaurant) error {
	if affected, err := engine.Update(modify, origin); err != nil {
		loghelper.Error.Println("updateRestaurant Error:", affected, err)
		return err
	}
	return nil
}
