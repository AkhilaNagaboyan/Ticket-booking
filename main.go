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
	//calling fuction
	greetusers(conferenceName, conferenceTickets, remainingTickets)
	//for loop
	for {
        firstName,lastName,Email,userTickets := getUserInput()
		isValidName,isValidEmail,isValidTicketNumber := ValidateUserInput(firstname,lastename,userTickets,remainingTickets)

		if isValidName && isVaildEmail && isValidTicketNumber {
			bookingTicket(remainingTickets uint,userTicket uint,bookings []string,firstName string,lastName string,email string,conferenceName string)

			//call function firstnames
			firstNames ;= getFirstName(bookings)
			fmt.Printf("first name of booking are : %v\n",firstName)

			if remainingTickets == 0 {
				//end program
				fmt.Println("our conference is booked out.Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isVaildEmail {
				fmt.Println("email address you entered doesn't  contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}
		}
	}
}

// functions
func greetusers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("we have total of %v Tickets and %v still are available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}
func getFirstName(bookings []string []string) {
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
	return firstNames
}
 fun validateUserInput(firstName string,lastName string,email string,userTickets uint,remainingTickets uint)(bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isVaildEmail := strings.Contains(Email, "@") // gives bool results back
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isvalidName , isvalidEmail,isvalidTicketsNumber
 }
 func getUserInput() (string,string,string,uint){
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
	 
	return firstName ,lastName,Email,userTickets
 }
 func bookingTicket(remainingTickets uint,userTicket uint,bookings []string,firstName string,lastName string,email string,conferenceName string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf(" Thank you  %v %v for booking Ticket . you will receive  a confirmation message\n", firstName, lastName, userTickets, Email)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)
 }