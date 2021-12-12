package main

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	Passenger_DB "PassengerMS/Database"

)

type passengerInfo struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	MobileNumber string `json:"MobileNumber"`
	EmailAddress string `json:"EmailAddress"`
}

// used for storing Passenger on the REST API
var passengers map[string]passengerInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

func allpassengers(w http.ResponseWriter, r *http.Request) {


	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}


	json.NewEncoder(w).Encode(passengers)

}

func passenger(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	params := mux.Vars(r)

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Passenger
		if r.Method == "POST" {

			// read the string sent to the service
			var newPassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newPassenger)

				//Check if all fields are filled up
				if newPassenger.FirstName == "" || newPassenger.LastName == "" || newPassenger.EmailAddress == "" || newPassenger.MobileNumber == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Passenger " +
							"information " + "in JSON format" + " & Provide all required Fields"))
					return
				}

				// check if Passenger exists; add only if
				// Passenger does not exist
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] = newPassenger
					Passenger_DB.InsertRecord(params["passengerid"], newPassenger.FirstName, newPassenger.LastName, newPassenger.MobileNumber, newPassenger.EmailAddress)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger registered: " +
						params["passengerid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Passenger ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information " +
					"in JSON format"))
			}
		}

		//---PUT is for Updating a Passenger's information
		if r.Method == "PUT" {
			var newPassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newPassenger)

				//Check if all fields are filled up
				if newPassenger.FirstName == "" || newPassenger.LastName == "" || newPassenger.EmailAddress == "" || newPassenger.MobileNumber == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Passenger " +
							" information " +
							"in JSON format"))
					return
				}

				// check if Passenger exists; add only if
				// Passenger does not exist
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] =
						newPassenger
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["passengerid"]))
				} else {
					// update Passenger
					passengers[params["passengerid"]] = newPassenger
					Passenger_DB.EditRecord(params["passengerid"], newPassenger.FirstName, newPassenger.LastName, newPassenger.MobileNumber, newPassenger.EmailAddress)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Passenger updated: " +
						params["passengerid"]))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"Passenger information " +
					"in JSON format"))
			}
		}

		
		if r.Method == "DELETE" {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Deleting of Account not allowed due to auditing reasons."))
		}

	}
}

func main() {

	// instantiate Passengers
	passengers = make(map[string]passengerInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/passengers", allpassengers)
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Passenger Rest API")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
