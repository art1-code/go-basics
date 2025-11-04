package main

import "fmt"

func main() {

	slice := make([]int, 5)

	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}

	fmt.Println(slice)
	fmt.Println(slice[2:5])
	fmt.Println(slice[0:3])

}
