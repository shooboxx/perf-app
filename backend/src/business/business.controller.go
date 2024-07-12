package business

import "github.com/gofiber/fiber/v2"

type businessController struct{}

func (controller *businessController) getBusiness(c *fiber.Ctx) error {
	return c.SendString("This is a changed business")
}
func (controller *businessController) createBusiness(c *fiber.Ctx) error {
	return c.SendString("This is a changed business")
}

func RegisterRoutes(a fiber.Router) {
	controller := businessController{}
	business := a.Group("/business")
	business.Get("/", controller.getBusiness)
	business.Post("/", controller.createBusiness)
}
