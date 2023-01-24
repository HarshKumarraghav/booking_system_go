package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("There are %v tickets and %v are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string
	for {
		if remainingTickets > 0 {
			fmt.Println("Please enter your first name:")
			fmt.Scanln(&firstName)
			fmt.Println("Please enter your last name:")
			fmt.Scanln(&lastName)
			fmt.Println("Please enter your email:")
			fmt.Scanln(&email)
			fmt.Println("Please enter the number of tickets you want to buy")
			fmt.Scanln(&userTickets)
			fmt.Printf("Thank you, %v %v, you have bought %v tickets! you will recieve the confirmation mail at %v \nSee you at the conference. \n", firstName, lastName, userTickets, email)
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			var firstNames []string
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				var firstName = name[0]
				firstNames = append(firstNames, firstName)
				fmt.Printf("firstName %v: \n", firstName)
			}
			fmt.Printf("These are our all the booking: %v \n and %v ticket are remaining.\n", bookings, remainingTickets)

		} else {
			break
		}

	}

}
