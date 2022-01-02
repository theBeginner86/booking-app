package helper

import (
	"fmt"
	"strings"
)

func GreetUser(conferenceName string, totalTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v, booking app :))\n", conferenceName)
	fmt.Printf("Total number of tickets: %v\nRemaining Tickets: %v\n", totalTickets, remainingTickets)
}

func GetFirstNames(bookings []string) ([]string){
	firstNames := []string{}

	for _, booking := range(bookings) {
		nameList := strings.Fields(booking)
		firstNames = append(firstNames, nameList[0])
	}

	return firstNames
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool){
	isValidName := len(firstName) >=2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func Input() (string, string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookingsLogic(bookings []string, firstName string, lastName string, remainingTickets int, userTickets int) ([]string, int){
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v, booked %v tickets, remaining tickets: %v\n", firstName, lastName, userTickets, remainingTickets)

	return bookings, remainingTickets
}
