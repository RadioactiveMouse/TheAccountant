package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/csv"
	"strconv"
	"strings"
)

type Record struct {
	Date string
	SortCode string
	AccountNumber string
	Value float64
	Type string
	Reference string
}

func (r Record) String() string {
	return fmt.Sprintf("Date : %v\nSort-Code : %+v\nAccount Number : %+v\nValue : £%0.2f\nPaymentType : %+v\nReference : %+v \n",r.Date,r.SortCode,r.AccountNumber,r.Value,r.Type,r.Reference)
}

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
	records := make([]Record, len(data))
	for i, item  := range data {
		_, err := strconv.ParseFloat(item[3], 32)
		if err == nil {
			details := strings.Split(item[2], " ")
			value, err := strconv.ParseFloat(item[3],32)
			if err != nil { log.Println(err) }
			record := Record{
				Date : item[1],
				SortCode : details[0],
				AccountNumber : details[1],
				Value : value,
				Type : item[4],
				Reference : item[5],
			}
			records[i-1] = record
			fmt.Print(record)
		} else {
			log.Println(err)
		}
	}
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
