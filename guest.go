package main

import (
	"book-mall/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kataras/go-sessions/v3"
)

type Pages struct {
	Pre, Cur, Next, Total uint64
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var step uint64 = 12
	data := struct {
		Category   string
		Categories []string
		Pages
		model.Books
	}{}
	query := r.URL.Query()

	// 填充类别信息
	data.Category = query.Get("category")
	data.Categories = model.GetCategories()

	// 填充分页信息
	data.Cur, _ = strconv.ParseUint(query.Get("page"), 0, 64)
	data.Total = model.CountBooks(data.Category)/step + 1
	if data.Cur == 0 {
		data.Cur = 1
	} else if data.Cur > data.Total {
		data.Cur = data.Total
	}
	data.Pre, data.Next = data.Cur-1, data.Cur+1

	// 填充图书信息
	data.Books = model.GetBooks(data.Category, uint64((data.Cur-1)*step), uint64(step))

	files := []string{"layout.html", "home.html", "navbar-guest.html"}

	if sessions.Start(w, r).Get("userID") != nil {
		files[2] = "navbar.html"
	}

	t, _ := template.ParseFiles(getFiles(files)...)

	t.ExecuteTemplate(w, "layout.html", data)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/layout-no-nav.html", "template/log-in.html")
	t.Execute(w, "layout-no-nav.html")
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/layout-no-nav.html", "template/sign-up.html")
	t.Execute(w, "layout-no-nav.html")
}
