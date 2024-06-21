package utils

import "github.com/gammazero/deque"

func InsertAsc(array *deque.Deque[float64], number float64) {
	index := BinarySearchAsc(*array, number)
	array.Insert(index, number)
}

func RemoveAsc(array *deque.Deque[float64], number float64) {
	index := BinarySearchAsc(*array, number)
	if index != -1 {
		array.Remove(index)
	}
}

func InsertDesc(array *deque.Deque[float64], number float64) {
	index := BinarySearchDesc(*array, number)
	array.Insert(index, number)
}

func RemoveDesc(array *deque.Deque[float64], number float64) {
	index := BinarySearchDesc(*array, number)
	if index != -1 {
		array.Remove(index)
	}
}

func BinarySearchAsc(array deque.Deque[float64], number float64) int {
	low := 0
	high := array.Len()

	for low < high {
		mid := (low + high) / 2
		if array.At(mid) < number {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func BinarySearchDesc(array deque.Deque[float64], number float64) int {
	low := 0
	high := array.Len()

	for low < high {
		mid := (low + high) / 2
		if array.At(mid) > number {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
