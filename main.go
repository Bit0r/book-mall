package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add-cart-item", handleAddCartItem)
	http.HandleFunc("/log-in", handleLogin)
	http.HandleFunc("/verify", handleVerify)
	http.HandleFunc("/sign-up", handleSignUp)
	http.HandleFunc("/registry", handleRegistry)

	http.Handle("/js/", http.FileServer(http.Dir("")))

	http.ListenAndServe("localhost:8080", nil)
}
