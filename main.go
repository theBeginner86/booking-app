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

		fmt.Println("Enter no of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {

			bookings = append(bookings, firstName + " " + lastName)
			remainingTickets = remainingTickets - userTickets
	
			firstNames := []string{}
	
			for _, booking := range(bookings) {
				nameList := strings.Fields(booking)
				firstNames = append(firstNames, nameList[0])
			}
	
			fmt.Printf("User %v %v, booked %v tickets, remaining tickets: %v\n", firstName, lastName, userTickets, remainingTickets)
			fmt.Printf("Total bookings %v\nbooking list:%v\n", len(firstNames), firstNames)
	
			if remainingTickets == 0 {
				fmt.Println("We are sold out. Come back next year ;]")
				break
			} 
		
		} else {
			fmt.Printf("Invalid request. Remaining tickets: %v\n", remainingTickets)
			continue
		}
	}

}   