package compsortingalgo

func merge(arr []int32, i, m, j int) {

	tmp := make([]int32, j-i+1)
	k := 0
	p := i
	q := m + 1

	for p <= m && q <= j {
		if arr[p] < arr[q] {
			tmp[k] = arr[p]
			p++
		} else {
			tmp[k] = arr[q]
			q++
		}
		k++
	}

	for p <= m {
		tmp[k] = arr[p]
		p++
		k++
	}

	for q <= j {
		tmp[k] = arr[q]
		q++
		k++
	}
	k = 0
	for p := i; p <= j; p++ {
		arr[p] = tmp[k]
		k++
	}
}

func mergeSortHelper(arr []int32, i, j int) {
	if i == j {
		return
	}

	m := (i + j) / 2

	mergeSortHelper(arr, i, m)
	mergeSortHelper(arr, m+1, j)

	merge(arr, i, m, j)

}

func MergeSort(arr []int32) []int32 {
	mergeSortHelper(arr, 0, len(arr)-1)
	return arr
}
