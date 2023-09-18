package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name type = expression, Either the type or the = expression can be omitted, but not both.
	var conferenceName = "Go Conference"
	// or conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your ticket here to attend")

	// var bookings = [50]string{"Nina", "Nicole", "Peter"}
	// var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	// for remainingTickets > 0 && len(bookings) < 50 {
	// for true {
	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("THe whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("slice type: %T\n", bookings)
			fmt.Printf("slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			var noTicketssRemaining bool = remainingTickets == 0
			if noTicketssRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}
