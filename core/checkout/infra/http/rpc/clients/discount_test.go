package clients_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/guil95/grpcApi/core/checkout/infra/http/rpc/clients"
	"github.com/guil95/grpcApi/core/checkout/infra/http/rpc/clients/mock"
	"github.com/guil95/grpcApi/core/discount"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscountGrpcClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Test GetDiscount success should return percentage discount", func(t *testing.T) {
		cli := mock.NewMockDiscountClient(ctrl)

		cli.
			EXPECT().
			GetDiscount(
				context.Background(),
				&discount.GetDiscountRequest{ProductID: 1},
				).
			Times(1).Return(&discount.GetDiscountResponse{Percentage: 0.05}, nil)

		grpcClient := clients.NewDiscountGrpcClient(cli)

		response, _ := grpcClient.GetDiscount(1)

		assert.Equal(t, float32(0.05), response)
	})

	t.Run("Test GetDiscount with client error should return 0", func(t *testing.T) {
		cli := mock.NewMockDiscountClient(ctrl)

		cli.
			EXPECT().
			GetDiscount(
				context.Background(),
				&discount.GetDiscountRequest{ProductID: 1},
			).
			Times(1).Return(nil, errors.New("Errouu"))

		grpcClient := clients.NewDiscountGrpcClient(cli)

		response, _ := grpcClient.GetDiscount(1)

		assert.Equal(t, float32(0), response)
	})
}
