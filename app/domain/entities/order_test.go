package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {

	t.Run("ensure to create an order correctly", func(t *testing.T) {

		dish := Dish{
			ID:          1,
			Name:        "Ramen",
			Description: "Delicious ramen dish",
			Photo:       nil,
		}

		orderItem := OrderItem{
			ID:     1,
			Dish:   dish,
			Amount: 2,
		}

		createdAt := time.Now()

		order := Order{
			ID:        1,
			Items:     []OrderItem{orderItem},
			CreatedAt: createdAt,
		}

		assert.Equal(t, 1, order.ID)
		assert.Equal(t, []OrderItem{orderItem}, order.Items)
		assert.Equal(t, createdAt, order.CreatedAt)
	})

}
