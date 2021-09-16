package domain_test

import (
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChart(t *testing.T) {
	t.Run("Create chart with invalid productid should return error", func(t *testing.T) {
		chart := domain.Chart{Products: []domain.ProductChart{{1,0}}}

		err := chart.Validate()

		assert.NotNil(t, err)
	})

	t.Run("Create chart with invalid quantity should return error", func(t *testing.T) {
		chart := domain.Chart{Products: []domain.ProductChart{{0,1}}}

		err := chart.Validate()

		assert.NotNil(t, err)
	})

	t.Run("Create chart should not return error", func(t *testing.T) {
		chart := domain.Chart{Products: []domain.ProductChart{{1,1}}}

		err := chart.Validate()

		assert.Nil(t, err)
	})
}
