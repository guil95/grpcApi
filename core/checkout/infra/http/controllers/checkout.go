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
		checkout(ctx, use_cases.NewCreateCheckoutUseCase(
			clients.NewDiscountGrpcClient(discount.NewDiscountClient(conn)),
			repositories.NewFileRepository(file),
			),
		)
		return nil
	})
}

func checkout(ctx *fiber.Ctx, service domain.UseCase) error {
	var chartPayload = new(domain.Chart)

	if err := ctx.BodyParser(chartPayload); err != nil {
		log.Println(err)

		err := ctx.Status(http.StatusUnprocessableEntity).JSON(newResponseError("Unprocessable entity"))

		if err != nil {
			log.Println(err)
			return nil
		}

		return nil
	}

	order, err := service.Checkout(chartPayload)

	if err != nil {
		log.Println(err.Error())

		switch err {
			case domain.ProductGiftError:
				err := ctx.Status(http.StatusUnprocessableEntity).JSON(newResponseError("Have a product gift in chart"))

				if err != nil {
					log.Println(err)
					return nil
				}
				return nil
			case domain.ProductNotFoundError :
				err := ctx.Status(http.StatusNotFound).JSON(newResponseError("Products not found"))

				if err != nil {
					log.Println(err)
					return nil
				}
				return nil
			default:
				err = ctx.Status(http.StatusInternalServerError).JSON(newResponseError("Internal Server Error"))

				if err != nil {
					log.Println(err)
					return nil
				}

				return nil
		}
	}

	err = ctx.Status(http.StatusOK).JSON(order)

	if err != nil {
		log.Println(err)
		return nil
	}

	return nil
}

type ResponseError struct {
	Message string `json:"message"`
}

func newResponseError(message string) ResponseError {
	return ResponseError{
		Message: message,
	}
}