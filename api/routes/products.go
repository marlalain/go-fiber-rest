package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest/pkg/entities"
	"go-fiber-rest/pkg/product"
)

func ProductRouter(app fiber.Router, service product.Service) {
	app.Get("/products", getProducts(service))
	app.Post("/products", postProduct(service))
	app.Put("/products", putProduct(service))
	app.Delete("/products", deleteProduct(service))
}

func deleteProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result entities.DeleteRequest
		err := c.BodyParser(&result)
		productID := result.ID

		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		dberr := service.Remove(productID)
		if dberr != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "updated",
		})
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

		update, dberr := service.Update(&result)
		return c.JSON(&fiber.Map{
			"status": update,
			"error":  dberr,
		})
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

		insert, dberr := service.Insert(&result)
		return c.JSON(&fiber.Map{
			"status": insert,
			"error":  dberr,
		})
	}
}

func getProducts(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.Fetch()
		var result fiber.Map

		if err != nil {
			result = fiber.Map{
				"status": false,
				"error":  err.Error(),
			}
		} else {
			result = fiber.Map{
				"status":   true,
				"products": fetched,
			}
		}

		return c.JSON(&result)
	}
}
