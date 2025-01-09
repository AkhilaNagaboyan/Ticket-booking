package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome  to %v Ticket booking\n", conferenceName)
	fmt.Printf("we have total of %v Tickets and %v still are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

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

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The Array type: %T\n", bookings)
	fmt.Printf("The Array lenght: %v\n", len(bookings))

	fmt.Printf("Thank you  %v %v for booking Ticket . you will receive  a confirmation message\n", firstName, lastName, userTickets, Email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}