# ETI_ASG1
Ride Sharing Platform

Design Considerations:

I designed the architecture in a way that respects the Microservice value of being independent and loosely coupled. When analysing the business requirements, three different domains come out : Driver , Passenger and Trip. Passenger and Driver are standalone MS to create and update accounts while Trip MS is the middleman as it requires information from both ends to function. In the real application, it is planned to have two UIs for the respective Passenger and Driver. Each MS would have it's own REST API for communication and have it's own Databases.




Architecture Diagram:
![ETI ASG I](https://user-images.githubusercontent.com/83932770/145717538-1716a17d-308c-4c50-ad9a-2565c1620fd4.png)

Instructions:

1.  Clone Repositary into your VSCODE
2.  Run the SQL script in MYSQLWORKBENCH to create the tables
3.  If packages are not recognised , run the following in the respective root folders:
   1. go mod init "filename"
   2. go mod tidy
