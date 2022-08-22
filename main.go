package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/person", personHandler)
	fmt.Println("Starting Server at 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Only Get Method Allowed", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, "Hello Sir!")
}
func personHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Only POST Method Allowed", http.StatusBadRequest)
		return
	}
	type Person struct {
		Name string
		Age  int
	}
	var body Person
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	fmt.Fprintf(res, "Hello %s!\nYour Age is %v", body.Name, body.Age)
}
