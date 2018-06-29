package server

import (
	//"fmt"
	//"net/http"
	//"path/filepath"
	//"strconv"
	//simplejson "github.com/bitly/go-simplejson"
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
	/* ---------------------------customer_handler.go--------------------------------*/
	mx.HandleFunc("/v1/customers", CustomerRegisterHandler(formatter)).Methods("POST")  //注册用户
	mx.HandleFunc("/v1/customers", ListAllCustomerHandler(formatter)).Methods("GET")   //显示所有用户
	mx.HandleFunc("/v1/customers/{name:[_a-zA-Z0-9]+}", GetCustomerByNameHandler(formatter)).Methods("GET")  //通过名字显示用户
	/* ---------------------------------end------------------------------------------*/

	/* -------------------------restaurant_handler.go--------------------------------*/
	mx.HandleFunc("/v1/restaurants", RestaurantRegisterHandler(formatter)).Methods("POST") //注册商家
	mx.HandleFunc("/v1/restaurants", ListAllRestaurantHandler(formatter)).Methods("GET") //显示所有商家
	mx.HandleFunc("/v1/restaurant/update", RestaurantUpdateHandler(formatter)).Methods("POST") //更改商家信to息
	/* ---------------------------------end------------------------------------------*/

	/* ---------------------------menufood_handler.go--------------------------------*/
	mx.HandleFunc("/v1/addfood", MenufoodRegisterHandler(formatter)).Methods("POST")  //添加菜品
	mx.HandleFunc("/v1/menufoods", ListAllMenufoodHandler(formatter)).Methods("GET")  //显示所有菜品
	mx.HandleFunc("/v1/menufood/update", MenufoodUpdateHandler(formatter)).Methods("POST")  //更改菜品
	mx.HandleFunc("/v1/menufood/delete", MenufoodDeleteHandlerByID(formatter)).Methods("POST")  //删除菜品
	mx.HandleFunc("/v1/menufood/tag", ListAllMenufoodThroughTagHandler(formatter)).Methods("GET") //显示菜单根据标签返回
	/* ---------------------------------end------------------------------------------*/
	
	/* ---------------------------orderfood_handler.go--------------------------------*/
	mx.HandleFunc("/v1/orders", OrderRegisterHandler(formatter)).Methods("POST")  //添加订单
	mx.HandleFunc("/v1/orders", ListAllOrderHandler(formatter)).Methods("GET")  //显示所有订单
	mx.HandleFunc("/v1/order/delete", OrderDeleteHandlerByID(formatter)).Methods("POST")  //删除订单
	mx.HandleFunc("/v1/orderid", GetOrderByIDHandler(formatter)).Methods("GET")  //通过id获取订单
	mx.HandleFunc("/v1/orderphone", GetOrderByPhoneHandler(formatter)).Methods("GET")  //通过手机号获取订单
	/* ---------------------------------end------------------------------------------*/
	
	/* ---------------------------comment_handler.go--------------------------------*/
	mx.HandleFunc("/v1/comments", CommentRegisterHandler(formatter)).Methods("POST")  //添加评论
	mx.HandleFunc("/v1/comments", ListAllCommentHandler(formatter)).Methods("GET")  //显示所有评论
	mx.HandleFunc("/v1/comment/tags", GetCommentCountsByTagHandler(formatter)).Methods("GET") //显示评论标签的数目
	mx.HandleFunc("/v1/comments/scores", ListServiceHandler(formatter)).Methods("GET")  //显示商家分数
	mx.HandleFunc("/v1/comment/offset", GetCommentByCountHandler(formatter)).Methods("GET")  //根据偏移量返回n条评论
	/* ---------------------------------end------------------------------------------*/
	
	/* ---------------------------categorys_handler.go--------------------------------*/
	mx.HandleFunc("/v1/menufood/categorys", ListAllCategorysHandler(formatter)).Methods("GET")  //显示所有菜单种类
	mx.HandleFunc("/v1/tags", TagRegisterHandler(formatter)).Methods("POST") //添加tag
	/* ---------------------------------end------------------------------------------*/
}