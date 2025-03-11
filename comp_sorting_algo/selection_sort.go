package compsortingalgo

func SelectionSort(arr []int32) []int32 {
	n := len(arr)
	for i := range n {
		smallest_idx := i
		for j := i + 1; j < n; j++ {
			if arr[smallest_idx] > arr[j] {
				smallest_idx = j
			}
		}
		swap(&arr[i], &arr[smallest_idx])
	}
	return arr
}
