package main

import (
	"fmt"
)

//var result = make([]string, 0)

func main() {
	//处理输入 ——> 数组
	var ipOri string
	fmt.Scan(&ipOri)
	nums := make([]int, len(ipOri))
	for i := 0; i < len(ipOri); i++ {
		a := ipOri[i]
		nums[i] = int((a - '0'))
		//println(nums[i])
	}
	//贪心
	for i := 0; i < 3 && i < len(nums)-1; i++ {
		for j := i + 1; j < i+4 && j < len(nums)-1; j++ {
			for k := j + 1; k < j+4 && k < len(nums)-1; k++ {
				var ires, jres, kres, mres int
				for m := 0; m < i+1; m++ {
					ires = nums[m]*sqr(i-m) + ires
				}
				for m := i + 1; m < j+1; m++ {
					jres = nums[m]*sqr(j-m) + jres
				}
				for m := j + 1; m < k+1; m++ {
					kres = nums[m]*sqr(k-m) + kres
				}
				for m := k + 1; m < len(nums); m++ {
					mres = nums[m]*sqr(len(nums)-m-1) + mres
				}
				//var resString = fmt.Sprintf("%d:%d:%d:%d", ires, jres, kres, mres)
				//println(resString)
				if inRange(ires) && inRange(jres) && inRange(kres) && inRange(mres) {
					var resString = fmt.Sprintf("%d:%d:%d:%d", ires, jres, kres, mres)
					println(resString)
				}
			}
		}
	}

}
func sqr(n int) int {
	sqrRes := 1
	for i := 0; i < n; i++ {
		sqrRes *= 10
	}
	return sqrRes
}
func inRange(a int) bool {
	if 255 >= a && a >= 0 {
		return true
	}
	return false
}
