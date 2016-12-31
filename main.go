package main

import (
	"fmt"
	"encoding/json"
	"github.com/packageman/seed/helpers"
	"github.com/packageman/seed/extensions"
	"github.com/packageman/seed/models"
)

func main() {
	fmt.Printf(helpers.Reverse("\n!oG, olleH"))

	_, body, errs := extensions.NewRestClient().Get("https://jsonplaceholder.typicode.com/users").End()
	if (errs != nil) {
	    fmt.Println(errs)
	}

	var users []models.User
	err := json.Unmarshal([]byte(body), &users)
    if (err != nil) {
    	fmt.Println(err)
    }

    fmt.Printf("users' type: %T\nusers' length: %d\n", users, len(users))

    user := users[0];
    fmt.Printf("Company's type: %T\n", user.Company)
	
	for _, user := range users {
		fmt.Printf("%d\n%s\n%s\n%s\n==========\n", 
			user.Id, 
			user.Name, 
			user.Email, 
			user.Company.(map[string]interface{})["name"])
	}
}
