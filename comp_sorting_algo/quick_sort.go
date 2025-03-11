package compsortingalgo

func partition(arr []int32, i, j int32) int32 {

	pivot := i

	smaller := []int32{}
	greater := []int32{}
	pivotValue := arr[pivot]
	for k := i + 1; k <= j; k++ {
		if arr[k] <= arr[pivot] {
			smaller = append(smaller, arr[k])
		} else {
			greater = append(greater, arr[k])
		}
	}

	itr := i
	for _, val := range smaller {
		arr[itr] = val
		itr++
	}
	newPivot := itr
	arr[itr] = pivotValue
	itr++
	for _, val := range greater {
		arr[itr] = val
		itr++
	}

	return newPivot
}

func quickSortHelper(arr []int32, i, j int32) {
	if i >= j {
		return
	}

	pivot := partition(arr, i, j)
	quickSortHelper(arr, i, pivot-1)
	quickSortHelper(arr, pivot+1, j)

}

func QuickSort(arr []int32) []int32 {

	quickSortHelper(arr, 0, int32(len(arr)-1))

	return arr

}
