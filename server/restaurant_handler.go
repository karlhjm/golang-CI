package server

import (
	//"go-agenda-service/service/entity"
	//"go-agenda-service/service/service"
	"fmt"
	"net/http"

	"github.com/karl-jm-huang/golang-CI/loghelper"
	"github.com/karl-jm-huang/golang-CI/service"
	//"path/filepath"
	//"strconv"
	simplejson "github.com/bitly/go-simplejson"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

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
			}{Status: 1, Success: "添加餐馆成功"}) // expected a user id
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
