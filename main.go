package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/csv"
)


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Served Home")
}

func importData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w,"Import endpoint")
}

//my account
func account(w http.ResponseWriter, r *http.Request) {
	log.Print("Account endpoint triggered")
	scanDir() // change this so it returns the data for vanilla printing
}

// scanDir should examine the data directory and add the data to the UI
func scanDir() {
	file , err := os.Open("data.csv")
	if err != nil {
		log.Fatal("Unable to find the data.")
	}
	defer file.Close()
	read := csv.NewReader(file)
	data, err := read.ReadAll()
	if err != nil {
		log.Fatal("Error during csv read")
	}
	log.Print(data)
}

func main() {
	// do some parsing of the main files
	http.HandleFunc("/", home)
	http.HandleFunc("/import", importData)
	http.HandleFunc("/myaccount",account)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
