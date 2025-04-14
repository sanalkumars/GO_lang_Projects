package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r*http.Request){

	if r.URL.Path != "/hello"{
		http.Error(w,"404 NOT FOUND",http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w,"sorry only get method",http.StatusNotFound)
	}
	fmt.Fprintf(w,"Hello")
}
func formHandler(w http.ResponseWriter, r*http.Request){

	if r.URL.Path != "/hello"{
		http.Error(w,"404 NOT FOUND",http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w,"sorry only get method",http.StatusNotFound)
	}
	fmt.Fprintf(w,"Hello")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("server starting .....")

	if err := http.ListenAndServe(":3000",nil);err != nil{
		log.Fatal(err)
	}
}