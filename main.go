package main

import (
	"kasir/database"
	"kasir/middleware"
	"kasir/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setUpRouter(app *fiber.App) {
	//User Endpoint
	app.Post("/users", routes.CreateUser)
	app.Get("/users", middleware.AuthMiddleware, routes.GetAllUser)
	app.Get("/users/:id", routes.GetUserByID)
	app.Put("/users/:id", routes.UpdateUserByID)
	app.Delete("/users/:id", routes.DeleteUser)

	//Login Endpoint
	app.Post("/login", routes.Login)

	//Product Endpoint
	app.Post("/product", routes.CreateProduct)
	app.Get("/product", routes.GetAllProduct)
	app.Get("/product/:id", routes.GetProductByID)
	app.Put("/product/:id", routes.UpdateProductByID)
	app.Delete("/users/:id", routes.DeleteProduct)

	//Order Endpoint
	app.Post("/order", routes.CreateOrder)
	app.Get("/order", routes.GetAllOrders)
	app.Get("/order/:id", routes.GetOrderByID)

}

func main() {
	database.ConnectDB()
	app := fiber.New()

	setUpRouter(app)

	log.Fatal(app.Listen(":3000"))
}
