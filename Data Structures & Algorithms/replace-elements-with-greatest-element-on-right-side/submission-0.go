import (
	"slices"
)

func replaceElements(arr []int) []int {
	r := len(arr) - 1

	for l := 0; l < r; l++ {
		arr[l] = slices.Max(arr[l+1:])
	}

	arr[r] = -1

	return arr
}
