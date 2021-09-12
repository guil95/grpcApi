package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	controllers "github.com/guil95/grpcApi/core/infra/http/controllers"
	"log"
	"net/http"
	"os"
)

func Run(file []byte) {
	port := os.Getenv("API_PORT")

	log.Println("Listen server on "+ port)

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		err := context.SendStatus(http.StatusOK)
		if err != nil {
			return nil
		}

		return context.JSON(fiber.Map{"message": "Welcome to hash discount api"})
	})

	controllers.CreateApi(app, file)

	log.Fatal(app.Listen(fmt.Sprint(":8000")))
}
