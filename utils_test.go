package main_test

import (
	"testing"

	gac "github.com/VitoNaychev/github-actions-course"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("adds 1 + 2 to equal 3", func(t *testing.T) {
		assert.Equal(t, 3, gac.Sum(1, 2))
	})

	t.Run("adds 1 + 2 to equal 4", func(t *testing.T) {
		assert.Equal(t, 4, gac.Sum(1, 2))
	})
}
