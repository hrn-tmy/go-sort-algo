package main

import (
	"fmt"
)

func main() {
	ascNums := []int{3, 1, 4, 2, 6, 5, 9, 7, 8}
	fmt.Println("ascNums Before: ", ascNums)
	for i := 0; i < len(ascNums); i++ { // ラウンド数
		for j := 0; j < len(ascNums)-1; j++ { // 実際の処理(昇順処理)
			if ascNums[j] > ascNums[j+1] {
				ascNums[j], ascNums[j+1] = ascNums[j+1], ascNums[j]
			}
		}
	}
	fmt.Println("ascNums After: ", ascNums)

	descNums := []int{3, 1, 4, 2, 6, 5, 9, 8, 7}
	fmt.Println("descNums Before: ", descNums)
	for i := 0; i < len(descNums); i++ { // ラウンド数
		for j := len(descNums) - 1; j > 0; j-- { // 実際の処理(降順処理)
			if descNums[j] > descNums[j-1] {
				descNums[j], descNums[j-1] = descNums[j-1], descNums[j]
			}
		}
	}
	fmt.Println("descNums After: ", descNums)
}
