package service

import (
	"github.com/moandy/canyonsysu/db"
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
)

func MenufoodRegister(name string, price float64, restaurant_id int, categorys string, detail string, src string) (bool, error) {
	var v entity.Menufood
	v.Name = name
	v.Price = price
	v.Restaurant_id = restaurant_id
	v.Categorys = categorys
	v.Detail = detail
	v.Src = src
	if err := db.CreateMenufood(&v); err != nil {
		//fmt.Println("Menufood Register: Already exist Menufood")
		loghelper.Error.Println(err)
		return false, err
	}
	return true, nil
}

func ListAllMenufoods() []entity.Menufood {
	return db.QueryMenufood(func(u *entity.Menufood) bool {
		return true
	})
}

func ListAllMenufoodsThroughTags() []entity.Menufood_ins {
	menufoods := db.QueryMenufood(func(u *entity.Menufood) bool {
		return true
	})
	categorys := db.FindAllCategorys()
	var str []string
	for _, value := range categorys {
		str = append(str, value.Categorys)
	}
	//str := entity.FindAllCategorys()
	var res []entity.Menufood_ins
	for _, tag := range str {
		var a []entity.Menufood
		for _, value := range menufoods {
			if value.Categorys == tag {
				a = append(a, value)
			}
		}
		var b entity.Menufood_ins
		b.Categorys = tag
		b.Menufoods = a
		res = append(res, b)
	}
	return res
}

func GetMenufoodByName(rname string) *entity.Menufood {
	return db.QueryMenufoodByName(rname)
}

func UpdateMenufood(id int, src string, name string, price float64, detail string, categorys string) int {
	filter := func(m *entity.Menufood) bool {
		return m.ID == id
	}
	return db.UpdateMenufood(filter, func(arg1 *entity.Menufood) {
		arg1.Name = name
		arg1.Src = src
		arg1.Price = price
		arg1.Detail = detail
		arg1.Categorys = categorys
	})
}

func DeleteMenufood(id int) int {
	return db.DeleteMenufood(func(m *entity.Menufood) bool {
		return m.ID == id
	})
}
