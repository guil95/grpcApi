package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guil95/grpcApi/core/domain"
	"github.com/guil95/grpcApi/core/infra/http/clients"
	"github.com/guil95/grpcApi/core/infra/repositories"
	application "github.com/guil95/grpcApi/core/use_cases"
	"log"
	"net/http"
)

func CreateApi(app *fiber.App, file []byte) {
	app.Post("/checkout", func(ctx *fiber.Ctx) error {
		checkout(ctx, application.NewService(clients.NewDiscountGrpcClient(), repositories.NewFileRepository(file)))
		return nil
	})
}

func checkout(ctx *fiber.Ctx, service *application.Service) {
	var chartPayload = new(domain.Chart)

	if err := ctx.BodyParser(chartPayload); err != nil {
		log.Println(err)

		err := ctx.Status(http.StatusUnprocessableEntity).JSON(NewResponseError("Unprocessable entity"))

		if err != nil {
			return
		}

		return
	}

	order, err := service.Checkout(chartPayload)

	if err != nil {
		if err == domain.ProductGiftError {
			err := ctx.Status(http.StatusUnprocessableEntity).JSON(NewResponseError("Have a product gift in chart"))

			if err != nil {
				return
			}
			return
		}

		log.Println(err.Error())

		err = ctx.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

		if err != nil {
			return
		}

		return
	}

	err = ctx.Status(http.StatusOK).JSON(order)

	if err != nil {
		return
	}

	return
}

type ResponseError struct {
	Message string `json:"message"`
}

func NewResponseError(message string) ResponseError {
	return ResponseError{
		Message: message,
	}
}