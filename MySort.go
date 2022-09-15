package main

func BubbleSort(nums *[]int) {
	n := len(*nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (*nums)[i] > (*nums)[j] {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			}
		}
	}
	//fmt.Println(*nums)
}

func QuickSort(nums *[]int) {
	if len(*nums) == 0 {
		return
	}
	var partition func(left int, right int) int
	partition = func(left int, right int) int {
		pivot := (*nums)[left]
		for left < right {
			for left < right && (*nums)[right] >= pivot {
				right--
			}
			(*nums)[left] = (*nums)[right]
			for left < right && (*nums)[left] <= pivot {
				left++
			}
			(*nums)[right] = (*nums)[left]
		}
		(*nums)[left] = pivot
		return left
	}
	var _quickSort func(left int, right int)
	_quickSort = func(left int, right int) {
		if left < right {
			index := partition(left, right)
			_quickSort(left, index-1)
			_quickSort(index+1, right)
		}
	}
	_quickSort(0, len(*nums)-1)
}

func SelectionSort(nums *[]int) {
	n := len(*nums)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := minIndex + 1; j < n; j++ {
			if (*nums)[j] < (*nums)[minIndex] {
				minIndex = j
			}
		}
		(*nums)[i], (*nums)[minIndex] = (*nums)[minIndex], (*nums)[i]
	}
}
