package main

import "strings"

func generateParenthesis(n int) (ans []string) {
	var path []string
	var f func()
	open, close := 0, 0
	f = func() {
		if open < close || open > n || close > n {
			return
		}
		if open+close == 2*n {
			ans = append(ans, strings.Join(path, ""))
		}
		path = append(path, "(")
		open++
		f()
		open--
		path = path[:len(path)-1]
		close++
		path = append(path, ")")
		f()
		close--
		path = path[:len(path)-1]

	}
	f()
	return
}
