package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers2(length int, max int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(max)
	}
	return numbers
}

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	var length, max int
	fmt.Print("Masukkan panjang array: ")
	fmt.Scan(&length)
	fmt.Print("Masukkan nilai maksimum: ")
	fmt.Scan(&max)
	numbers := generateRandomNumbers2(length, max)
	sorted := bubbleSort(numbers)
	fmt.Print("Array yang diurutkan: ", sorted)
}
