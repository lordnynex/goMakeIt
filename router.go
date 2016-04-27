package main

import (
	"net/http"
	"fmt"
)

func HelloMMG(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, world")
}

func NewRouter() http.Handler{
	router := http.NewServeMux()
	router.HandleFunc("/", HelloMMG)
	return router
}