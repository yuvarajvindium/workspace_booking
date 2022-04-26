package db

import (
	"context"
	"fmt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetUsers() []User {
	fmt.Println("Fetching")

	rows, err := dbPool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil
	}

	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var email string
		var id string
		var name string
		var password string

		err = rows.Scan(&name, &id, &email, &password)
		if err != nil {
			fmt.Println("Failed", err)
			return []User{}
		}

		users = append(users, User{ID: id, Name: name, Email: email, Password: password})
	}

	return users
}
