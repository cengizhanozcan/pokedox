package main

import (
	"net/http"
	"pokemon/middleware"
)

func main() {

	http.ListenAndServe(":8080", middleware.NewRouter().InitRouter())
}
