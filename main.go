package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Date          string
	SortCode      string
	AccountNumber string
	Value         float64
	Type          string
	Reference     string
}

func (r Record) String() string {
	return fmt.Sprintf("Date : %v\nSort-Code : %v\nAccount Number : %v\nValue : £%0.2f\nPaymentType : %v\nReference : %v\n", r.Date, r.SortCode, r.AccountNumber, r.Value, r.Type, r.Reference)
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Served Home")
}

func importData(w http.ResponseWriter, r *http.Request) {
	log.Println("Import endpoint")
}

//endpoint to export data out of the system
func export(w http.ResponseWriter, r *http.Request) {
	log.Print("Export endpoint triggered")
	res, err := scanFile("data.csv")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, res)
	}
}

// endpoint to show the x last transactions
func latestTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Latest transactions endpoint")
	fmt.Fprintf(w,"Latest Transactions")
	fmt.Fprintf(w,"#########")
	// loop through the last 5 transactions
}

// scanFile should examine the file and add the data to the UI
func scanFile(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to find the data.")
	}
	defer file.Close()
	read := csv.NewReader(file)
	data, err := read.ReadAll()
	if err != nil {
		log.Fatal("Error during csv read")
	}
	records := make([]Record, len(data)-1)
	for i, item := range data {
		if i != 0 {
			details := strings.Split(item[2], " ")
			value, err := strconv.ParseFloat(item[3], 32)
			if err != nil {
				return nil, errors.New("Error during parsing.")
			} else {
				record := Record{
					Date:          item[1],
					SortCode:      details[0],
					AccountNumber: details[1],
					Value:         value,
					Type:          item[4],
					Reference:     item[5],
				}
				records[i-1] = record
			}
		}
	}
	return records, nil
}

func main() {
	// do some parsing of the main files
	http.HandleFunc("/", home)
	http.HandleFunc("/import", importData)
	http.HandleFunc("/export", export)
	http.HandleFunc("/latest", latestTransaction)

	log.Println("The Accountant started on localhost:8000/")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
