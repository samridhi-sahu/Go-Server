package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful\n")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")
	pin := r.FormValue("pin")
	city := r.FormValue("city")
	fmt.Fprintf(w, "email = %s\n", email)
	fmt.Fprintf(w, "password = %s\n", password)
	fmt.Fprintf(w, "address = %s\n", address)
	fmt.Fprintf(w, "pin = %s\n", pin)
	fmt.Fprintf(w, "city = %s\n", city)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 error not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Start at port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
