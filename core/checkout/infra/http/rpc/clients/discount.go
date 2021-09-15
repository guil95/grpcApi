package clients

import (
	"context"
	"github.com/guil95/grpcApi/core/discount"
	"log"
)

type DiscountGrpcClient struct {
	cli discount.DiscountClient
}

func NewDiscountGrpcClient(cli discount.DiscountClient) *DiscountGrpcClient {
	return &DiscountGrpcClient{cli: cli}
}

func (d *DiscountGrpcClient) GetDiscount(productId int32) (float32, error) {
	response, err := d.cli.GetDiscount(
		context.Background(),
		&discount.GetDiscountRequest{ProductID: productId})

	if err != nil {
		log.Println(err)
		return 0, nil
	}

	return response.GetPercentage(), nil
}
