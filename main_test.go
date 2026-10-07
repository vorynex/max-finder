package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	// Пустой слайс
	require.Empty(t, generateRandomElements(0))

	// Отрицательный размер
	require.Empty(t, generateRandomElements(-5))

	// Нормальный размер
	require.Len(t, generateRandomElements(100), 100)
}

func TestMaximum(t *testing.T) {
	// Пустой слайс
	require.Equal(t, 0, maximum([]int{}))

	// Один элемент
	require.Equal(t, 5, maximum([]int{5}))

	// Несколько элементов
	require.Equal(t, 10, maximum([]int{1, 5, 10, 3, 7}))

	// Все одинаковые
	require.Equal(t, 3, maximum([]int{3, 3, 3}))

	// Максимум в конце
	require.Equal(t, 100, maximum([]int{1, 2, 3, 100}))

	// Максимум в начале
	require.Equal(t, 100, maximum([]int{100, 2, 3, 1}))
}

func TestMaxChunks(t *testing.T) {
	// Пустой слайс
	require.Equal(t, 0, maxChunks([]int{}))

	// Один элемент
	require.Equal(t, 5, maxChunks([]int{5}))

	// Маленький слайс (меньше CHUNKS)
	require.Equal(t, 10, maxChunks([]int{1, 5, 10, 3, 7}))

	// Большой слайс
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	require.Equal(t, 999, maxChunks(data))

	// Максимум в середине
	data = make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	data[500] = 10000
	require.Equal(t, 10000, maxChunks(data))
}