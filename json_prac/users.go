package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users array
type Users struct {
	Users []User `json:users`
}

// User structure
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:age`
	Social Social `json:"social"`
}

// Social media handles
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	usersFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println("Failed to open")
	}
	defer usersFile.Close()

	byteValue, _ := ioutil.ReadAll(usersFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

}
