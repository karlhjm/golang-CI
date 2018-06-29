package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	//"github.com/moandy/canyonsysu/loghelper"
	"github.com/moandy/canyonsysu/service"
	"fmt"
	"net/http"
	"path/filepath"	
	"github.com/unrolled/render"
)

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
			//http.Redirect(w, req, "users/"+req.PostForm[`username`][0], 201) //?
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
