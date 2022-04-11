package main

import (
	"database/sql"
	"fmt"
	"log"
	"workspace_booking/config"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func main() {

	// connection string
	psqlconn := config.GetDBConnectionURL()

	println(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	app := fiber.New()

	println(config.GetServerPort())
	log.Fatalln(app.Listen(config.GetServerPort()))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
