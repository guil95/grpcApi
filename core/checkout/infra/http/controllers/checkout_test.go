package http

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/guil95/grpcApi/core/checkout/infra/http/controllers/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCheckoutApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockCreateCheckoutUseCase(ctrl)
	r := mux.NewRouter()
	MakeCheckoutHandler(r, service)

	t.Run("Test create checkout with success should return 200", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		service.EXPECT().Checkout(&domain.Chart{Products: []domain.ProductChart{{1,1}}})

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": 1,
					"quantity": 1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Test create checkout with invalid payload should return 422", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": 0,
					"quantity": 0 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	})

	t.Run("Test create checkout with invalid chart should return 422", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": -1,
					"quantity": -1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	})

	t.Run("Test create checkout with invalid json should return 500", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
					"id": 1,
					"quantity": 1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("Test create checkout with gift product should return 422", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		service.EXPECT().Checkout(&domain.Chart{Products: []domain.ProductChart{{1,1}}}).Return(nil, domain.ProductGiftError)

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": 1,
					"quantity": 1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	})

	t.Run("Test create checkout with non exists product should return 404", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		service.EXPECT().Checkout(&domain.Chart{Products: []domain.ProductChart{{1,1}}}).Return(nil, domain.ProductNotFoundError)

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": 1,
					"quantity": 1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Test create checkout with unexpected error should return 500", func(t *testing.T) {
		path, err := r.GetRoute("create_checkout").GetPathTemplate()
		assert.Nil(t, err)
		assert.Equal(t, "/checkout", path)

		service.EXPECT().Checkout(&domain.Chart{Products: []domain.ProductChart{{1,1}}}).Return(nil, errors.New("unexpected error"))

		ts := httptest.NewServer(checkout(service))
		defer ts.Close()

		payload := fmt.Sprintf(`{
			"products": [
				{
					"id": 1,
					"quantity": 1 
				}
			]
		}`)

		res, err := http.Post(ts.URL, "application/json", strings.NewReader(payload))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
