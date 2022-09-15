package main

import "math"

func maxProfit(prices []int) (ans int) {
	minPrice := math.MaxInt32
	for _, x := range prices {
		if x < minPrice {
			minPrice = x
		} else {
			if x > ans+minPrice {
				ans = x - minPrice
			}
		}
	}
	return
}
