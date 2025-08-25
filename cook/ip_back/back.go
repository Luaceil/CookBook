package main

import (
	"fmt"
)

func main() {
	var ipOri string
	fmt.Scan(&ipOri)

	results := restoreIpAddresses(ipOri)
	for _, ip := range results {
		fmt.Println(ip)
	}
}

func restoreIpAddresses(s string) []string {
	var result []string
	var path []string
	backtrack(s, 0, path, &result)
	return result
}

func backtrack(s string, start int, path []string, result *[]string) {
	// 如果已经分割为4段且用完所有字符，则找到一个有效IP
	if len(path) == 4 {
		if start == len(s) {
			ip := fmt.Sprintf("%s.%s.%s.%s", path[0], path[1], path[2], path[3])
			*result = append(*result, ip)
		}
		return
	}

	// 尝试不同长度的段（1-3位）
	for i := start; i < start+3 && i < len(s); i++ {
		segment := s[start : i+1]

		// 检查段是否有效
		if isValid(segment) {
			path = append(path, segment)
			backtrack(s, i+1, path, result)
			path = path[:len(path)-1] // 回溯
		}
	}
}

func isValid(segment string) bool {
	// 检查前导零（除了"0"本身）
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}

	// 检查数值范围
	num := 0
	for i := 0; i < len(segment); i++ {
		num = num*10 + int(segment[i]-'0')
	}

	return num >= 0 && num <= 255
}
