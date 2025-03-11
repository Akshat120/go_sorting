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
}

func doTask(arr []int32, algo Algo, wg *sync.WaitGroup) {
	start := time.Now().UnixMilli()
	algo.fn(arr)
	end := time.Now().UnixMilli()

	fmt.Printf("%v completed in %vms\n", algo.fnName, end-start)

	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	n := 100000

	arr := []int32{}

	for range n {
		arr = append(arr, int32(rand.Intn(n)))
	}

	algos := []Algo{
		{
			fn:     compsortingalgo.BubbleSort,
			fnName: "BubbleSort",
		},
		{
			fn:     compsortingalgo.InsertionSort,
			fnName: "InsertionSort",
		},
		{
			fn:     compsortingalgo.SelectionSort,
			fnName: "SelectionSort",
		},
		{
			fn:     compsortingalgo.MergeSort,
			fnName: "MergeSort",
		},
	}

	for _, algo := range algos {
		wg.Add(1)
		new_arr := make([]int32, len(arr))
		copy(new_arr, arr)
		go doTask(new_arr, algo, &wg)
	}

	wg.Wait()

}
