package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	"github.com/moandy/canyonsysu/loghelper"
	"github.com/moandy/canyonsysu/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/customers", CustomerRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/customers", ListAllCustomerHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/customers/{name:[_a-zA-Z0-9]+}", GetCustomerByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/restaurants", RestaurantRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/restaurants", ListAllRestaurantHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/v1/restaurants/{name:[_a-zA-Z0-9]+}", GetRestaurantByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/restaurant/update", RestaurantUpdateHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/menufoods", MenufoodRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/menufoods", ListAllMenufoodHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/v1/menufoods/{name:[_a-zA-Z0-9]+}", GetMenufoodByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/menufood/update", MenufoodUpdateHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/menufood/delete", MenufoodDeleteHandlerByID(formatter)).Methods("POST")
	mx.HandleFunc("/v1/orders", OrderRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/orders", ListAllOrderHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/order/delete", OrderDeleteHandlerByID(formatter)).Methods("POST")
	mx.HandleFunc("/v1/orderid", GetOrderByIDHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/orderphone", GetOrderByPhoneHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/comments", CommentRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/comments", ListAllCommentHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/tags", GetCommentCountsByTagHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/comments/scores", ListServiceHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/menufood/tag", ListAllMenufoodThroughTagHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/comment/offset", GetCommentByCountHandler(formatter)).Methods("GET")
}

func GetCommentByCountHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		begin, _ := strconv.Atoi(req.Form["begin"][0])
		offset, _ := strconv.Atoi(req.Form["offset"][0])
		res := service.ListCommentsByCount(begin, offset)
		formatter.JSON(w, 200, res)
	}
}

func ListAllMenufoodThroughTagHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		res := service.ListAllMenufoodsThroughTags()
		formatter.JSON(w, 200, res)
	}
}

func CustomerRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		name := req.PostForm[`name`][0]
		password := req.PostForm[`password`][0]
		phone := req.PostForm[`phone`][0]
		flag, _ := service.CustomerRegister(name, password, 1, phone)
		//flag, _ := api.CustomerRegister(req.PostForm)
		if flag == true {
			formatter.JSON(w, 201, nil)                                      // expected a user id
			http.Redirect(w, req, "users/"+req.PostForm[`username`][0], 201) //?
		} else {
			formatter.JSON(w, 404, nil)
		}
	}
}

func ListAllCustomerHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		res := service.ListAllCustomers()
		formatter.JSON(w, 200, res)
	}
}

func GetCustomerByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//fmt.Println("Get by name")
		//vars := mux.Vars(req)
		//name := vars["name"]
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		fmt.Println(name)
		customer := service.GetCustomerByName(name)
		fmt.Println(customer)
		if customer != nil {
			r.JSON(w, 200, customer)
		} else {
			r.JSON(w, 404, nil)
		}
	}
}

func RestaurantRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		name, _ := js.Get("name").String()
		address, _ := js.Get("location").String()
		//fmt.Println()
		certificates, _ := js.Get("certificates").String()
		servertime, _ := js.Get("server_time").String()
		flag, _ := service.RestaurantRegister(name, address, certificates, servertime)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
				// Name         string `json:"name"`
				// Location     string `json:"location"`
				// Certificates string `json:"certificates"`
				// Server_time  string `json:"server_time"`
			}{Status: 1, Success: "添加餐馆成功"}) // expected a user id
			//http.Redirect(w, req, "users/"+req.PostForm[`name`][0], 201) //?
			loghelper.Info.Println("添加", name, "餐馆成功")
		} else {
			formatter.JSON(w, 404, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 0, Success: "名字重复"})
			loghelper.Warning.Println(name, "名字重复")
			loghelper.Info.Println(name, "名字重复")
		}
	}
}

func ListAllRestaurantHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		if len(req.Form["name"][0]) != 0 {
			restaurant := service.GetRestaurantByName(req.Form["name"][0])
			if restaurant != nil {
				formatter.JSON(w, 200, restaurant)
				loghelper.Info.Println("Get the Restaurant by name!")
			} else {
				formatter.JSON(w, 404, nil)
				loghelper.Info.Println("None restaurant call the name!")
			}
			return
		}
		res := service.ListAllRestaurants()
		if len(res) == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("There isn't any restaurant now!")
		} else {
			formatter.JSON(w, 200, res)
			loghelper.Info.Println("Get All Restaurants!")
		}
	}
}

// func GetRestaurantByNameHandler(r *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin","*")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		req.ParseForm()
// 		path := filepath.FromSlash(req.RequestURI)
// 		_, name := filepath.Split(path)
// 		fmt.Println(name)
// 		restaurant := service.GetRestaurantByName(name)
// 		fmt.Println(restaurant)
// 		if restaurant != nil {
// 			r.JSON(w, 200, restaurant)
// 			loghelper.Info.Println("Get the Restaurant by name!")
// 		} else {
// 			r.JSON(w, 404, nil)
// 			loghelper.Info.Println("None restaurant call the name!")
// 		}
// 	}
// }

func RestaurantUpdateHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("update restaurant!")
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		name, _ := js.Get("name").String()
		location, _ := js.Get("location").String()
		certificates, _ := js.Get("certificates").String()
		server_time, _ := js.Get("server_time").String()
		flag := service.UpdateRestaurant(name, location, server_time, certificates)
		if flag == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("None restaurant call the name!")
		} else {
			formatter.JSON(w, 201, struct {
				Name         string `json:"name"`
				Location     string `json:"location"`
				Certificates string `json:"certificates"`
				Server_time  string `json:"server_time"`
			}{Name: name, Location: location, Certificates: certificates, Server_time: server_time})
			loghelper.Info.Println("Update the restaurant!")
		}
	}
}

func MenufoodRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		//id, _ := js.Get("id").Int()
		name, _ := js.Get("name").String()
		src, _ := js.Get("src").String()
		detail, _ := js.Get("detail").String()
		categorys, _ := js.Get("categorys").String()
		price, _ := js.Get("price").Float64()
		flag, _ := service.MenufoodRegister(name, price, 0, categorys, detail, src)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "添加食品成功"}) // expected a user id
			loghelper.Info.Println("添加", name, "食品成功")
			//http.Redirect(w, req, "users/"+req.PostForm[`name`][0], 201) //?
		} else {
			formatter.JSON(w, 404, nil)
			loghelper.Warning.Println(name, "名字重复")
			loghelper.Info.Println(name, "名字重复")
		}
	}
}

func ListAllMenufoodHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		if len(req.Form["name"][0]) != 0 {
			menufood := service.GetMenufoodByName(req.Form["name"][0])
			if menufood != nil {
				formatter.JSON(w, 200, menufood)
				loghelper.Info.Println("Get the menufood by name!")
			} else {
				formatter.JSON(w, 404, nil)
				loghelper.Info.Println("No such menufood called the name!")
				loghelper.Warning.Println("No such menufood called the name!")
			}
			return
		}
		res := service.ListAllMenufoods()
		if len(res) == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("There isn't any menufood now!")
			loghelper.Warning.Println("There isn't any menufood now!")
		} else {
			formatter.JSON(w, 200, res)
			loghelper.Info.Println("Get All Menufoods!")
		}
	}
}

func GetMenufoodByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		//fmt.Println(name)
		menufood := service.GetMenufoodByName(name)
		//fmt.Println(menufood)
		if menufood != nil {
			r.JSON(w, 200, menufood)
		} else {
			r.JSON(w, 404, nil)
			loghelper.Info.Println("No such menufood called the name!")
			loghelper.Warning.Println("No such menufood called the name!")
		}
	}
}

func MenufoodUpdateHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		id, _ := js.Get("id").Int()
		name, _ := js.Get("name").String()
		src, _ := js.Get("src").String()
		detail, _ := js.Get("detail").String()
		categorys, _ := js.Get("categorys").String()
		price, _ := js.Get("price").Float64()
		flag := service.UpdateMenufood(id, src, name, price, detail, categorys)
		if flag == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("No such food couldn't be update by id!")
		} else {
			formatter.JSON(w, 201, struct {
				Status    int     `json:"status"`
				Success   string  `json:"success"`
				Src       string  `json:"src"`
				Name      string  `json:"name"`
				Detail    string  `json:"detail"`
				Prices    float64 `json:"prices"`
				Categorys string  `json:"categorys"`
				//Published_at time.Time `json:"published_at"`
			}{Status: 1, Success: "更新成功", Src: src, Name: name, Detail: detail, Categorys: categorys, Prices: price})
			loghelper.Info.Println("Update the menufood!")
		}
	}
}

func MenufoodDeleteHandlerByID(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Println("delete menufood!")
		// req.ParseForm()
		// name := req.PostForm[`name`][0]
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		id, _ := js.Get("id").Int()
		flag := service.DeleteMenufood(id)
		if flag == 0 {
			formatter.JSON(w, 404, nil)
			loghelper.Info.Println("No such food couldn't be delete by id!")
			loghelper.Warning.Println("No such food couldn't be delete by id!")
		} else {
			formatter.JSON(w, 200, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "删除食品成功"})
			loghelper.Info.Println("Delete Menufood by ID!")
		}
	}
}

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
		w.Header().Set("Access-Control-Allow-Origin","*")
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

func CommentRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		js, _ := simplejson.NewFromReader(req.Body)
		order_id, _ := js.Get("order_id").Int()
		usr_name, _ := js.Get("usr_name").String()
		usr_photo, _ := js.Get("usr_photo").String()
		comment_at, _ := js.Get("comment_at").String()
		comment_star, _ := js.Get("comment_star").Int()
		tag, _ := js.Get("tag").String()
		client_text, _ := js.Get("client_text").String()
		merchant_text, _ := js.Get("merchant_text").String()
		flag, _ := service.CommentRegister(order_id, comment_star, comment_at, usr_name, tag, usr_photo, client_text, merchant_text)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "添加评论成功"}) // expected a user id
			loghelper.Info.Println("添加评论成功!")
		} else {
			formatter.JSON(w, 404, nil)
			loghelper.Error.Println("添加评论失败!")
		}
	}
}

func ListAllCommentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		res := service.ListAllComments()
		if len(res) == 0 {
			formatter.JSON(w, 404, res)
			loghelper.Info.Println("No Comment Now!")
		} else {
			formatter.JSON(w, 200, res)
			loghelper.Info.Println("List ALl Comments!")
		}
	}
}

func GetCommentCountsByTagHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["tag"][0]) != 0 {
			tag := req.Form["tag"][0]
			fmt.Println(tag)
			nums := service.GetCommentCountByTag(tag)
			if nums != 0 {
				r.JSON(w, 200, struct{
					Tag string `json:"tag"`
					Count int `json:"count"`
				}{Tag: tag, Count:nums})
				loghelper.Info.Println("Get CommentCounts by Tag!")
			} else {
				r.JSON(w, 404, nil)
				loghelper.Error.Println("Get CommentCounts by Tag fail!")
				loghelper.Info.Println("Get CommentCounts by Tag fail!")
			}
			return
		}
		res := service.ListAllTags()
		if len(res) == 0 {
			r.JSON(w, 404, nil)
		 	loghelper.Info.Println("No Tags Now!")
		} else {
			r.JSON(w, 200, res)
			loghelper.Info.Println("List All Tags!")
		}
	}
}

func ListServiceHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, 200, struct {
			Taste_score float32 `json:"taste_score"`
			Sight_score float32 `json:"sight_score"`
			Overall_score float32 `json:"overall_score"`
			Service_score float32 `json:"service_score"`
		} {Taste_score: 4.7, Sight_score: 4.7, Overall_score: 4.9, Service_score: 4.9})
	}
}