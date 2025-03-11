package compsortingalgo

func BubbleSort(arr []int32) []int32 {
	n := len(arr)

	for i := range n {
		for j := range n - 1 - i {

			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}

		}
	}

	return arr
}
