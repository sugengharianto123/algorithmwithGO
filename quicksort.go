package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers3(length int, max int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(max)
	}
	return numbers
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	} else {
		pivot := rand.Intn(len(numbers))
		less := []int{}
		greater := []int{}
		for i, v := range numbers {
			if i != pivot {
				if v <= numbers[pivot] {
					less = append(less, v)
				} else {
					greater = append(greater, v)
				}
			} else {
				continue
			}
		}
		sortedLess := quickSort(less)
		sortedGreater := quickSort(greater)
		return append(append(sortedLess, numbers[pivot]), sortedGreater...)
	}

}

func main() {
	var length, max int
	fmt.Print("Masukkan panjang array: ")
	fmt.Scan(&length)
	fmt.Print("Masukkan nilai maksimum: ")
	fmt.Scan(&max)
	numbers := generateRandomNumbers3(length, max)
	sorted := quickSort(numbers)
	fmt.Print("Array yang diurutkan: ", sorted)
}
