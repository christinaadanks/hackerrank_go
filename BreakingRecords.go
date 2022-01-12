package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var min, max, temp, countMin, countMax int32 = scores[0], scores[0], 0, 0, 0
	for i := 0; i < len(scores); i++ {
		temp = scores[i]
		if temp < min {
			countMin++
			min = temp
		}
		if temp > max {
			countMax++
			max = temp
		}
	}
	result := []int32{countMax, countMin}
	return result
}
func main() {
	scores := []int32{10,5,20,20,4,5,2,25,1}
	fmt.Println(breakingRecords(scores))
	scores2 := []int32{3,4,21,36,10,28,35,5,24,42}
	fmt.Println(breakingRecords(scores2))
}
