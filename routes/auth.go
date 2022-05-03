package routes

import (
	"workspace_booking/db"
	"workspace_booking/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func setupAuthRoutes(apiApp fiber.Router) {
	// Handles login process / JWT based
	apiApp.Post("/auth/login", func(c *fiber.Ctx) error {
		u := new(db.User)
		reqBody := new(db.User)

		if err := c.BodyParser(reqBody); err != nil {
			return err
		}

		// Check if the password and email values are valid
		if reqBody.Password == "" || reqBody.Email == "" {
			return c.JSON(map[string]string{"message": "Email/Password cannot be empty"})
		}

		// check the user exist in db
		u = u.GetUserByEmail(reqBody.Email)
		if u == nil {
			return c.JSON(map[string]string{"message": "Invalid Email / Password"})
		}

		if !services.ComparePassword(u.Password, reqBody.Password) {
			return c.JSON(map[string]string{"message": "Invalid Email / Password"})
		}

		token, err := services.GenerateJWT(u.ID)
		if err != nil {
			return c.JSON(map[string]string{"message": "Failed to login"})
		}

		return c.JSON(map[string]string{"token": token})
	})

	// Forgot password
	apiApp.Post("/auth/forgot-password", func(c *fiber.Ctx) error {
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
	apiApp.Put("/auth/update-password", func(c *fiber.Ctx) error {
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

}
