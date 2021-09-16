package api

import (
	"github.com/gorilla/mux"
	"github.com/guil95/grpcApi/config"
	checkoutControllers "github.com/guil95/grpcApi/core/checkout/infra/http/controllers"
	"github.com/guil95/grpcApi/core/checkout/infra/http/rpc/clients"
	"github.com/guil95/grpcApi/core/checkout/infra/repositories"
	"github.com/guil95/grpcApi/core/checkout/use_cases"
	"github.com/guil95/grpcApi/core/discount"
	pkg "github.com/guil95/grpcApi/pkg/grpc"
	"io/ioutil"
	"log"
	"net/http"
)

func Run() {
	conf, err := config.RetrieveConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listen server on "+ conf.ApiPort)

	r := mux.NewRouter()

	conn := pkg.Conn()
	client := clients.NewDiscountGrpcClient(discount.NewDiscountClient(conn))

	file, _ := ioutil.ReadFile(conf.DbFile)

	repo := repositories.NewFileRepository(file)

	checkoutControllers.MakeCheckoutHandler(r, use_cases.NewCreateCheckoutUseCase(client, repo, conf))

	log.Fatal(http.ListenAndServe(":"+conf.ApiPort, r))
}
