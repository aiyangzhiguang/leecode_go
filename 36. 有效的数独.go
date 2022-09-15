package main

import (
	"fmt"
	"sync"
)

var (
	ch = make(chan int, 27)
	wg sync.WaitGroup
)

func deepCopySudoku(board [][]byte) [][]byte {
	var tmp [][]byte
	for i := 0; i < 9; i++ {
		var t []byte
		for j := 0; j < 9; j++ {
			t = append(t, board[i][j])
		}
		tmp = append(tmp, t)
	}
	return tmp
}

func isValidSudoku(board [][]byte) bool {
	wg.Add(27)
	for i := 0; i < 9; i++ {
		i := i
		go func() {
			wg.Done()
			m := make(map[byte]int)
			for u := 0; u < 9; u++ {
				x := board[i][u]
				if x != '.' {
					if m[x] > 0 {
						ch <- -1
						return
					}
				}
				m[x]++
			}
			ch <- 1
		}()
	}
	for i := 0; i < 9; i++ {
		i := i
		go func() {
			wg.Done()
			m := make(map[byte]int)
			for u := 0; u < 9; u++ {
				x := board[u][i]
				if x != '.' {
					if m[x] > 0 {
						ch <- -1
						return
					}
				}
				m[x]++
			}
			ch <- 1
		}()
	}
	for i := 0; i < 9; i += 3 {
		i := i
		for j := 0; j < 9; j += 3 {
			j := j
			go func() {
				wg.Done()
				hs := make(map[byte]int)
				for m := i; m < i+3; m++ {
					for n := j; n < j+3; n++ {
						x := board[m][n]
						if x != '.' {
							if hs[x] > 0 {
								ch <- -1
								return
							}
						}
						hs[x]++
					}
				}
				ch <- 1
			}()
		}
	}
	wg.Wait()
	i := 27
	for i > 0 {
		x, _ := <-ch
		if x == -1 {
			return false
		} else if x == 1 {
			i--
		} else {
			fmt.Println('1')
		}
	}
	return true
}
