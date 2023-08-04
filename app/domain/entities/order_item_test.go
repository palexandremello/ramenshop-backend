package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderItem(t *testing.T) {

	t.Run("ensure to create a OrderItem correctly", func(t *testing.T) {

		dish := Dish{
			ID:          1,
			Name:        "Ramen",
			Description: "Delicious ramen dish",
			Photo:       &Photo{URL: "http://example.com/photo.jpg"},
		}

		orderItem := OrderItem{
			ID:     1,
			Dish:   dish,
			Amount: 2,
		}

		assert.Equal(t, 1, orderItem.ID)
		assert.Equal(t, dish, orderItem.Dish)
		assert.Equal(t, 2, orderItem.Amount)
	})
}
