package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {

	t.Run("should return an error if name is empty", func(t *testing.T) {
		c, err := NewClient("", Male, 29)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name is required", err.Error())

	})

	t.Run("should return an error if name have only spaces", func(t *testing.T) {

		c, err := NewClient("         ", Female, 25)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return an error if name have less than 5 characters", func(t *testing.T) {

		c, err := NewClient("taok", Male, 20)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "name must have at least 5 characters", err.Error())
	})

	t.Run("should return an error if age is invalid", func(t *testing.T) {

		c, err := NewClient("Alexandre", Female, 140)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "age must be between 0 and 120", err.Error())
	})

	t.Run("should return an error if gender is not contemplated", func(t *testing.T) {

		c, err := NewClient("Alexandre", 3, 29)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, "For now we have only two genders, sorry. We Will fix it soon", err.Error())
	})

	t.Run("ensure should return a valid client", func(t *testing.T) {

		c, err := NewClient("Alexandre", Male, 29)

		assert.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, "Alexandre", c.Name)
		assert.Equal(t, Male, c.Gender)
		assert.Equal(t, 29, c.Age)
	})

}
