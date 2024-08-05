package main

import (
	"testServer/db"
	"testServer/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	server := fiber.New()
	routes.Setup(server)
	server.Listen(":" + db.GetPort())
}
