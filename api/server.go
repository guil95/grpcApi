package api

import (
	"github.com/gorilla/mux"
	checkoutControllers "github.com/guil95/grpcApi/core/checkout/infra/http/controllers"
	"github.com/guil95/grpcApi/core/checkout/infra/http/rpc/clients"
	"github.com/guil95/grpcApi/core/checkout/infra/repositories"
	"github.com/guil95/grpcApi/core/checkout/use_cases"
	"github.com/guil95/grpcApi/core/discount"
	pkg "github.com/guil95/grpcApi/pkg/grpc"
	"log"
	"net/http"
	"os"
)

func Run(file []byte) {
	port := os.Getenv("API_PORT")

	log.Println("Listen server on "+ port)

	r := mux.NewRouter()

	conn := pkg.Conn()
	client := clients.NewDiscountGrpcClient(discount.NewDiscountClient(conn))
	repo := repositories.NewFileRepository(file)

	checkoutControllers.MakeCheckoutHandler(r, use_cases.NewCreateCheckoutUseCase(client, repo))

	log.Fatal(http.ListenAndServe(":8000", r))
}
