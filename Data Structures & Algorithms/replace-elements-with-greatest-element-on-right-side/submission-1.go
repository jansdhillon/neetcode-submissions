func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)

	for i := range n {
		rightMax := -1
		for j := i + 1; j < n; j++ {
			if arr[j] > rightMax {
				rightMax = arr[j]
			}
		}

		ans[i] = rightMax

	}

	return ans
}
