package server

import (
	"encoding/json"
	"fmt"
	"os"
)

type Db = []User

func initDb() Db {
	usersData, err := os.ReadFile("/data/users.json")
	if err != nil {
		fmt.Println("'users.json' does not exist. Creating slice.")
		return makeUsers()
	}

	var users []User
	err = json.Unmarshal([]byte(usersData), &users)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return []User{}
	}
	return users
}

func makeUsers() []User {
	users := make([]User, 0, 5)
	for i := range 5 {
		nextUser := User{
			Guid:      fmt.Sprintf("dummy%d", (i+1)*100),
			FirstName: "User",
			LastName:  fmt.Sprintf("%d", i),
			Age:       (i * 3) + 10,
			Admin:     (i%2 == 0),
		}
		users = append(users, nextUser)
	}
	return users
}
