package main

import (
	DB "TripMS/Database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type tripInfo struct {
	TripID        string `json:"TripID"`
	PickUpPostal  string `json:"PickUpPostal"`
	DropOffPostal string `json:"DropOffPostal"`
	DriverID      string `json:"DriverID"`
	PassengerID   string `json:"PassengerID"`
	DateTime      string `json:"DateTime"`
	Status        string `json:"Status"`
}

// used for storing Trips on the REST API
var trips map[string]tripInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb298" {
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

//Driver View Trips in that are opened
func alltrips(w http.ResponseWriter, r *http.Request) {
	

	OpenTrips := DB.GetOpenTrips()
	fmt.Println(OpenTrips)
	response := json.NewEncoder(w)
	response.SetIndent("", "  ")
	response.Encode(OpenTrips)



}

//Passenger View Trips in reverse chronological order
func allpassengertrips(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	PassengerTrips := DB.GetPassengerTrips(params["passengerid"])
	fmt.Println(PassengerTrips)
	response := json.NewEncoder(w)
	response.SetIndent("", "  ")
	response.Encode(PassengerTrips)

}

func trip(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	params := mux.Vars(r)

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Trip
		if r.Method == "POST" {

			// read the string sent to the service
			var newTrip tripInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newTrip)


				// Check if required Fields filled up
				if newTrip.TripID == "" || newTrip.DropOffPostal == "" || newTrip.PickUpPostal == "" || newTrip.Status == "" || newTrip.DateTime == "" || newTrip.PassengerID == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Trip " +
							"information " + "in JSON format" + " & Provide all required Fields"))
					return
				}

				// check if Trip exists; add only if
				// Trip does not exist
				if _, ok := trips[params["tripid"]]; !ok {
					trips[params["tripid"]] = newTrip

					DB.InsertRecord(newTrip.TripID, newTrip.PickUpPostal, newTrip.DropOffPostal, newTrip.DriverID, newTrip.PassengerID, newTrip.DateTime, newTrip.Status)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Trip Created , Awaiting Driver "))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Trip ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Trip information " +
					"in JSON format"))
			}
		}

		//---PUT is for Updating Trip
		if r.Method == "PUT" {
			var newTrip tripInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newTrip)

				
				// Check if required Fields filled up
				if newTrip.TripID == "" || newTrip.Status == "" || newTrip.DriverID == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Trip " +
							" information " +
							"in JSON format"))
					return
				}

				//Display Messages based on Trip Status
				if newTrip.Status == "On-Going" {
					w.Write([]byte("202 - Trip Started: " +
						params["tripid"]))
				} else if newTrip.Status == "Finished" {
					w.Write([]byte("202 - Trip Finished: " +
						params["tripid"]))
				}

				DB.EditRecord(newTrip.TripID, newTrip.DriverID, newTrip.Status)

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"Trip information " +
					"in JSON format"))
			}
		}

		if r.Method == "GET" {
			if _, ok := trips[params["tripid"]]; ok {
				json.NewEncoder(w).Encode(
					trips[params["tripid"]])

			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - No Trips found"))
			}
		}

		if r.Method == "DELETE" {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Deleting of Trip not allowed due to auditing reasons."))
		
		}

	}
}

func main() {

	// instantiate Trips
	trips = make(map[string]tripInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/trips", alltrips)
	router.HandleFunc("/api/v1/trips/{tripid}", trip).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/api/v1/trips/passenger/{passengerid}", allpassengertrips).Methods("GET")

	fmt.Println("Trip REST API")
	fmt.Println("Listening at port 6000")
	log.Fatal(http.ListenAndServe(":6000", router))
}
