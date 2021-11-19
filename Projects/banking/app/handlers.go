package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Customer!!")
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]interface{})
	timezones := r.URL.Query().Get("tz")
	if timezones == "" {
		res["current_time"] = time.Now().Local().String()
		json.NewEncoder(w).Encode(res)
	} else {
		tzs := strings.Split(timezones, ",")
		for _, tz := range tzs {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezones %s in input", tz)))
				return
			}
			res[tz] = time.Now().In(loc)
		}
		json.NewEncoder(w).Encode(res)
	}

}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Augustine", "Seoul", "103456"},
		{"David", "Hanoi", "1063566"},
		{"Juth", "London", "143436"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "creating...")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, id)
}
