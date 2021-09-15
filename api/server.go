package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	checkoutControllers "github.com/guil95/grpcApi/core/checkout/infra/http/controllers"
	pkg "github.com/guil95/grpcApi/pkg/grpc"
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

	conn := pkg.Conn()
	checkoutControllers.CreateApi(app, file, conn)

	log.Fatal(app.Listen(fmt.Sprint(":8000")))
}
