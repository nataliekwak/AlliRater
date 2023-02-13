// https://youtu.be/a3YRzhyy3VY
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var choice string
	fmt.Println("Would you like to register[button] or login[button]?")
	fmt.Scanln(&choice)

	if choice == "register" {
		register()
	} else if choice == "login" {
		valid := false

		var username string
		fmt.Println("Enter username: ")
		fmt.Scanln(&username)

		var password string
		fmt.Println("Enter password: ")
		fmt.Scanln(&password)

		file, err := os.Open("login.txt")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		defer file.Close()

		var temp string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			temp = scanner.Text()

			if valid {
				if temp == password {
					fmt.Println("Login successful!")
					break
				} else {
					fmt.Println("Wrong password.")
					os.Exit(0) //added to exit program before asking to rate meal
					break
				}
			}
			if strings.Compare(temp, username) == 0 {
				valid = true
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("File reading error", err)
		}

		if !valid {
			fmt.Println("Username does not exist.")
			os.Exit(0)
		}

		rating()
	} else {
		fmt.Println("Please type one of the choices above. Try again.")
	}
}

func register() {
	var user string
	var pass string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter login")
	user, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter password")
	pass, _ = reader.ReadString('\n')

	fmt.Println("Thanks for registering as a user with AlliRater")

	file, err := os.OpenFile("./login.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(user)
	file.WriteString(pass)

}

func rating() {
	var food string
	var userRating string

	//Front end
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What meal would you like to rate?: ")
	food, _ = reader.ReadString('\n') //food

	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please rate the %v between 1 and 5: ", food)
	//userRating = string? mynumRating = float
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("\n Thanks for rating %v with %v star rating.\n\n Your rating was recorded in our system at %v", food, mynumRating, time.Now().Format(time.Stamp))

	//writing ratings to data.txt
	file, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("Meal: ")
	file.WriteString(food)
	file.WriteString("Rating: ")
	file.WriteString(userRating)
	file.WriteString("~")
}
