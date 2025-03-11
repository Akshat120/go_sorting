package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	compsortingalgo "github.com/Akshat120/go_comp_sorting_algo/comp_sorting_algo"
)

type Algo struct {
	fn     func([]int32) []int32
	fnName string
	debug  bool
	skip   bool
}

func doTask(arr []int32, algo Algo, results [][]int32, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	if algo.skip {
		return
	}
	start := time.Now().UnixMilli()
	sortedArr := algo.fn(arr)
	end := time.Now().UnixMilli()

	if algo.debug {
		fmt.Println("Sorted Arr:", arr)
	}

	fmt.Printf("%v completed in %vms\n", algo.fnName, end-start)

	results[idx] = sortedArr
}

func main() {
	var wg sync.WaitGroup
	n := 1000000
	arr := []int32{}

	for range n {
		arr = append(arr, int32(rand.Intn(n)))
	}

	// fmt.Println("Original Arr:", arr)

	algos := []Algo{
		{
			fn:     compsortingalgo.BubbleSort,
			fnName: "BubbleSort",
			// skip:   true,
			// debug: true,
		},
		{
			fn:     compsortingalgo.InsertionSort,
			fnName: "InsertionSort",
			// skip:   true,
			// debug: true,
		},
		{
			fn:     compsortingalgo.SelectionSort,
			fnName: "SelectionSort",
			// skip:   true,
			// debug: true,
		},
		{
			fn:     compsortingalgo.MergeSort,
			fnName: "MergeSort",
			// skip:   true,
			// debug: true,
		},
		{
			fn:     compsortingalgo.QuickSort,
			fnName: "QuickSort",
			// debug:  true,
		},
	}

	results := make([][]int32, len(algos))
	for idx, algo := range algos {
		wg.Add(1)
		new_arr := make([]int32, len(arr))
		copy(new_arr, arr)
		go doTask(new_arr, algo, results, idx, &wg)
	}

	wg.Wait()

	fmt.Printf("CrossChecking.")
	for i := range len(results) - 1 {
		for j := range len(results[i]) {
			if results[i][j] == results[i+1][j] {
				continue
			} else {
				fmt.Println("ERROR values in different algos are not matching!!")
				return
			}
		}
		fmt.Printf(".")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Completed")

}
