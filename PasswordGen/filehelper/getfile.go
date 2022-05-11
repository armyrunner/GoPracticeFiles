package filehelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetFileInfo(name string) {

	jf, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer jf.Close()

	fmt.Println("Sucessfully Opened file.")

	byteValue, _ := ioutil.ReadAll(jf)

	var users Accounts
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Accounts); i++{
		fmt.Println("Name: " + users.Accounts[i].Name)
		fmt.Println("URL: " + users.Accounts[i].URL)
		fmt.Println("Username: " + users.Accounts[i].Username)
		fmt.Println("Password: " + users.Accounts[i].Password)
	}

}
