package main

import "fmt"

// 计算正整数分拆数（integer partition number）。
// 分拆数 p(n) 表示将正整数 n 表示为若干个正整数之和（不计顺序）的方法数。
// 例如 p(4) = 5: 4, 3+1, 2+2, 2+1+1, 1+1+1+1。
func Partition(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}

	memo := make(map[[2]int]int)
	var dfs func(sum, low int) int
	dfs = func(sum, low int) int {
		if sum == 0 {
			return 1
		}
		if sum < low {
			return 0
		}

		if val, ok := memo[[2]int{sum, low}]; ok {
			return val
		}

		res := 0
		for num := low; num <= sum; num++ {
			res += dfs(sum-num, num)
		}
		memo[[2]int{sum, low}] = res
		return res
	}

	return dfs(n, 1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d 的分拆数是 %d\n", n, Partition((n)))
}
