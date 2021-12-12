package DB

import (
	"database/sql"
	"fmt"
	

	_ "github.com/go-sql-driver/mysql"
)

type Trip struct { // map this type to the record in the table

	TripID        string
	PickUpPostal  string
	DropOffPostal string
	DriverID      string
	PassengerID   string
	DateTime      string
	Status        string
}

//Edit Trip Status & Identification Number
func EditRecord(TripID string, DriverID string,  Status string) {
	
	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/trip_db")

	
	query := fmt.Sprintf(
		"UPDATE Trip SET TripStatus='%s', IdentificationNumber='%s' WHERE TripID='%s'",
		Status, DriverID, TripID)
	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Create a Trip
func InsertRecord(TripID string, PickUpPostal string, DropOffPostal string, DriverID string, PassengerID string, DateTime string, Status string) {

	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/trip_db")

	query := fmt.Sprintf("INSERT INTO Trip (TripID, PickUpPostal, DropOffPostal, IdentificationNumber, PassengerID, TripDateTime, TripStatus) VALUES ('%s', '%s', '%s', '%s', '%s', '%s' , '%s')",
		TripID, PickUpPostal, DropOffPostal, DriverID, PassengerID, DateTime, Status)

	_, err = db.Query(query)

	if err != nil {

		panic(err.Error())
	}

	db.Close()
}


//Get all Open Trips for Driver to View
func GetOpenTrips() (res []Trip){

	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/trip_db")

	//SQL statement to only select Trips that are Oen
	results, err := db.Query("Select * FROM trip_db.Trip WHERE TripStatus='Open'")

	if err != nil {
		panic(err.Error())
	}

	OpenTrips := make([]Trip, 0)
	
	for results.Next() {
		// map this type to the record in the table
		var trip Trip
		err = results.Scan(&trip.TripID, &trip.PickUpPostal,
			&trip.DropOffPostal, &trip.DriverID, &trip.PassengerID, &trip.DateTime, &trip.Status)
		if err != nil {
			panic(err.Error())
		}

		
		fmt.Println(trip.TripID, trip.PickUpPostal,
			trip.DropOffPostal, trip.DriverID, trip.PassengerID, trip.DateTime, trip.Status)
		
			
		//Append each result to List
		OpenTrips = append(OpenTrips, trip)
	}
	
	//Pass back list to RestAPI
	return OpenTrips
	

}

//Get Passenger Trips in Reverse Chronological Order
func GetPassengerTrips(PassengerID string) (res []Trip){

	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/trip_db")

	//SQL statement to get rows in Reverse Chronological Order
	query := fmt.Sprintf("Select * FROM trip_db.Trip WHERE PassengerID='%s' ORDER BY TripDateTime DESC",
		 PassengerID)

	results, err := db.Query(query)

	
	
	if err != nil {
		panic(err.Error())
	}

	PassengerTrips := make([]Trip, 0)
	
	for results.Next() {
		// map this type to the record in the table
		var trip Trip
		err = results.Scan(&trip.TripID, &trip.PickUpPostal,
			&trip.DropOffPostal, &trip.DriverID, &trip.PassengerID, &trip.DateTime, &trip.Status)
		if err != nil {
			panic(err.Error())
		}

		
		fmt.Println(trip.TripID, trip.PickUpPostal,
			trip.DropOffPostal, trip.DriverID, trip.PassengerID, trip.DateTime, trip.Status)
		
		//Append each result to List
		PassengerTrips = append(PassengerTrips, trip)
	}
	
	//Pass back list to RestAPI
	return PassengerTrips
	

}

func main() {
	fmt.Println("Trip DB")
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/trip_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	
	defer db.Close()
}
