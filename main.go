package main

import "fmt"

func check(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

//func main() {
//	var wg sync.WaitGroup
//	var wg2 sync.WaitGroup
//	for j := 0; j < 10; j++ {
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			wg2.Add(1)
//			//i := i
//			go func() {
//				rand.Seed(time.Now().UnixMilli())
//				n := rand.Intn(100)
//				var nums []int
//				for m := 0; m < n; m++ {
//					nums = append(nums, rand.Intn(100000))
//				}
//				SelectionSort(&nums)
//				fmt.Println(check(nums))
//				//fmt.Println()
//				defer wg.Done()
//				defer wg2.Done()
//			}()
//		}
//		fmt.Println(j)
//		wg.Wait()
//		time.Sleep(time.Second)
//	}
//	wg2.Wait()
//}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(board))
}
