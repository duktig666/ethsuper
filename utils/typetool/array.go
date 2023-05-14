package util

import (
	"encoding/json"
	"sort"
)

func EqualUint64Array(arr1, arr2 []uint64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	arr1Json, _ := json.Marshal(arr1)
	arr2Json, _ := json.Marshal(arr2)
	return string(arr1Json) == string(arr2Json)
}

func EqualUint64ArraySorted(arr1, arr2 []uint64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	arr1Json, _ := json.Marshal(arr1)
	arr2Json, _ := json.Marshal(arr2)
	return string(arr1Json) == string(arr2Json)
}
