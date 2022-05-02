package routes

import (
	"fmt"
	"workspace_booking/db"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func initUserRoutes(apiApp fiber.Router) {
	// Get the list of users // search
	apiApp.Get("/users", func(c *fiber.Ctx) error {
		u := db.User{}
		users := u.GetUsers()
		fmt.Println(users)
		return c.JSON(users)
	})

	// Create new user
	apiApp.Post("/users", func(c *fiber.Ctx) error {
		u := new(db.User)
		if err := c.BodyParser(u); err != nil {
			return err
		}

		u.ID = uuid.New().String()
		if err := u.Save(); err != nil {
			return err
		}

		return c.JSON(u)
	})

	// Update user by Id
	apiApp.Put("/users/:id", func(c *fiber.Ctx) error {
		u := new(db.User)
		id := c.Params("id")

		if err := c.BodyParser(u); err != nil {
			return err
		}

		if err := u.Update(id); err != nil {
			return err
		}

		return c.JSON(u)
	})

	// Delete user by Id
	apiApp.Delete("/users/:id", func(c *fiber.Ctx) error {
		u := new(db.User)
		id := c.Params("id")
		if err := u.Delete(id); err != nil {
			return err
		}

		res := map[string]string{
			"message": "user with ID" + id + "has been deleted successfully",
		}

		return c.JSON(res)
	})

}
