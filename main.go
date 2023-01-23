package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("There are %v tickets and %v are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend", conferenceTickets)
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string
	for {
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
		fmt.Printf("These are our all the booking: %v \n and %v ticket are remaining.\n", bookings, remainingTickets)
	}

}
