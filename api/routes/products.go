package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest/pkg/entities"
	"go-fiber-rest/pkg/product"
	"net/http"
)

func ProductRouter(app fiber.Router, service product.Service) {
	app.Get("/products", getProducts(service))
	app.Get("/products/:id", getProduct(service))
	app.Post("/products", postProduct(service))
	app.Put("/products", putProduct(service))
	app.Delete("/products/:id", deleteProduct(service))
}

func getProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")
		prod, err := service.FindOne(productID)
		if err != nil {
			return err
		}

		err = c.JSON(prod)
		if err != nil {
			return err
		}

		return nil
	}
}

func deleteProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")
		err := service.Remove(productID)
		if err != nil {
			return err
		}

		c.Status(http.StatusAccepted)
		return nil
	}
}

func putProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result entities.Product
		err := c.BodyParser(&result)

		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		_, dberr := service.Update(&result)
		if dberr != nil {
			return c.SendString(dberr.Error())
		}

		c.Status(http.StatusCreated)
		return nil
	}
}

func postProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result entities.Product
		err := c.BodyParser(&result)

		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		_, dberr := service.Insert(&result)
		if dberr != nil {
			return c.SendString(dberr.Error())
		}

		c.Status(http.StatusCreated)
		return nil
	}
}

func getProducts(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.Fetch()

		if err != nil {
			return err
		}

		return c.JSON(&fetched)
	}
}
