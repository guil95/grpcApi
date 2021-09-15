package repositories_test

import (
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/guil95/grpcApi/core/checkout/infra/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepository(t *testing.T) {
	db := []byte("[" +
		"{\"id\": 1,\"title\": \"Ergonomic Wooden Pants\",\"description\": \"Deleniti beatae porro.\",\"amount\": 15157,\"is_gift\": false},  " +
		"{\"id\": 2, \"title\": \"Ergonomic Cotton Keyboard\",\"description\": \"Iste est ratione excepturi repellendus adipisci qui.\",\"amount\": 93811,\"is_gift\": false}," +
		"{\"id\": 3,\"title\": \"Gorgeous Cotton Chips\",\"description\": \"Nulla rerum tempore rem.\",\"amount\": 60356,\"is_gift\": false}," +
		"{\"id\": 4,\"title\": \"Fantastic Frozen Chair\",\"description\": \"Et neque debitis omnis quam enim cupiditate.\",\"amount\": 56230,\"is_gift\": false}," +
		"{\"id\": 5,\"title\": \"Incredible Concrete Soap\",\"description\": \"Dolorum nobis temporibus aut dolorem quod qui corrupti.\",\"amount\": 42647,\"is_gift\": false}," +
		"{\"id\": 6,\"title\": \"Handcrafted Steel Towels\",\"description\": \"Nam ea sed animi neque qui non quis iste.\",\"amount\": 900,\"is_gift\": true}" +
		"]")
	repo := repositories.NewFileRepository(db)

	t.Run("Test Get gift products", func(t *testing.T) {
		products := repo.GetGiftProducts()

		assert.True(t, hasGiftProduct(products))
	})

	t.Run("Test Get all products", func(t *testing.T) {
		products := repo.GetProducts()

		assert.True(t, len(products) == 6)
	})

	t.Run("Test get products by chart pass an nonexistent product than return only existent products", func(t *testing.T) {
		chart := domain.Chart{Products: []domain.ProductChart{{1,1}, {1,2}}}

		products := repo.GetProductsByChart(&chart)

		assert.True(t, len(products) == 2)
	})
}

func hasGiftProduct(products []domain.Product) bool {
	for _,product := range products {
		if product.Gift == true {
			return true
		}
	}

	return false
}