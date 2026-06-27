package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rows := rand.Intn(20) + 1
	cols := rand.Intn(20) + 1

	matrix := make([][]int, rows)
	used := make(map[int]bool)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for {
				num := rand.Intn(1000) + 1
				if !used[num] {
					used[num] = true
					matrix[i][j] = num
					break
				}
			}
		}
	}

	fmt.Printf("Двумерный массив %dx%d:\n", rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}
}
