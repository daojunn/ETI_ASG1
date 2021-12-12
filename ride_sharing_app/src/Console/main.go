package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const PassengerURL = "http://localhost:5000/api/v1/passengers"
const Passengerkey = "2c78afaf-97da-4816-bbee-9ad239abb296"

const TripURL = "http://localhost:6000/api/v1/trips"
const Tripkey = "2c78afaf-97da-4816-bbee-9ad239abb298"

const DriverURL = "http://localhost:7000/api/v1/drivers"
const Driverkey = "2c78afaf-97da-4816-bbee-9ad239abb295"

//POST to create Passenger
func addPassenger(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(PassengerURL+"/"+code+"?key="+Passengerkey,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//PUT to update Passenger
func updatePassenger(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		PassengerURL+"/"+code+"?key="+Passengerkey,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//GET to get passenger trips in reverese chronological order
func getPassengerTrips(code string) {

	url := TripURL
	if code != "" {
		url = TripURL + "/passenger/" + code + "?key=" + Tripkey
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}

}

//Driver get all opened trips
func getTrips(code string) {
	url := TripURL
	if code != "" {
		url = TripURL + "/" + code + "?key=" + Tripkey
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		response.Body.Close()
	}
}

//POST to add trip when passenger booked a trip
func addTrip(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(TripURL+"/"+code+"?key="+Tripkey,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//Update Trip status, driver identification number
func updateTrip(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		TripURL+"/"+code+"?key="+Tripkey,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//POST to create driver
func addDriver(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(DriverURL+"/"+code+"?key="+Driverkey,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//PUT to update driver
func updateDriver(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		DriverURL+"/"+code+"?key="+Driverkey,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func main() {

	var option int
	var id int = 1
	for true {
		fmt.Println("Ride Sharing Application:")
		fmt.Println("")
		fmt.Println("Select User:")
		fmt.Println("Passsenger [1]")
		fmt.Println("Driver [2]")
		fmt.Println("")
		fmt.Print("Enter your option:")
		fmt.Scanln(&option)
		if option == 1 {

			for true {
				fmt.Println("Passenger Menu:")
				fmt.Println("")
				fmt.Println("Select Function:")
				fmt.Println("Create A Passenger Account [1]")
				fmt.Println("Update A Passenger Account [2]")
				fmt.Println("Book A Trip [3]")
				fmt.Println("View Trips [4]")
				fmt.Println("")
				fmt.Print("Enter your option:")
				fmt.Scanln(&option)

				// Create Passenger
				if option == 1 {

					var FirstName string
					var LastName string
					var EmailAddress string
					var MobileNumber string

					fmt.Print("Enter First Name:")
					fmt.Scanln(&FirstName)
					fmt.Println("")
					fmt.Print("Enter Last Name:")
					fmt.Scanln(&LastName)
					fmt.Println("")
					fmt.Print("Enter Mobile Number:")
					fmt.Scanln(&MobileNumber)
					fmt.Println("")
					fmt.Print("Enter Email Address:")
					fmt.Scanln(&EmailAddress)

					jsonData := map[string]string{"FirstName": "", "LastName": "", "EmailAddress": "", "MobileNumber": ""}
					jsonData["FirstName"] = FirstName
					jsonData["LastName"] = LastName
					jsonData["EmailAddress"] = EmailAddress
					jsonData["MobileNumber"] = MobileNumber

					addPassenger(fmt.Sprintf("%v", id), jsonData)

					id += 1

					break

				} else if option == 2 {

					//Update Passenger

					var PassengerId string
					var FirstName string
					var LastName string
					var EmailAddress string
					var MobileNumber string

					// Specifying ID acts as a "Login" Page
					fmt.Print("Enter Passenger ID:")
					fmt.Scanln(&PassengerId)
					fmt.Println("")
					fmt.Print("Enter First Name:")
					fmt.Scanln(&FirstName)
					fmt.Println("")
					fmt.Print("Enter Last Name:")
					fmt.Scanln(&LastName)
					fmt.Println("")
					fmt.Print("Enter Mobile Number:")
					fmt.Scanln(&MobileNumber)
					fmt.Println("")
					fmt.Print("Enter Email Address:")
					fmt.Scanln(&EmailAddress)

					jsonData := map[string]string{"FirstName": "", "LastName": "", "EmailAddress": "", "MobileNumber": ""}
					jsonData["FirstName"] = FirstName
					jsonData["LastName"] = LastName
					jsonData["EmailAddress"] = EmailAddress
					jsonData["MobileNumber"] = MobileNumber

					updatePassenger(PassengerId, jsonData)

					break

				} else if option == 3 {

					//Passenger Book a trip
					var TripID string = fmt.Sprintf("%v", id)
					var PickUpPostal string
					var DropOffPostal string
					var PassengerID string
					var DateTime string = time.Now().Format(time.RFC3339)
					var Status string = "Open"

					fmt.Print("Enter Passenger ID:")
					fmt.Scanln(&PassengerID)
					fmt.Print("Enter Pick Up Postal:")
					fmt.Scanln(&PickUpPostal)
					fmt.Print("Enter Drop Off Postal:")
					fmt.Scanln(&DropOffPostal)

					jsonData := map[string]string{"TripID": "", "PickUpPostal": "", "DropOffPostal": "", "DriverID": "", "PassengerID": "", "DateTime": "", "Status": ""}
					jsonData["TripID"] = TripID
					jsonData["PickUpPostal"] = PickUpPostal
					jsonData["DropOffPostal"] = DropOffPostal
					jsonData["PassengerID"] = PassengerID
					jsonData["DateTime"] = DateTime
					jsonData["Status"] = Status

					addTrip(TripID, jsonData)
					id += 1
					break

				} else if option == 4 {

					//View Passenger Ride History

					var PassengerID string
					fmt.Print("Enter Passenger ID:")
					fmt.Scanln(&PassengerID)
					getPassengerTrips(PassengerID)
					break

				}
			}
		} else if option == 2 {

			for true {
				fmt.Println("Driver Menu:")
				fmt.Println("")
				fmt.Println("Select Function:")
				fmt.Println("Create A Driver Account [1]")
				fmt.Println("Update A Driver Account [2]")
				fmt.Println("Start A Trip [3]")
				fmt.Println("")
				fmt.Print("Enter Option:")
				fmt.Scanln(&option)
				if option == 1 {

					//Create Driver Account

					var IdentificationNumber string
					var FirstName string
					var LastName string
					var EmailAddress string
					var MobileNumber string
					var CarLicenseNumber string

					fmt.Print("Enter Identification Number:")
					fmt.Scanln(&IdentificationNumber)
					fmt.Print("Enter First Name:")
					fmt.Scanln(&FirstName)
					fmt.Print("Enter Last Name:")
					fmt.Scanln(&LastName)
					fmt.Print("Enter Email Address:")
					fmt.Scanln(&EmailAddress)
					fmt.Print("Enter Car License Number:")
					fmt.Scanln(&CarLicenseNumber)
					fmt.Print("Enter Mobile Number:")
					fmt.Scanln(&MobileNumber)

					jsonData := map[string]string{"IdentificationNumber": "", "FirstName": "", "LastName": "", "EmailAddress": "", "MobileNumber": "", "CarLicenseNumber": ""}
					jsonData["IdentificationNumber"] = IdentificationNumber
					jsonData["FirstName"] = FirstName
					jsonData["LastName"] = LastName
					jsonData["EmailAddress"] = EmailAddress
					jsonData["MobileNumber"] = MobileNumber
					jsonData["CarLicenseNumber"] = CarLicenseNumber

					addDriver(fmt.Sprintf("%v", id), jsonData)
					id += 1
					break

				} else if option == 2 {

					//Update Driver Account

					var IdentificationNumber string
					var FirstName string
					var LastName string
					var EmailAddress string
					var MobileNumber string
					var CarLicenseNumber string

					fmt.Println("Update Information:")
					fmt.Println("")
					fmt.Print("Enter Identification Number:")
					fmt.Scanln(&IdentificationNumber)
					fmt.Print("Enter First Name:")
					fmt.Scanln(&FirstName)
					fmt.Print("Enter Last Name:")
					fmt.Scanln(&LastName)
					fmt.Print("Enter Car License Number:")
					fmt.Scanln(&CarLicenseNumber)
					fmt.Print("Enter Mobile Number:")
					fmt.Scanln(&MobileNumber)
					fmt.Print("Enter Email Address:")
					fmt.Scanln(&EmailAddress)

					jsonData := map[string]string{"IdentificationNumber": "", "FirstName": "", "LastName": "", "EmailAddress": "", "MobileNumber": "", "CarLicenseNumber": ""}
					jsonData["IdentificationNumber"] = IdentificationNumber
					jsonData["FirstName"] = FirstName
					jsonData["LastName"] = LastName
					jsonData["EmailAddress"] = EmailAddress
					jsonData["MobileNumber"] = MobileNumber
					jsonData["CarLicenseNumber"] = CarLicenseNumber

					updateDriver(fmt.Sprintf("%v", id), jsonData)
					id += 1
					break

				} else if option == 3 {

					//Driver View Open Trips

					var TripOption int
					var IdentificationNumber string
					var TripID string
					var Status string = "On-Going"

					fmt.Print("Enter your Driver Identification Number:")
					fmt.Scanln(&IdentificationNumber)
					getTrips("")

					fmt.Print("Enter Trip ID of Trip to start the Trip:")
					fmt.Scanln(&TripID)

					jsonData := map[string]string{"DriverID": "", "TripID": "", "Status": ""}
					jsonData["DriverID"] = IdentificationNumber
					jsonData["TripID"] = TripID
					jsonData["Status"] = Status

					// Add Driver to Trip and Change Trip Status to On-going
					updateTrip(TripID, jsonData)

					fmt.Println("Trip Started!")
					fmt.Println("")
					fmt.Print("Type [0] to Complete the Trip:")
					fmt.Scanln(&TripOption)
					var UpdatedStatus string = "Finished"

					//When Driver Completed Trip
					if TripOption == 0 {

						jsonData = map[string]string{"DriverID": "", "TripID": "", "Status": ""}
						jsonData["DriverID"] = IdentificationNumber
						jsonData["TripID"] = TripID
						jsonData["Status"] = UpdatedStatus

						//Change Trip Status to Completed
						updateTrip(TripID, jsonData)
						fmt.Println("Trip Completed!")
						break

					}

				}
			}
		} else {
			fmt.Println("Please select option 1 or 2:")
		}
	}
}
