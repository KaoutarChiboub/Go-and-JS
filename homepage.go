package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request, params httpRouter.params) {
	fmt.Println("Serving homepage")
	http.ServeFile(w, r, "./html/homepage.html")
}
