package main

import "fmt"

func main() {
	fmt.Println(groupingNumbers(0, 1000))
}

func checkPrimeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func groupingNumbers(start int, end int) map[string][]int {
	obj := make(map[string][]int)

	if end >= 2 {
		obj["primeNumbers"] = append(obj["primeNumbers"], 2)
		obj["primeNumbersLess100"] = append(obj["primeNumbersLess100"], 2)
		for i := start; i <= end; i++ {
			if i%2 == 0 {
				obj["evenList"] = append(obj["evenList"], i)
			}
			if i%2 == 1 {
				obj["oddList"] = append(obj["oddList"], i)
			}
			if i%5 == 0 {
				obj["dividedBy5"] = append(obj["dividedBy5"], i)
			}
			if i > 2 {
				if checkPrimeNumber(i) {
					obj["primeNumbers"] = append(obj["primeNumbers"], i)
					if i < 100 {
						obj["primeNumbersLess100"] = append(obj["primeNumbersLess100"], i)
					}
				}
			}
		}
	}

	return obj
}
