package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
