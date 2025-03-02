package main

import (

	"fmt"
	"encoding/json"
)

type User struct {
	Name string 	`json:"name"`
	Age int 	`json:"age"`
	Email string 	`json:"email"`
	Active bool	`json:"active"`
}

// Method of User struct
// func (u User) PrintValue() {
// 	fmt.Println(u)
// }

func main() {
	
	user := User{
		Name: "omit zero",
		Age: 0, // Zero value
		Email: "", // Zero value
		Active: false, // Zero value
	}

	jsonData, _ := json.Marshal(user)
	// user.PrintValue()
	fmt.Println(string(jsonData)) // Print JSON data correctly

}

