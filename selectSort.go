package main

import "fmt"

func main() {
	data := [9]int{5, 6, 1, 3, 7, 4, 8, 2, 9}

	var max, min int

	for j := 0; j <= len(data)/2-1; j++ {

		max, min = j, j

		for i := j; i <= len(data)-j-2; i++ {
			if data[max] < data[i+1] {
				max = i + 1
			}

			if data[min] > data[i+1] {
				min = i + 1
			}
		}

		data[j], data[min] = data[min], data[j]

		data[len(data)-j-1], data[max] = data[max], data[len(data)-j-1]
	}

	fmt.Println(data)
}
