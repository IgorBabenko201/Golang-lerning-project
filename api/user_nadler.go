package api

import (
	"github.com/IgorBabenko201/Golang-lerning-project.git/types"
	"github.com/gofiber/fiber/v2"
)

func HandleListUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the watercolor",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James 2")
}
