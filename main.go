package main

import (
	"net/http"
)

var templateRoot = "template/"

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add-cart-item", handleAddCartItem)
	http.HandleFunc("/log-in", handleLogin)
	http.HandleFunc("/verify", handleVerify)
	http.HandleFunc("/sign-up", handleSignUp)
	http.HandleFunc("/registry", handleRegistry)
	http.HandleFunc("/order", handleOrder)
	http.HandleFunc("/shopping-cart", handleShoppingCart)
	http.HandleFunc("/update-cart-item", handleUpdateCartItem)

	http.Handle("/js/", http.FileServer(http.Dir("")))

	http.ListenAndServe("localhost:8080", nil)
}

func getFiles(files []string) []string {
	for idx, file := range files {
		files[idx] = templateRoot + file
	}
	return files
}
