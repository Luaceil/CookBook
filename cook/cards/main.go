package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([]int, 0)
	counts := 1
	scores := 0
	for i := 0; i < n; i++ {
		var count int
		var score int
		fmt.Scan(&score, &count)
		if count > 0 {
			scores += score
			counts += count - 1
		} else {
			cards = append(cards, score)
		}
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i] < cards[j] {
			return false
		} else {
			return true
		}
	})
	for i := 0; i < counts && i < len(cards); i++ {
		scores += cards[i]
	}

	fmt.Println(scores)

}
