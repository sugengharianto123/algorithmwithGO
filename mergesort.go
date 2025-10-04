package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers4(length int, max int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(max)
	}
	return numbers
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	} else {
		mid := len(numbers) / 2
		left := mergeSort(numbers[:mid])
		right := mergeSort(numbers[mid:])
		return merge(left, right)
	}
}

func main() {
	var length, max int
	fmt.Print("Masukkan panjang array: ")
	fmt.Scan(&length)
	fmt.Print("Masukkan nilai maksimum: ")
	fmt.Scan(&max)
	numbers := generateRandomNumbers4(length, max)
	sorted := mergeSort(numbers)
	fmt.Print("Array yang diurutkan: ", sorted)
}
