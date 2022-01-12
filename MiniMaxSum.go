package main

import "fmt"

func miniMaxSum(arr []int32) {
	var min, max, sum, temp int64 = 0, 0, 0, 0
	for i := 0; i < len(arr); i++ {
		temp += int64(arr[i])
	}
	sum = temp
	for j := 0; j < len(arr); j++	{
		sum -= int64(arr[j])
		if min == 0 {
			min = sum
		}
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
		sum = temp
	}
	fmt.Println(min, max)
}
func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)
}
