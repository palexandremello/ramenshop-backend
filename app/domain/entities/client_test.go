package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {

	t.Run("should return an error if name is empty", func(t *testing.T) {
		c, err := NewClient(1, "", Male, 29)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name is required", err.Error())

	})

	t.Run("should return an error if name have only spaces", func(t *testing.T) {

		c, err := NewClient(1, "         ", Female, 25)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return an error if name have less than 5 characters", func(t *testing.T) {

		c, err := NewClient(1, "taok", Male, 20)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name must have at least 5 characters", err.Error())
	})

}
