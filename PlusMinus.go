package main
/**
@author christinaathecoder
@version 1/11/2022
 */
import "fmt"

func plusMinus (arr []int32)	{
	var countNeg, countPos, countZ int = 0,0,0
	for i := 0; i < len(arr); i++ {
		if num := arr[i]; num < 0 {
			countNeg++
		} else if num == 0 {
			countZ++
		} else if arr[i] > 0 {
			countPos++
		}
	}
	pos := float64(countPos)/float64(len(arr))
	fmt.Printf("%.6f", pos)
	fmt.Println()
	neg := float64(countNeg)/float64(len(arr))
	fmt.Printf("%.6f", neg)
	fmt.Println()
	zero := float64(countZ)/float64(len(arr))
	fmt.Printf("%.6f", zero)
}
func main() {
	arr := []int32{-4,3,-9,0,4,1}
	plusMinus(arr)
}

