package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/karlhjm/golang-CI/entity"
	"github.com/karlhjm/golang-CI/loghelper"
	//"github.com/go-xorm/xorm"
)

func insertMenufood(v *entity.Menufood) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertMenufood Error:", affected, err)
		return err
	}
	//fmt.Println(v.ID)
	return nil
}

func findAllMenufoods() []entity.Menufood {
	vec := make([]entity.Menufood, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllMenufoods Error:", err)
	}
	return vec
}

func findMenufoodByName(name string) *entity.Menufood {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &entity.Menufood{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		loghelper.Error.Println("findMenufoodByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func deleteMenufood(v *entity.Menufood) error {
	if affected, err := engine.Delete(v); err != nil {
		loghelper.Error.Println("deleteMenufood Error:", affected, err)
		return err
	}
	return nil
}

func updateMenufood(origin, modify *entity.Menufood) error {
	if affected, err := engine.Update(modify, origin); err != nil {
		loghelper.Error.Println("updateMenufood Error:", affected, err)
		return err
	}
	return nil
}
