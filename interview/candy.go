package main

import "fmt"
import "sort"

func candy(ratings []int) (int) {
	var count []int = make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		count[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i - 1] {
			count[i] = count[i - 1] + 1
		}
	}
	var sum int = 0
	for i := len(ratings) - 1; i >= 1; i-- {
		sum += count[i]
		if ratings[i - 1] > ratings[i] && 
			count[i - 1] <= count[i] {
			count[i - 1] = count[i] + 1
		} 
	}
	sum += count[0]
	return sum
}

func main() {
    p := []int{2, 3, 7, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    var num = candy(p)
    fmt.Println("total candy:", num)
    sort.Ints(p)
    fmt.Println("after sorting: ==", p)
}