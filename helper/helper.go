package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func GreetUser(conferenceName string, totalTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v, booking app :))\n", conferenceName)
	fmt.Printf("Total number of tickets: %v\nRemaining Tickets: %v\n", totalTickets, remainingTickets)
}

func GetFirstNames(bookings []map[string]string) ([]string){
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
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

func BookingsLogic(bookings []map[string]string, firstName string, lastName string, email string, remainingTickets int, userTickets int) ([]map[string]string, int){
	userDetails := make(map[string]string)

	userDetails["firstName"] = firstName
	userDetails["lastName"] = lastName
	userDetails["email"] = email
	userDetails["userTickets"] = strconv.FormatInt(int64(userTickets), 10)
	bookings = append(bookings, userDetails)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you, %v %v for bookings %v tickets. A confirmation mail would be sent to %v\n", firstName, lastName, userTickets, email)

	return bookings, remainingTickets
}
