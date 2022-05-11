package main

import (
	"fmt"
	"os"

	// pg "passwordgen/generator"
	db "passwordgen/dbconnection"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createPasswordFile(password string) {
	f, err := os.Create("passfile")
	check(err)

	defer f.Close()

	np := []byte(password)
	_, err = f.Write(np)
	check(err)

}

func main() {
	// var length int
	// var minSpecChar int
	// var minNum int
	// var minUpperCase int
	// var minLowerCase int
	// var name string

	var host string
	var port int
	var user string
	var password string
	var dbname string

	// fmt.Println("Welcome to the Password Generator!")

	// fmt.Print("Length of Your Password -> ")
	// fmt.Scan(&length)

	// fmt.Print("Number of Special Characters -> ")
	// fmt.Scan(&minSpecChar)

	// fmt.Print("How many numbers in the password -> ")
	// fmt.Scan(&minNum)

	// fmt.Print("Number of Upper Case Letters do you need -> ")
	// fmt.Scan(&minUpperCase)

	// fmt.Print("Number of Lower Case Letters do you need -> ")
	// fmt.Scan(&minLowerCase)

	// password, err := pg.GeneratePassword(length, minSpecChar, minNum, minUpperCase, minLowerCase)
	// check(err)

	// fmt.Print("Where is your json file located -> ")
	// fmt.Scan(&name)

	fmt.Print("Connect to the database \n")
	fmt.Print("Host: ")
	fmt.Scan(&host)
	fmt.Print("Port: ")
	fmt.Scan(&port)
	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	fmt.Print("DBName: ")
	fmt.Scan(&dbname)

	constr, err := db.PostgresConnStr(host, port, user, password, dbname)
	check(err)
	fmt.Print(constr)

	// createPasswordFile(password)

}
