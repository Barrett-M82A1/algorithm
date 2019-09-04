package main

import "fmt"

func main() {
	data := [9]int{9, 6, 1, 3, 7, 2, 8, 4, 5}

	pointer := 0

	for j := 0; j <= len(data)-1; j++ {
		for i := j + 1; i <= len(data)-2; i++ {
			if data[i] < data[pointer] {
				pointer = i
			}
		}

		data[j], data[pointer] = data[pointer], data[j]
	}

	fmt.Println(data)
}
