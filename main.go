package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conf"
	const totalTickets = 200
	remainingTickets := totalTickets
	bookings := []string{}

	for {

		greetUser(conferenceName, totalTickets, remainingTickets)

		firstName, lastName, email, userTickets := input()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookings, remainingTickets = bookingsLogic(bookings, firstName, lastName, remainingTickets, userTickets)

			firstNames := getFirstNames(bookings)
			fmt.Printf("Total bookings %v\nbooking list:%v\n", len(firstNames), firstNames)	

			if remainingTickets == 0 {
				fmt.Println("We are sold out. Come back next year ;]")
				break
			} 

		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short!!")
			} 
			if !isValidEmail {
				fmt.Println("email address is not valid!!")
			}
			if !isValidTicketNumber {
				fmt.Println("ticket number is invalid!!")
			}
		}
	}

}   

func greetUser(conferenceName string, totalTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v, booking app :))\n", conferenceName)
	fmt.Printf("Total number of tickets: %v\nRemaining Tickets: %v\n", totalTickets, remainingTickets)
}

func getFirstNames(bookings []string) ([]string){
	firstNames := []string{}

	for _, booking := range(bookings) {
		nameList := strings.Fields(booking)
		firstNames = append(firstNames, nameList[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool){
	isValidName := len(firstName) >=2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func input() (string, string, string, int){
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

func bookingsLogic(bookings []string, firstName string, lastName string, remainingTickets int, userTickets int) ([]string, int){
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v, booked %v tickets, remaining tickets: %v\n", firstName, lastName, userTickets, remainingTickets)

	return bookings, remainingTickets
}

