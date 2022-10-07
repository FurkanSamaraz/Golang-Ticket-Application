package api

import (
	"main/model"

	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestTicketOptions(t *testing.T) {
	app := fiber.New()

	app.Post("/ticket_options", func(c *fiber.Ctx) error {
		body := &model.Ticket_options{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/ticket_options")
	rb.Json(&model.Ticket_options{Id: 1, Name: "furkan", Desc: "deneme", Allocation: 100, Create_Date: "2022"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Ticket_options{Id: 1, Name: "furkan", Desc: "deneme", Allocation: 100, Create_Date: "2022"})

}
func TestPurchases(t *testing.T) {
	app := fiber.New()

	app.Post("/Ticket_purchases", func(c *fiber.Ctx) error {
		body := &model.Ticket_purchases{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/Ticket_purchases")
	rb.Json(&model.Ticket_purchases{Id: 1, Quantity: 10, User_id: "1"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Ticket_purchases{Id: 1, Quantity: 10, User_id: "1"})

}
func TestUser(t *testing.T) {
	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		body := &model.Ticket_user{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/users")
	rb.Json(&model.Ticket_user{Id: 1, User_id: "2", Name: "furkan"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Ticket_user{Id: 1, User_id: "2", Name: "furkan"})

}
