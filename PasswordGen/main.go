package main

import (
	"fmt"
	"log"

	db "passwordgen/dbconnection"
	pg "passwordgen/generator"
)

func check(e error) {
	if e != nil {
		log.Printf("Error %s", e)
		return
	}
}

func CalcPassChar()

func main() {
	var length int
	var minSpecChar int
	var minNum int
	var minUpperCase int
	var minLowerCase int
	var webName string
	var URL string
	var uname string
	var pword string
	var dbname string

	fmt.Println("Welcome to the Password Generator!")

	fmt.Print("Length of Your Password -> ")
	fmt.Scan(&length)

	fmt.Print("Number of Special Characters -> ")
	fmt.Scan(&minSpecChar)

	fmt.Print("How many numbers in the password -> ")
	fmt.Scan(&minNum)

	fmt.Print("Number of Upper Case Letters do you need -> ")
	fmt.Scan(&minUpperCase)

	fmt.Print("Number of Lower Case Letters do you need -> ")
	fmt.Scan(&minLowerCase)

	password, err := pg.GeneratePassword(length, minSpecChar, minNum, minUpperCase, minLowerCase)
	check(err)

	fmt.Print("Connect and Create a Database \n")

	fmt.Print("DBName: ")
	fmt.Scan(&dbname)

	newdb, err := db.CreateDB(dbname)
	check(err)

	fmt.Print("What is Name to save for the website? -> ")
	fmt.Scan(&webName)

	fmt.Print("What is the URL for the website? -> ")
	fmt.Scan(&URL)

	fmt.Print("What is the Username? -> ")
	fmt.Scan(&uname)

	pword = password

	fmt.Print("What is the Web Name to Username and Password -> ")
	fmt.Scan(&webName)

	defer newdb.Close()

	newdb, err = db.OpenDB(dbname)
	check(err)

	err = db.CreateTables(newdb)
	check(err)

	err = db.InsertIntoDB(webName, URL, uname, pword, newdb)
	check(err)

	username, password, err := db.SelectingAUser(webName, newdb)
	check(err)

	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %v\n", password)

	// fmt.Print("What Account do you not use anymore -> ")
	// fmt.Scan(&webName)

	// err = db.DeleteUserInfo(webName, newdb)
	// if err != nil {
	// 	log.Printf("Error %s", err)
	// 	return
	// }

	err = db.UpdateIntoDB(webName, password, newdb)
	check(err)

}


