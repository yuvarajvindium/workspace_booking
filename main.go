package main

import (
	"database/sql"
	"fmt"
	"log"
	"workspace_booking/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
)

func main() {

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
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	println(config.GetServerPort())
	log.Fatalln(app.Listen(config.GetServerPort()))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
