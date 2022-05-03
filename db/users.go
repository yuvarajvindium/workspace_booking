package db

import (
	"context"
	"fmt"
	"workspace_booking/services"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	password, err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = password

	sql := `INSERT INTO users("id", "name", "email", "password") VALUES($1, $2, $3, $4)`
	d, err := dbPool.Exec(context.Background(), sql, u.ID, u.Name, u.Email, u.Password)
	fmt.Println(d)
	return err
}

func (u *User) Update(id string) error {
	sqlQueryParts := "SET "
	i := 1
	values := make([]interface{}, 0)

	if u.Password != "" {
		password, err := services.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = password
		sqlQueryParts += fmt.Sprintf(" password=$%d,", i)
		i++
		values = append(values, u.Password)
	}

	if u.Email != "" {
		sqlQueryParts += fmt.Sprintf(" email = $%d,", i)
		i++
		values = append(values, u.Email)
	}

	if u.Name != "" {
		sqlQueryParts += fmt.Sprintf(" name=$%d", i)
		i++
		values = append(values, u.Name)
	}

	sql := fmt.Sprintf("UPDATE users %s WHERE id='%s'", sqlQueryParts, id)
	fmt.Println(sql)
	d, err := dbPool.Exec(context.Background(), sql, values...)
	fmt.Println(d)
	return err
}

func (u *User) Delete(id string) error {
	sql := "DELETE FROM users WHERE id=$1"
	d, err := dbPool.Exec(context.Background(), sql, id)
	fmt.Println(d)
	return err
}

func (u *User) GetUserByEmail(email string) *User {
	var id string
	var name string
	var password string

	row := dbPool.QueryRow(context.Background(), "SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&name, &id, &email, &password)

	if err != nil {
		return nil
	}

	u.ID = id
	u.Name = name
	u.Email = email
	u.Password = password

	return u
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
