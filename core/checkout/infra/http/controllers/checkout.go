package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/guil95/grpcApi/core/checkout/infra/http/rpc/clients"
	"github.com/guil95/grpcApi/core/checkout/infra/repositories"
	"github.com/guil95/grpcApi/core/checkout/use_cases"
	"github.com/guil95/grpcApi/core/discount"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func CreateApi(app *fiber.App, file []byte, conn grpc.ClientConnInterface) {
	app.Post("/checkout", func(ctx *fiber.Ctx) error {
		checkout(ctx, use_cases.NewService(
			clients.NewDiscountGrpcClient(discount.NewDiscountClient(conn)),
			repositories.NewFileRepository(file),
			),
		)
		return nil
	})
}

func checkout(ctx *fiber.Ctx, service *use_cases.CreateCheckoutUseCase) {
	var chartPayload = new(domain.Chart)

	if err := ctx.BodyParser(chartPayload); err != nil {
		log.Println(err)

		err := ctx.Status(http.StatusUnprocessableEntity).JSON(NewResponseError("Unprocessable entity"))

		if err != nil {
			log.Println(err)
			return
		}

		return
	}

	order, err := service.Checkout(chartPayload)

	if err != nil {
		log.Println(err.Error())

		switch err {
			case domain.ProductGiftError:
				err := ctx.Status(http.StatusUnprocessableEntity).JSON(NewResponseError("Have a product gift in chart"))

				if err != nil {
					log.Println(err)
					return
				}
				return
			case domain.ProductNotFoundError :
				err := ctx.Status(http.StatusNotFound).JSON(NewResponseError("Products not found"))

				if err != nil {
					log.Println(err)
					return
				}
				return
			default:
				err = ctx.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

				if err != nil {
					log.Println(err)
					return
				}

				return
		}
	}

	err = ctx.Status(http.StatusOK).JSON(order)

	if err != nil {
		log.Println(err)
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