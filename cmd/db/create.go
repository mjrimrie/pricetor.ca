package main

import (
	"flag"
	"fmt"
	"github.com/mjrimrie/priceator/internal/datalayer"
	"os"
)


func validateParams(user *string, password *string, host *string, port *int) bool {
	valid := true
	if *user == "" {
		fmt.Println("user cannot be blank")
		valid = false
	}
	if *password == "" {
		fmt.Println("password cannot be blank")
		valid = false
	}
	if *host == "" {
		fmt.Println("host cannot be blank")
		valid = false
	}
	return valid
}

func main() {
	user := flag.String("user", "", "database server user with db create permission")
	password := flag.String("password", "", "password of database user")
	host := flag.String("host", "", "host of the database")
	port := flag.Int("port", 5432, "database server port")

	if len(os.Args) < 5 {
		fmt.Println("expecting all the required flag")
		os.Exit(1)
	}
	flag.Parse()
	paramsValid := validateParams(user, password, host, port)
	if !paramsValid {
		os.Exit(1)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		*host, *port, *user, *password)

	err := datalayer.Initialize(&psqlInfo)
	if err != nil {
		os.Exit(1)
	}


}
