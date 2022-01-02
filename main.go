package main

import (
	"fmt"
	"booking-app/helper"
)

func main(){
	conferenceName := "Go Conf"
	const totalTickets = 200
	remainingTickets := totalTickets
	bookings := []string{}

	for {

		helper.GreetUser(conferenceName, totalTickets, remainingTickets)

		firstName, lastName, email, userTickets := helper.Input()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookings, remainingTickets = helper.BookingsLogic(bookings, firstName, lastName, remainingTickets, userTickets)

			firstNames := helper.GetFirstNames(bookings)
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

