package main

import (
	"book-mall/model"
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
