package stablesort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountingSort(t *testing.T) {
	var max = 10
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, arr, CountingSort(arr, max))
}
