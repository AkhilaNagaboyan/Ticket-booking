package main

import (
	"fmt"
	"strings"
)

func main() {
	//variables
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var booking []string
	bookings := []string{}
	// Greeting Messages
	fmt.Printf("Welcome  to %v Ticket booking\n", conferenceName)
	fmt.Printf("we have total of %v Tickets and %v still are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
	//for loop
	for {
		//User Input array
		var firstName string
		var lastName string
		var Email string
		var userTickets uint
		//ask user for their name
		fmt.Printf("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your Email:")
		fmt.Scan(&Email)

		fmt.Printf("Enter your Number tickets of users:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			// Booking Logic
			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			//slice
			bookings = append(bookings, firstName+" "+lastName)
			// Debugging Information
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The slice lenght: %v\n", len(bookings))

			fmt.Printf(" Thank you  %v %v for booking Ticket . you will receive  a confirmation message\n", firstName, lastName, userTickets, Email)
			fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)
			//for each loop
			firstNames := []string{}
			//iteration
			for _, booking := range bookings {
				//logic
				var names = strings.Fields(booking)
				//var firstName = names[0]
				firstNames = append(firstNames, names[0])

				//fmt.Printf("Thank you  %v %v for booking Ticket . you will receive  a confirmation message\n", firstName, lastName, userTickets, Email)
				//fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			}
			//slice usecase
			fmt.Printf("These first names of bookings are: %v\n ", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("our conference is booked out.Come back next year.")
				break
			}
		} else {
			fmt.Printf("we have only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		}
	}
}
