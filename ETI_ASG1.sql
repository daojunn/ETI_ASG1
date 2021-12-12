CREATE USER 'user2'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user2'@'localhost';

CREATE database passenger_db;

USE passenger_db;

CREATE TABLE Passenger (PassengerID INT NOT NULL PRIMARY KEY, FirstName VARCHAR(30) NOT NULL, LastName VARCHAR(30) NOT NULL, EmailAddress VARCHAR(30) NOT NULL, MobileNumber VARCHAR(8) NOT NULL); 


CREATE database driver_db;

USE driver_db;

CREATE TABLE driver (IdentificationNumber VARCHAR(20) NOT NULL PRIMARY KEY, FirstName VARCHAR(30) NOT NULL, LastName VARCHAR(30) NOT NULL, EmailAddress VARCHAR(30) NOT NULL, MobileNumber VARCHAR(8) NOT NULL, CarLicenseNumber VARCHAR(20) NOT NULL); 

CREATE database trip_db;

use trip_db;

CREATE TABLE Trip (TripID  INT NOT NULL PRIMARY KEY, PickUpPostal VARCHAR(6)  NOT NULL, DropOffPostal  VARCHAR(6)  NOT NULL, IdentificationNumber VARCHAR(20), PassengerID VARCHAR(10), TripDateTime datetime, TripStatus VARCHAR(20)); 


