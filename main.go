package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("\nWelcome to '%v' booking application, \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend \n")

	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Print("What is your first name? ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name? ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email address? ")
	fmt.Scan(&userEmail)

	fmt.Print("How many tickets you would like to buy? ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", 
			firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}