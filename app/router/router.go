package router

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/moaabid/golang-bitly/database"
	"github.com/moaabid/golang-bitly/model"
	"github.com/moaabid/golang-bitly/utils"
)

func SetupRouter() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/r/:redirect", redirect)
	app.Get("/getallbitlies", getAllBitly)
	app.Get("/getbitly/:id", getBitly)
	app.Post("/createbitly", createBitly)
	app.Patch("/updatebitly", updateBitly)
	app.Delete("/deletebitly/:id", deletebitly)

	app.Listen(":3000")

}

func redirect(c *fiber.Ctx) error {
	bitlyURL := c.Params("redirect")

	bitly, err := model.FindByBitlyUrl(bitlyURL)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not bitly" + err.Error(),
		})
	}

	bitly.Clicked += 1
	err = model.UpdateBitly(bitly)
	if err != nil {
		fmt.Printf("error update bitly %v\n", err)
	}
	return c.Redirect(bitly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllBitly(c *fiber.Ctx) error {
	bitlies, err := model.GetAllBitlies()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all bitly links" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(bitlies)
}

func getBitly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error could not parse id" + err.Error(),
		})
	}

	bitly, err := model.GetBilty(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on getting bitly url" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(bitly)

}

func createBitly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var bitly database.Bitly

	err := c.BodyParser(&bitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on creating bitly url" + err.Error(),
		})
	}

	if bitly.Random {
		bitly.Bitly = utils.RandomURL(8)
	}

	err = model.CreateBitly(bitly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create bitly" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(bitly)

}
func updateBitly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var bitly database.Bitly

	err := c.BodyParser(&bitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on updating bitly url" + err.Error(),
		})
	}

	if bitly.Random {
		bitly.Bitly = utils.RandomURL(8)
	}

	err = model.UpdateBitly(bitly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update bitly" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(bitly)

}

func deletebitly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error could not parse id" + err.Error(),
		})
	}

	err = model.DeleteBitly(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on getting bitly url" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "bitly deleted successfully"})

}
