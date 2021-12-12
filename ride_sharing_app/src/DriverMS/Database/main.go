package Driver_DB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct { // map this type to the record in the table

	IdentificationNumber string
	FirstName    string
	LastName     string
	MobileNumber string
	EmailAddress string
	CarLicenseNumber string

}

//Update Driver Information
func EditRecord(IdentificationNumber string, FirstName string, LastName string, MobileNumber string, EmailAddress string, CarLicenseNumber string) {
	
	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/driver_db")
	
	//Update all fields except identification number
	query := fmt.Sprintf(
		"UPDATE Driver SET FirstName='%s', LastName='%s', MobileNumber='%s', EmailAddress='%s', CarLicenseNumber='%s' WHERE IdentificationNumber='%s'",
		FirstName, LastName, MobileNumber, EmailAddress, CarLicenseNumber ,IdentificationNumber)
	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}


//Creation of Driver
func InsertRecord(IdentificationNumber string, FirstName string, LastName string, MobileNumber string, EmailAddress string, CarLicenseNumber string) {
	
	// Open DB
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/driver_db")

	//Insert into Driver Table
	query := fmt.Sprintf("INSERT INTO Driver (IdentificationNumber, FirstName, LastName, EmailAddress, MobileNumber, CarLicenseNumber) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		IdentificationNumber, FirstName, LastName, EmailAddress, MobileNumber, CarLicenseNumber)

	_, err = db.Query(query)

	if err != nil {

		panic(err.Error())
	}
}



func main() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}


	// defer the close till after the main function has finished executing
	defer db.Close()
}
