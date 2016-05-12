package main

import (
	"fmt"
	"net/http"
	"os"
)

func getPort() (port string) {
	return os.Getenv("PORT")
}

func main() {
	fmt.Fprintf(os.Stdout, "Listening on %s", getPort())
	http.ListenAndServe(":"+getPort(), newRouter())
}
