package main

import (
	"fmt"
	"strings"
)

func main() {

	data := `NAME, CATEGORY, PRICE
	Xiaomi Redmi 5A, Smartphone, 1199000
	Macbook Air, Laptop, 13775000
	Samsung Galaxy J7, Smartphone, 3549000
	DELL XPS 13, Laptop, 26799000
	Xiaomi Mi 6, Smartphone, 5399000
	LG V30 Plus, Smartphone, 10499000`

	row := strings.Split(data, "\n")
	key := strings.Split(row[0], ", ")
	output := []map[string]string{}

	for i := 1; i < len(row); i++ {
		object := make(map[string]string)

		column := strings.Split(row[i], ", ")

		for j := 0; j < len(column); j++ {
			object[strings.ToLower(key[j])] = column[j]
		}

		output = append(output, object)
	}

	fmt.Println(output)
	// fmt.Println(output[0]["name"])

}
