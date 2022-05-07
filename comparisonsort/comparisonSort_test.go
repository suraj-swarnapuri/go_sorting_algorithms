package comparisonsort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSortWithInt(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, expected, BubbleSort(expected))
}

func TestBubbleSortWithString(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	require.Equal(t, expected, BubbleSort(expected))
}

func TestBubbleSortWithFloats(t *testing.T) {
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	require.Equal(t, expected, BubbleSort(expected))
}

func TestInsertionSortWithInt(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, expected, InsertionSort(expected))
}

func TestInsertionSortWithString(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	require.Equal(t, expected, InsertionSort(expected))
}

func TestInsertionSortWithFloats(t *testing.T) {
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	require.Equal(t, expected, InsertionSort(expected))
}

func TestMergeSortWithInt(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, expected, MergeSort(expected))
}

func TestMergeSortWithString(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	require.Equal(t, expected, MergeSort(expected))
}

func TestMergeSortWithFloats(t *testing.T) {
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	require.Equal(t, expected, MergeSort(expected))
}
