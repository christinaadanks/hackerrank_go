package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversions (s string) string {
	var time string = ""
	arr := strings.Split(s, ":")
	if strings.Contains(arr[2], "PM") {
		arr[2] = strings.Replace(arr[2], "PM", "", -1)
		if arr[0] == "12" {
			time = arr[0] + ":" + arr[1] + ":" + arr[2]
		} else {
			hour, _ := strconv.Atoi(arr[0])
			hour += 12
			strHour := strconv.Itoa(hour)
			time = strHour + ":" + arr[1] + ":" + arr[2]
		}
	}
	if strings.Contains(arr[2], "AM") {
		arr[2] = strings.Replace(arr[2], "AM", "", -1)
		if arr[0] == "12" {
			time = "00:" + arr[1] + ":" + arr[2]
		} else {
			time = arr[0] + ":" + arr[1] + ":" + arr[2]
		}
	}
	return time
}

func main() {
	s := "07:05:45PM"
	fmt.Println(timeConversions(s))
}
