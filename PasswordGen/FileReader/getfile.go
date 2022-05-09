package filereader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

)

func getFileInfo(name string) error {

	jf, err := os.Open(name)
	if err != nil {
		fmt.Println("Can't open file or file doesn't exist", err)
	}

	defer jf.Close()

	byteValue, _ := ioutil.ReadAll(jf)

	var users Accounts
	json.Unmarshal([]byte(byteValue), &users)

	for i := 0; i < len(users.Accounts); i++ {
		fmt.Printf("Name: %s", users.Accounts[i].Name)
		fmt.Printf("URL: %s", users.Accounts[i].URL)
		fmt.Printf("Username: %s", users.Accounts[i].Username)
		fmt.Printf("Password: %s", users.Accounts[i].Password)
	}

	return nil

}
