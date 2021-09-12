package clients

import (
	"context"
	"github.com/guil95/grpcApi/core/discount"
	pkg "github.com/guil95/grpcApi/pkg/grpc"
	"google.golang.org/grpc"
	"log"
)

type DiscountGrpcClient struct {
}

func NewDiscountGrpcClient() *DiscountGrpcClient {
	return &DiscountGrpcClient{}
}

func (d *DiscountGrpcClient) GetDiscount(productId int32) (float32, error) {
	conn := pkg.Conn()

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	cli := discount.NewDiscountClient(conn)

	response, err := cli.GetDiscount(context.Background(), &discount.GetDiscountRequest{ProductID: productId})

	if err != nil {
		log.Println(err)
		return 0, nil
	}

	return response.GetPercentage(), nil
}
