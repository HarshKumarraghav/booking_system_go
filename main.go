package main

import (
	"fmt"
	"strings"
)

type UserData struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   uint
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	greetUser(conferenceName, conferenceTickets, remainingTickets)
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	bookTickets(firstName, lastName, email, userTickets, remainingTickets)
}
func bookTickets(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) {
	var userData UserData
	var bookings = make([]UserData, 0)
	for {

		fmt.Println("Please enter your first name:")
		fmt.Scanln(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scanln(&lastName)
		fmt.Println("Please enter your email:")
		fmt.Scanln(&email)
		fmt.Println("Please enter the number of tickets you want to buy")
		fmt.Scanln(&userTickets)
		userData.FirstName = firstName
		userData.LastName = lastName
		userData.Email = email
		userData.Tickets = userTickets
		fmt.Printf("userData: %v \n", userData)
		fmt.Printf("Thank you, %v %v, you have bought %v tickets! you will recieve the confirmation mail at %v \nSee you at the conference. \n", firstName, lastName, userTickets, email)
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets
		if isValidEmail && isValidName && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userData)

			firstName := printFirstNames(bookings)
			fmt.Printf("firstName: %v \n", firstName)

		} else {
			fmt.Println("Your input value are invalid try again")
		}

	}
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("There are %v tickets and %v are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func printFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames
}
