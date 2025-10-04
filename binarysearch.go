package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers(lenght int, max int) []int {
	numbers := make([]int, lenght)
	for i := 0; i < lenght; i++ {
		numbers[i] = rand.Intn(max)
	}
	return numbers
}

func sortedNumbers(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]

			}
		}

	}
	return numbers
}

func binarySearch(numbers []int, target int) int {
	left, right := 0, len(numbers)-1

	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var length, max, target int
	fmt.Print("Masukkan panjang array: ")
	fmt.Scan(&length)
	fmt.Print("Masukkan nilai maksimum: ")
	fmt.Scan(&max)
	fmt.Print("Masukkan nilai yang dicari: ")
	fmt.Scan(&target)
	numbers := generateRandomNumbers(length, max)
	sorted := sortedNumbers(numbers)
	result := binarySearch(sorted, target)
	if result != -1 {
		fmt.Printf("Angka %d ditemukan pada indeks %d dalam array yang diurutkan: %v\n", target, result, sorted)
	} else {
		fmt.Printf("Angka %d tidak ditemukan dalam array yang diurutkan: %v\n", target, sorted)
	}

}
