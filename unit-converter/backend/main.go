package main

import (
	"fmt"
	"net/http"
)


func handleServer (w http.ResponseWriter, req *http.Request) {

	fmt.Println(w, "hello")
	headers := req.Body
	fmt.Println("req", headers)

}

func main() {

	fmt.Println("Hello from server")

	http.HandleFunc("/handleServer", handleServer)
	http.ListenAndServe(":8090", nil)

}