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

/*
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set headers
        (w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
        }
        fmt.Println("ok")
        next.ServeHTTP(w, r)
		return
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400) // bad request
		return
    }
    // get the body of the POST request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading post body: %v", err)
		w.WriteHeader(500) // internal server error
		return
    }
    // try to parse the body
	var login Login
	if err = json.Unmarshal(reqBody, &login); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(400) // bad request
		return
    }
    // Do the login comparison
	ok, err := validateCreds(login)
	if err != nil {
		log.Printf("Login internal server error: %v", err)
		w.WriteHeader(500) // Internal server error
		return
	}
	if !ok {
		log.Printf("Invalid username or password")
		w.WriteHeader(401) // unauthorized
    }

    w.WriteHeader(200)
}
*/

func main() {
	/*
			r := mux.NewRouter()
		    r.HandleFunc("/login", loginHandler)
		    r.handle("/", r)
		    log.Fatal(http.ListenAndServe(":10000", nil))*/

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
