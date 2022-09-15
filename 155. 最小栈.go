package main

import "math"

type MinStack struct {
	min           int
	minValueCount int
	stack         []int
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.min {
		this.min = val
		this.minValueCount = 1
	} else if val == this.min {
		this.minValueCount++
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top == this.min {
		this.minValueCount--
		if this.minValueCount == 0 {
			minV := math.MaxInt64
			count := 0
			for _, x := range this.stack {
				if x < minV {
					minV = x
					count = 1
				} else if x == minV {
					count++
				}
			}
			this.min = minV
			this.minValueCount = count
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
