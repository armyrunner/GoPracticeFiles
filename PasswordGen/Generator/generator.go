package generator

import (
	"math/rand"
	"strings"
	"time"
)
// Google Regex 
var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "~`!@#$%^&*()_-+={[}]|:;'<,>.?/"
	numberSet      = "0123456789"
	allChartSet    = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func GeneratePassword(length, minSpecChar, minNum, minUpperCase, minLowerCase int) (string, error) {

	var password strings.Builder

	//Special Character for Password
	for i := 0; i < minSpecChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Number for Password
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//UpperCase for Password
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//LowerCase for Password
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	//Makeing sure the lenght of the password is at correct length
	currentPassLength := length - minSpecChar - minNum - minUpperCase - minLowerCase
	for i := 0; i < currentPassLength; i++ {
		random := rand.Intn(len(allChartSet))
		password.WriteString(string(allChartSet[random]))
	}

	//Get full password and shuffle it to make a good password
	rand.Seed(time.Now().Unix())
	newpassword := []rune(password.String())
	rand.Shuffle(len(newpassword), func(i, j int) {
		newpassword[i], newpassword[j] = newpassword[j], newpassword[i]
	})

	return string(newpassword), nil

}
