package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceticket = 50
	var remainTicket = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of ", conferenceticket, "ticket and", remainTicket, "are still available")
	fmt.Println("Get your ticket here to attend ")
	var firstName string
	var lastName string
	var Email string
	var userTicket int

	// var updateTickets int
	fmt.Println("Enter your First name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email id:")
	fmt.Scan(&Email)
	fmt.Println("Enter No of tickets you want:")
	fmt.Scan(&userTicket)
	remainTicket = remainTicket - userTicket
	fmt.Printf("Thank you  %v %v for booking  %v tickets. You willrecive a confiramtion email at %v \n", firstName, lastName, userTicket, Email)
	fmt.Printf("%v tickets remaining for %v\n", remainTicket, conferenceName)
}
