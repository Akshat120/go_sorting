package compsortingalgo

func InsertionSort(arr []int32) []int32 {
	n := len(arr)
	for i := range n {
		j := i
		for j > 0 && (arr[j] < arr[j-1]) {
			swap(&arr[j], &arr[j-1])
			j--
		}
	}
	return arr
}
