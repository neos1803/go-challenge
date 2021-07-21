package main

import "fmt"

func main() {

	fmt.Println(leapYear(2000, 2016))

}

func leapYear(start int, end int) []int {
	output := []int{}

	for i := start; i <= end; i++ {
		if i%4 == 0 {
			output = append(output, i)
		}
	}

	return output

}
