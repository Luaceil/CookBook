package main

import (
	"fmt"
	"math"
)

func main() {
	var N, T int
	fmt.Scan(&N, &T)
	dp := make([]int, T+1)
	for i := 0; i < N; i++ {
		var count, score, t int
		fmt.Scan(&count, &score, &t)
		k := 1
		for count > 0 {
			cnt := int(math.Min(float64(k), float64(count)))
			value := cnt * score
			time := cnt * t
			for j := T; j >= time; j-- {
				if dp[j-time]+value > dp[j] {
					dp[j] = dp[j-time] + value
				}
			}
			count -= cnt
			k *= 2
		}
	}
	fmt.Println(dp[T])
}

// 回溯问题
