package main

import (

	"fmt"
	"encoding/json"
)

type User struct {
	Name string 	`json:"name"`
	Age int 		`json:"age,omitzero"`
	Email string 	`json:"email,omitzero"`
	Active bool		`json:"active,omitzero"`
}

func main() {
	
	user := User{
		Name: "omit zero",
		Age: 0, // Zero value
		Email: "", // Zero value
		Active: false, // Zero value
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData)) // Print JSON data correctly

}

