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

	username := "Jake"
	noOfOrderedTickets := 2

	remainingTickets = totalTickets - noOfOrderedTickets

	fmt.Printf("User %v, booked %v tickets, remaining tickets: %v\n", username, noOfOrderedTickets, remainingTickets)


}   