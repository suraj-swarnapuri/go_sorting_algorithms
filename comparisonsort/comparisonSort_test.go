package comparisonsort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	intInput    = []int{10, 6, 4, 7, 3, 9, 1, 8, 2, 5}
	floatInput  = []float64{10.0, 6.6, 4.4, 7.7, 3.3, 9.9, 1.1, 8.8, 2.2, 5.5}
	stringInput = []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}

	expectedInt    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedString = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	expectedFloat  = []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
)

func TestBubbleSortWithInt(t *testing.T) {
	require.Equal(t, expectedInt, BubbleSort(intInput))
}

func TestBubbleSortWithString(t *testing.T) {
	require.Equal(t, expectedString, BubbleSort(stringInput))
}

func TestBubbleSortWithFloats(t *testing.T) {
	require.Equal(t, expectedFloat, BubbleSort(floatInput))
}

func TestInsertionSortWithInt(t *testing.T) {
	require.Equal(t, expectedInt, InsertionSort(intInput))
}

func TestInsertionSortWithString(t *testing.T) {
	require.Equal(t, expectedString, InsertionSort(stringInput))
}

func TestInsertionSortWithFloats(t *testing.T) {
	require.Equal(t, expectedFloat, InsertionSort(floatInput))
}

func TestMergeSortWithInt(t *testing.T) {
	require.Equal(t, expectedInt, MergeSort(intInput))
}

func TestMergeSortWithString(t *testing.T) {
	require.Equal(t, expectedString, MergeSort(stringInput))
}

func TestMergeSortWithFloats(t *testing.T) {
	require.Equal(t, expectedFloat, MergeSort(floatInput))
}
