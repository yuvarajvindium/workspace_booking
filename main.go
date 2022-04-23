package main

import (
	"context"
	"fmt"
	"log"
	"workspace_booking/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	psqlconn := config.GetDBConnectionURL()
	println(psqlconn)
	db, err := pgxpool.Connect(context.Background(), psqlconn)

	// open database
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping(context.Background())
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
