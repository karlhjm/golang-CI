package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	"github.com/karl-jm-huang/golang-CI/loghelper"
	"github.com/karl-jm-huang/golang-CI/service"
	//"fmt"
	"net/http"
	//"path/filepath"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"

	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

func OrderRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		table_id, _ := js.Get("table_id").Int()
		total, _ := js.Get("total").Float64()
		phone, _ := js.Get("customer_phone").String()
		order_contain := js.Get("order_contain")
		order_num, _ := js.Get("order_nums").Int()
		time, _ := js.Get("order_time").String()
		flag, _ := service.OrderfoodRegister(phone, table_id, order_contain, total, order_num, time)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "添加订单成功"}) // expected a user id
			loghelper.Info.Println("添加订单成功!")
		} else {
			formatter.JSON(w, 404, nil)
			loghelper.Error.Println("添加订单失败!")
		}
	}
}

func ListAllOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		res := service.ListAllOrders()
		if len(res) == 0 {
			formatter.JSON(w, 404, res)
			loghelper.Info.Println("No Order Now!")
		} else {
			formatter.JSON(w, 200, res)
			loghelper.Info.Println("List ALl Orders!")
		}
	}
}

func OrderDeleteHandlerByID(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		id, _ := js.Get("order_id").Int()
		flag := service.DeleteOrderBy(id)
		if flag == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("删除食物失败!")
			loghelper.Error.Println("删除食物失败!")
		} else {
			formatter.JSON(w, 200, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "删除食品成功"})
			loghelper.Info.Println("删除食品成功!")
		}
	}
}

// /v1/orderid/{orderid}
func GetOrderByIDHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["order_id"][0]) != 0 {
			id, err := strconv.ParseInt(req.Form["order_id"][0], 10, 32)
			loghelper.Error.Println(err)
			order := service.GetOrderByID(int(id))
			if order != nil {
				r.JSON(w, 200, order)
				loghelper.Info.Println("Get order by ID!")
			} else {
				r.JSON(w, 404, nil)
				loghelper.Error.Println("Get order by ID fail!")
				loghelper.Info.Println("Get order by ID fail!")
			}
			return
		}
		res := service.ListAllOrders()
		if len(res) == 0 {
			r.JSON(w, 404, res)
			loghelper.Info.Println("No Order Now!")
		} else {
			r.JSON(w, 200, res)
			loghelper.Info.Println("List ALl Orders!")
		}
	}
}

//  /v1/orderphone/?phone=
func GetOrderByPhoneHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//fmt.Println("Get by Phone")
		req.ParseForm()
		if len(req.Form["phone"][0]) != 0 {
			phone := req.Form["phone"][0]
			order := service.GetOrderByPhone(phone)
			if order != nil {
				r.JSON(w, 200, order)
				loghelper.Info.Println("Get Order by Phone!")
			} else {
				r.JSON(w, 404, nil)
				loghelper.Info.Println("Get Order by Phone fail!")
				loghelper.Error.Println("Get Order by Phone fail!")
			}
			return
		}
		res := service.ListAllOrders()
		if len(res) == 0 {
			r.JSON(w, 404, res)
			loghelper.Info.Println("No Order Now!")
		} else {
			r.JSON(w, 200, res)
			loghelper.Info.Println("List ALl Orders!")
		}
	}
}
