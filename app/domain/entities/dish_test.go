package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDish(t *testing.T) {

	t.Run("it sets the ID, name, description and photo", func(t *testing.T) {
		photo := &Photo{URL: "http://example.com/photo.jpg"}
		dish := &Dish{ID: 1, Name: "Ramen", Description: "Delicious ramen dish", Photo: photo}

		assert.Equal(t, 1, dish.ID)
		assert.Equal(t, "Ramen", dish.Name)
		assert.Equal(t, "Delicious ramen dish", dish.Description)
		assert.Equal(t, photo, dish.Photo)
	})
}
