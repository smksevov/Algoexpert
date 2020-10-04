package main

func SelectionSort(array []int) []int {

	currentIndex := 0
	for currentIndex < len(array)-1 {
		smallestIndex := currentIndex

		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		array[smallestIndex], array[currentIndex] = array[currentIndex], array[smallestIndex]

		currentIndex++
	}
	return array
}

// My solution
func SelectionSort1(array []int) []int {

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}