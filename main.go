package main

import (
	"fmt"
)

func main(){
	conferenceName := "Go Conf"
	const totalTickets = 200
	remainingTickets := totalTickets

	fmt.Printf("Welcome to %v, booking app :))\n", conferenceName)
	fmt.Printf("Total number of tickets: %v\nRemaining Tickets: %v\n", totalTickets, remainingTickets)

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

	remainingTickets = totalTickets - userTickets

	fmt.Printf("User %v %v, booked %v tickets, remaining tickets: %v\n", firstName, lastName, userTickets, remainingTickets)


}   