package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MinHeap []int // 定义一个类型
type MaxHeap []int

func (h MinHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MinHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MinHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (h MaxHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MaxHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MaxHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	for _, x := range nums1 {
		if maxHeap.Len() == 0 {
			heap.Push(maxHeap, x)
			continue
		}
		if x > (*maxHeap)[0] {
			heap.Push(minHeap, x)
		} else {
			heap.Push(maxHeap, x)
		}
		for math.Abs(float64(maxHeap.Len()-minHeap.Len())) > 1 {
			if maxHeap.Len() > minHeap.Len() {
				heap.Push(minHeap, heap.Pop(maxHeap))
			} else {
				heap.Push(maxHeap, heap.Pop(minHeap))
			}
		}
	}
	for _, x := range nums2 {
		if maxHeap.Len() == 0 {
			heap.Push(maxHeap, x)
			continue
		}
		if x > (*maxHeap)[0] {
			heap.Push(minHeap, x)
		} else {
			heap.Push(maxHeap, x)
		}
		for math.Abs(float64(maxHeap.Len()-minHeap.Len())) > 1 {
			if maxHeap.Len() > minHeap.Len() {
				heap.Push(minHeap, heap.Pop(maxHeap))
			} else {
				heap.Push(maxHeap, heap.Pop(minHeap))
			}
		}
	}
	fmt.Println(*minHeap)
	fmt.Println(*maxHeap)

	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64((*minHeap)[0]) + float64((*maxHeap)[0])) / 2
	} else {
		if minHeap.Len() > maxHeap.Len() {
			return float64((*minHeap)[0])
		} else {
			return float64((*maxHeap)[0])
		}
	}

}
