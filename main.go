package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("\nWelcome to '%v' booking application, \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var firstName = strings.Fields(booking)[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("These are all our bookings : %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("\nSorry, our conference is booked out.")
				break
			}

		} else {
			fmt.Printf("\nSorry, we only have %v tickets remaining, so you cant book for %v tickets\n\n", remainingTickets, userTickets)
		}

	}

}
