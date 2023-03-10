package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type UserData struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   uint
}

var wg = sync.WaitGroup{}

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

// It takes in a first name, last name, email, user tickets and remaining tickets as arguments, and
// then it prints out the first name, last name, email, user tickets and remaining tickets
func bookTickets(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) {
	var userData UserData
	var bookings = make([]UserData, 0)
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
	wg.Add(1)
	go sendTicketViaEmailHandler(firstName, lastName, email, userTickets)
	wg.Wait()
}

// It takes in three parameters, a string, an int, and a uint, and prints out a message to the user
func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("There are %v tickets and %v are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Given a slice of UserData, return a slice of strings containing the first names of each UserData.
func printFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames
}
func sendTicketViaEmailHandler(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("******************************************")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", tickets, email)
	fmt.Println("******************************************")
	wg.Done()
}
