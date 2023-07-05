package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	firstNames := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining Tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		var firstName string
		var userTickets uint
		var lastName string
		var email string

		fmt.Println("Enter your first Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets

			// var bookings []string
			bookings := []string{}
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("The Slice type %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("The first names of booking are:%v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == remainingTickets {
			//do something else
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//continues to next iteration of loop
		}
	}
}
