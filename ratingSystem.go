// https://youtu.be/a3YRzhyy3VY
package main

import (
	"bufio"
	"fmt"

	//"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
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

	//Back end
	fmt.Printf("\n Thanks for rating %v with %v star rating.\n\n Your rating was recorded in our system at %v", food, mynumRating, time.Now().Format(time.Stamp))

	//attempting to write to a file
	/*
		file, err := os.Create("./data.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		file.WriteString("Writing a single line. \n")

		linesToWrite := []string{"Writing multiple lines", "one after the other", "using a string slice."}

		for _, line := range linesToWrite {
			file.WriteString(line + "\n")
		}

		data := []byte("Writing binary data\n")
		file.Write(data)
	*/

}
