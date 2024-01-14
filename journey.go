package main

import (
	"fmt"
	"localSht/algos"
	"time"
)

func main() {
	test_generic_popluate()
}

func test_generic_popluate() {
	sizes := []int{20, 200, 2000, 20000, 200000, 2000000, 1 << 22}
	fmt.Println("=> Populate array with non repeating integers")
	fmt.Println("=> performing tests on sizes", sizes)

	start_time := time.Now()
	defer func() {
		fmt.Println("=> Test took ", time.Since(start_time))
	}()
	for _, size := range sizes {
		var arr = make([]int, size)
		algos.PopulatearrayIntn(&arr)
	}
}
