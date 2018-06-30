package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	"github.com/karl-jm-huang/golang-CI/loghelper"
	"github.com/karl-jm-huang/golang-CI/service"
	//"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

func MenufoodRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		defer req.Body.Close()
		//js, _ := simplejson.NewFromReader(req.Body)
		req.ParseForm()
		name := req.Form["name"][0]
		src := req.Form["src"][0]
		detail := req.Form["detail"][0]
		categorys := req.Form["categorys"][0]
		price, _ := strconv.ParseFloat(req.Form["price"][0], 64)
		flag, _ := service.MenufoodRegister(name, price, 0, categorys, detail, src)
		if flag == true {
			formatter.JSON(w, 201, struct {
				Status  int    `json:"status"`
				Success string `json:"success"`
			}{Status: 1, Success: "添加食品成功"}) // expected a user id
			loghelper.Info.Println("添加", name, "食品成功")
		} else {
			formatter.JSON(w, 404, nil)
			loghelper.Warning.Println(name, "名字重复")
			loghelper.Info.Println(name, "名字重复")
		}
	}
}

func ListAllMenufoodHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//fmt.Println("delete menufood!")
		req.ParseForm()
		id, _ := strconv.ParseInt(req.Form[`id`][0], 10, 64)
		flag := service.DeleteMenufood(int(id))
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

func ListAllMenufoodThroughTagHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		res := service.ListAllMenufoodsThroughTags()
		if len(res) != 0 {
			formatter.JSON(w, 200, res)
		} else {
			formatter.Text(w, 404, "No Menufoods!")
		}
	}
}
