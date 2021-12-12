package Passenger_DB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct { // map this type to the record in the table

	PassengerID  string
	FirstName    string
	LastName     string
	MobileNumber string
	EmailAddress string
}

//Edit Passenger Information
func EditRecord(PassengerID string, FirstName string, LastName string, MobileNumber string, EmailAddress string) {
	
	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/passenger_db")
	
	query := fmt.Sprintf(
		"UPDATE Passenger SET FirstName='%s', LastName='%s', MobileNumber='%s', EmailAddress='%s' WHERE PassengerID='%s'",
		FirstName, LastName, MobileNumber, EmailAddress, PassengerID)
	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Add Passenger
func InsertRecord(PassengerID string, FirstName string, LastName string, MobileNumber string, EmailAddress string) {
	
	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/passenger_db")
	
	query := fmt.Sprintf("INSERT INTO Passenger (PassengerID, FirstName, LastName, EmailAddress, MobileNumber) VALUES ('%s', '%s', '%s', '%s', '%s')",
		PassengerID, FirstName, LastName, EmailAddress, MobileNumber)

	_, err = db.Query(query)

	if err != nil {

		panic(err.Error())
	}

	defer db.Close()
}


func main() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	defer db.Close()
}
