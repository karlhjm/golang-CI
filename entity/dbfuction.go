package entity

import (
	"github.com/moandy/canyonsysu/loghelper"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // for init
	//"github.com/go-xorm/xorm"
)

func insertCategorys(v *Categorys) error {
	if affected, err := engine.Insert(v); err != nil {
		fmt.Println("insertCategorys Error:", affected, err)
		return err
	}
	return nil
}

func insertCustomer(v *Customer) error {
	if affected, err := engine.Insert(v); err != nil {
		fmt.Println("insertCustomer Error:", affected, err)
		return err
	}
	return nil
}

func findAllCustomers() []Customer {
	vec := make([]Customer, 0)
	if err := engine.Find(&vec); err != nil {
		fmt.Println("findAllCustomers Error:", err)
	}
	return vec
}

func findCustomerByName(name string) *Customer {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &Customer{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		fmt.Println("findCustomerByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func insertRestaurant(v *Restaurant) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		//fmt.Println("insertRestaurant Error:", affected, err)
		loghelper.Error.Println("insertRestaurant Error:", affected, err)
		return err
	}
	return nil
}

func findAllRestaurants() []Restaurant {
	vec := make([]Restaurant, 0)
	if err := engine.Find(&vec); err != nil {
		fmt.Println("findAllRestaurants Error:", err)
	}
	return vec
}

func findRestaurantByName(name string) *Restaurant {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &Restaurant{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		fmt.Println("findRestaurantByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func updateRestaurant(origin, modify *Restaurant) error {
	if affected, err := engine.Update(modify, origin); err != nil {
		fmt.Println("updateRestaurant Error:", affected, err)
		return err
	}
	return nil
}

func insertMenufood(v *Menufood) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		fmt.Println("insertMenufood Error:", affected, err)
		return err
	}
	fmt.Println(v.ID)
	return nil
}

func findAllMenufoods() []Menufood {
	vec := make([]Menufood, 0)
	if err := engine.Find(&vec); err != nil {
		fmt.Println("findAllMenufoods Error:", err)
	}
	return vec
}

func findMenufoodByName(name string) *Menufood {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &Menufood{Name: name}
	has, err := engine.Get(u)
	if err != nil {
		fmt.Println("findMenufoodByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func deleteMenufood(v *Menufood) error {
	if affected, err := engine.Delete(v); err != nil {
		fmt.Println("deleteMenufood Error:", affected, err)
		return err
	}
	return nil
}

func updateMenufood(origin, modify *Menufood) error {
	if affected, err := engine.Update(modify, origin); err != nil {
		fmt.Println("updateMenufood Error:", affected, err)
		return err
	}
	return nil
}

func insertOrder(v *Orders) (int, error) {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		fmt.Println("insertOrder Error:", affected, err)
		return -1, err
	}
	return v.ID, nil
}
func findAllOrders() []Orders {
	vec := make([]Orders, 0)
	if err := engine.Find(&vec); err != nil {
		fmt.Println("findAllOrders Error:", err)
	}
	return vec
}

func insertOrderfood(v *Orderfood) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		fmt.Println("insertOrderfood Error:", affected, err)
		return err
	}
	return nil
}

func findAllOrderfoods() []Orderfood {
	vec := make([]Orderfood, 0)
	if err := engine.Find(&vec); err != nil {
		fmt.Println("findAllOrderfoods Error:", err)
	}
	return vec
}

func findOrderByID(id int) *Orders {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &Orders{ID: id}
	has, err := engine.Get(u)
	if err != nil {
		fmt.Println("findOrderByID Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func findOrderByPhone(phone string) *[]Orders {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	vec := make([]Orders, 0)
	err := engine.Where("customer_phone = ?", phone).Find(&vec)
	if err != nil {
		fmt.Println("findMenufoodByOrder_id Error:", err)
	}
	return &vec
}

func findOrderfoodByID(id int) *[]Orderfood {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	// u := &Orderfood{Order_id: id}
	// has, err := engine.Get(u)
	vec := make([]Orderfood, 0)
	err := engine.Where("order_id = ?", id).Find(&vec)
	if err != nil {
		fmt.Println("findMenufoodByOrder_id Error:", err)
	}
	return &vec
}

func deleteOrder(v *Orders) error {
	if affected, err := engine.Delete(v); err != nil {
		fmt.Println("deleteOrder Error:", affected, err)
		return err
	}
	return nil
}

func deleteOrderfood(v *Orderfood) error {
	if affected, err := engine.Delete(v); err != nil {
		fmt.Println("deleteOrderfood Error:", affected, err)
		return err
	}
	return nil
}

func insertComment(v *Comment) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertComment Error:", affected, err)
		return err
	}
	return nil
}

func findAllComments() []Comment {
	vec := make([]Comment, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllComments Error:", err)
	}
	return vec
}

func findCommentCountsByTag(tag string) int {
	comment := new(Comment)
	total, err := engine.Where("tags =?", tag).Count(comment)
	if err != nil {
		loghelper.Error.Println(err)
	}
	return int(total)
}

func findAllMenufoodByTag() (int, []string) {
	menufood := new(Menufood)
	tagnum, err := engine.Where("").Count(menufood)
	if err != nil {
		loghelper.Error.Println(err)
	}
	vec := make([]Menufood, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println(err)
	}
	var str []string
	for _, value := range vec {
		str = append(str, value.Categorys)
	}
	return int(tagnum), str
}

func findCommentByCount(begin int, offset int) []Comment {
	vec := make([]Comment, 0)
	sql := "select * from comment order by i_d desc limit ?, ?"
	engine.Sql(sql, begin, offset).Find(&vec)
	return vec
}

func FindAllCategorys() []Categorys {
	sql := "select * from categorys"
	vec := make([]Categorys, 0)
	engine.Sql(sql).Find(&vec)	
	return vec
}