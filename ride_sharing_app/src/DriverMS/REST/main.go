package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	Driver_DB "DriverMS/Database"
)

type driverInfo struct {
	IdentificationNumber string `json:"IdentificationNumber"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	MobileNumber string `json:"MobileNumber"`
	EmailAddress string `json:"EmailAddress"`
	CarLicenseNumber string `json:"CarLicenseNumber"`
}

// used for storing Passenger on the REST API
var drivers map[string]driverInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb295" {
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

func alldrivers(w http.ResponseWriter, r *http.Request) {
	
	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}


	json.NewEncoder(w).Encode(drivers)

}

func driver(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	params := mux.Vars(r)

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Driver
		if r.Method == "POST" {

			// read the string sent to the service
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newDriver)

				//Check if all fields are filled up
				if newDriver.IdentificationNumber=="" || newDriver.FirstName == "" || newDriver.LastName == "" || newDriver.EmailAddress == "" || newDriver.MobileNumber == "" || newDriver.CarLicenseNumber  == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Driver " +
							"information " + "in JSON format" + " & Provide all required Fields"))
					return
				}

				// check if Driver exists; add only if
				// Driver does not exist
				if _, ok := drivers[params["driverid"]]; !ok {
					drivers[params["driverid"]] = newDriver

					//Call Insert Statement of Driver DB File
					Driver_DB.InsertRecord(newDriver.IdentificationNumber, newDriver.FirstName, newDriver.LastName, newDriver.MobileNumber, newDriver.EmailAddress, newDriver.CarLicenseNumber)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver registered: " +
						params["driverid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Driver ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Driver information " +
					"in JSON format"))
			}
		}

		//---PUT is Updating Driver's Information
		if r.Method == "PUT" {
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newDriver)

				//Check if all fields are filled up
				if newDriver.IdentificationNumber=="" || newDriver.FirstName == "" || newDriver.LastName == "" || newDriver.EmailAddress == "" || newDriver.MobileNumber == "" || newDriver.CarLicenseNumber  == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Driver " +
							" information " +
							"in JSON format"))
					return
				}

				
				drivers[params["driverid"]] = newDriver
				//Call Update Statement of Driver DB File
				Driver_DB.EditRecord(newDriver.IdentificationNumber, newDriver.FirstName, newDriver.LastName, newDriver.MobileNumber, newDriver.EmailAddress, newDriver.CarLicenseNumber)
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Driver updated: " +
				params["driverid"]))

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"Driver information " +
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

	// instantiate Drivers
	drivers = make(map[string]driverInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/drivers", alldrivers)
	router.HandleFunc("/api/v1/drivers/{driverid}", driver).Methods("GET", "PUT", "POST", "DELETE")

	fmt.Println("Driver Rest Api")
	fmt.Println("Listening at port 7000")
	log.Fatal(http.ListenAndServe(":7000", router))
}
