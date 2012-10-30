package main

import (
	"fmt"
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve Home")
}

func import(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w,"Import endpoint")	
}

func main() {
	
	http.HandleFunc("/", home)
	http.HandleFunc("/import", import)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
