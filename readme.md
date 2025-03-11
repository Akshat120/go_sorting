# Go Sorting Algorithm Benchmark

This is a Go program that benchmarks different sorting algorithms by measuring their execution time on a randomly generated array of integers. The program uses goroutines and WaitGroups to run multiple sorting algorithms concurrently.

## Features

- Benchmarks four sorting algorithms:
  - Bubble Sort
  - Insertion Sort
  - Selection Sort
  - Merge Sort
- Runs algorithms concurrently using goroutines
- Measures and displays execution time for each algorithm
- Uses a randomly generated array of 100,000 integers