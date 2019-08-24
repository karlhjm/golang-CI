package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	"github.com/karlhjm/golang-CI/loghelper"
	"github.com/karlhjm/golang-CI/service"
	//"fmt"
	"net/http"
	//"path/filepath"
	//"strconv"

	//simplejson "github.com/bitly/go-simplejson"

	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

func TagRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		tag := req.Form["tag"][0]
		flag, err := service.CategoryRegister(tag)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "添加种类成功"}) // expected a user id
			loghelper.Info.Println("添加", tag, "种类成功")
		} else {
			formatter.JSON(w, 404, err)
			loghelper.Warning.Println(err)
			loghelper.Info.Println(err)
			loghelper.Error.Println(err)
		}
	}
}

func ListAllCategorysHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		categorys := service.ListAllCategorys()
		if len(categorys) != 0 {
			formatter.JSON(w, 200, categorys)
		} else {
			formatter.Text(w, 404, "No categorys!")
		}
	}
}
