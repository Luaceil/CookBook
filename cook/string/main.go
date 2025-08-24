package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var length int
		fmt.Scan(&length)
		var tStr, s string
		fmt.Scan(&s, &tStr)
		println(tStr)
		diffCount := 0
		for i := 0; i < N; i++ {
			if s[i] != tStr[i] {
				diffCount++
			}
		}
		case1 := diffCount == 1

		// 检查字符频率是否相同，判断单字符错位
		var freqS, freqT [26]int
		for _, c := range s {
			freqS[c-'a']++
		}
		for _, c := range tStr {
			freqT[c-'a']++
		}
		case2 := false
		if freqS == freqT {
			// 双指针统计按顺序匹配的字符数
			i, j := 0, 0
			match := 0
			for i < N && j < N {
				if s[i] == tStr[j] {
					match++
					i++
					j++
				} else {
					j++
				}
			}
			if match == N-1 {
				case2 = true
			}
		}

		// 综合判断两种错误是否最多出现一种
		if (case1 && !case2) || (case2 && !case1) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
