package main

import (
	"book-mall/model"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kataras/go-sessions/v3"
)

func handleAddCartItem(w http.ResponseWriter, r *http.Request) {
	s := sessions.Start(w, r)
	id, ok := s.Get("userID").(uint64)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bookID, _ := strconv.ParseUint(r.URL.Query().Get("id"), 0, 64)
	err := model.AddCartItem(id, bookID)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	id := model.VerifyUser(r.FormValue("name"), r.FormValue("passwd"))
	if id != 0 {
		sessions.Start(w, r).Set("userID", id)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func handleRegistry(w http.ResponseWriter, r *http.Request) {
	id := model.AddUser(r.FormValue("name"), r.FormValue("passwd"))
	if id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		sessions.Start(w, r).Set("userID", id)
	}
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if _, ok := sessions.Start(w, r).Get("userID").(uint64); ok {

	} else {
		http.Redirect(w, r, "/log-in", http.StatusFound)
	}
}

func handleShoppingCart(w http.ResponseWriter, r *http.Request) {
	id, ok := sessions.Start(w, r).Get("userID").(uint64)
	if !ok {
		http.Redirect(w, r, "/log-in", http.StatusFound)
		return
	}
	files := []string{"layout.html", "navbar.html", "shopping-cart.html"}
	for idx, file := range files {
		files[idx] = templateRoot + file
	}
	t, _ := template.ParseFiles(files...)

	data := struct {
		model.CartBooks
		model.CartAddresses
	}{
		model.GetCartItems(id),
		model.GetCartAddresses(id)}

	t.ExecuteTemplate(w, "layout.html", data)
}

func handleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := sessions.Start(w, r).Get("userID").(uint64)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		values := r.URL.Query()
		bookID, _ := strconv.ParseUint(values.Get("id"), 0, 64)
		quantity, _ := strconv.Atoi(values.Get("quantity"))
		err := model.UpdateCartItem(userID, bookID, uint(quantity))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
