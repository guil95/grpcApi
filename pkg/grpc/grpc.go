package pkg

import (
	"github.com/guil95/grpcApi/config"
	"google.golang.org/grpc"
	"log"
)

func Conn() *grpc.ClientConn {
	conf, err := config.RetrieveConfig()

	if err != nil {
		log.Fatal(err)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(conf.DiscountClientHost+":"+conf.DiscountClientPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
