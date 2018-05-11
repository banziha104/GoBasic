package main

import (
	"fmt"
)

func main() {
	fruits := []string{"사과","바나나","토마토"}
	fmt.Println(fruits[:2]) // 슬라이싱
	fruits = append(fruits,"포도","딸기") //슬라이싱 덧붙이기
	fmt.Println(cap(fruits)) // 용량
	fmt.Println(len(fruits)) // 길이
	nums := make([]int, 0 ,15) // 15개 까지만 가능한 슬라이스
	nums = append(nums, 1)

	dest := make([]string,len(fruits))
	copy(dest,fruits) // 복사


}

