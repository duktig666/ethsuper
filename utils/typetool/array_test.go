package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEqualUint64Array(t *testing.T) {
	arr1 := []uint64{1, 2, 3, 4}
	arr2 := []uint64{1, 2, 4, 3}
	arr3 := []uint64{1, 2, 4, 4}

	require.Equal(t, true, EqualUint64Array(arr1, arr1))
	require.Equal(t, false, EqualUint64Array(arr1, arr2))
	require.Equal(t, false, EqualUint64Array(arr1, arr3))
	require.Equal(t, false, EqualUint64Array(arr2, arr3))

	require.Equal(t, true, EqualUint64ArraySorted(arr1, arr1))
	require.Equal(t, true, EqualUint64ArraySorted(arr1, arr2))
	require.Equal(t, false, EqualUint64ArraySorted(arr1, arr3))
	require.Equal(t, false, EqualUint64ArraySorted(arr2, arr3))
}
