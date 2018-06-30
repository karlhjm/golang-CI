package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/karl-jm-huang/golang-CI/entity"
	"github.com/karl-jm-huang/golang-CI/loghelper"
	//"github.com/go-xorm/xorm"
)

func insertOrder(v *entity.Orders) (int, error) {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertOrder Error:", affected, err)
		return -1, err
	}
	return v.ID, nil
}

func findAllOrders() []entity.Orders {
	vec := make([]entity.Orders, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllOrders Error:", err)
	}
	return vec
}

func insertOrderfood(v *entity.Orderfood) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertOrderfood Error:", affected, err)
		return err
	}
	return nil
}

func findAllOrderfoods() []entity.Orderfood {
	vec := make([]entity.Orderfood, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllOrderfoods Error:", err)
	}
	return vec
}

func findOrderByID(id int) *entity.Orders {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &entity.Orders{ID: id}
	has, err := engine.Get(u)
	if err != nil {
		loghelper.Error.Println("findOrderByID Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func findOrderByPhone(phone string) *[]entity.Orders {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	vec := make([]entity.Orders, 0)
	err := engine.Where("customer_phone = ?", phone).Find(&vec)
	if err != nil {
		loghelper.Error.Println("findMenufoodByOrder_id Error:", err)
	}
	return &vec
}

func findOrderfoodByID(id int) *[]entity.Orderfood {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	// u := &Orderfood{Order_id: id}
	// has, err := engine.Get(u)
	vec := make([]entity.Orderfood, 0)
	err := engine.Where("order_id = ?", id).Find(&vec)
	if err != nil {
		loghelper.Error.Println("findMenufoodByOrder_id Error:", err)
	}
	return &vec
}

func deleteOrder(v *entity.Orders) error {
	if affected, err := engine.Delete(v); err != nil {
		loghelper.Error.Println("deleteOrder Error:", affected, err)
		return err
	}
	return nil
}

func deleteOrderfood(v *entity.Orderfood) error {
	if affected, err := engine.Delete(v); err != nil {
		loghelper.Error.Println("deleteOrderfood Error:", affected, err)
		return err
	}
	return nil
}
