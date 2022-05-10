package main

import (
	"fmt"
	"os"

	// pg "passwordgen/Generator"
	"passwordgen/filehelper"
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

	name := "./JsonFile/newfile.json"

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

	filehelper.GetFileInfo(name)

	// createPasswordFile(password)

}
