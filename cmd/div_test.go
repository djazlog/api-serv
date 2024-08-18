package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiv(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		a := 10.0
		b := 5.0

		expected := 2.0
		require.Equal(t, expected, Div(a, b))
	})
	t.Run("test 2", func(t *testing.T) {
		a := 10.0
		b := 0.0

		expected := 0.0
		require.Equal(t, expected, Div(a, b))
	})
}
