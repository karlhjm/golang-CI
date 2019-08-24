package server

import (
	"net/http"

	"github.com/karlhjm/golang-CI/loghelper"
	"github.com/karlhjm/golang-CI/service"
	//"path/filepath"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/unrolled/render"
	//"github.com/rs/cors"
)

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
			//fmt.Println(tag)
			nums := service.GetCommentCountByTag(tag)
			if nums != 0 {
				r.JSON(w, 200, struct {
					Tag   string `json:"tag"`
					Count int    `json:"count"`
				}{Tag: tag, Count: nums})
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

func GetCommentByCountHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		req.ParseForm()
		begin, _ := strconv.Atoi(req.Form["begin"][0])
		offset, _ := strconv.Atoi(req.Form["offset"][0])
		res := service.ListCommentsByCount(begin, offset)
		if len(res) != 0 {
			formatter.JSON(w, 200, res)
		} else {
			formatter.JSON(w, 200, struct {
				Status  int    `json:"status"`
				Message string `json:"message"`
			}{Status: -1, Message: "没有更多评论"})
		}
	}
}

func ListServiceHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, 200, struct {
			Taste_score   float32 `json:"taste_score"`
			Sight_score   float32 `json:"sight_score"`
			Overall_score float32 `json:"overall_score"`
			Service_score float32 `json:"service_score"`
		}{Taste_score: 4.7, Sight_score: 4.7, Overall_score: 4.9, Service_score: 4.9})
	}
}
