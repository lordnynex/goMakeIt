package main

import (
	"net/http"
	"fmt"
)

// HelloMMG is a http handler which says hello to you
func HelloMMG(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, world")
}

func newRouter() http.Handler{
	router := http.NewServeMux()
	router.HandleFunc("/", HelloMMG)
	return router
}