package algos

import "math/rand"
import "time"
import "fmt"

func PopulatearrayIntn(arr *[]int) {
	
	var Len int = len(*arr)
	bench_startTime := time.Now()
	defer func() {
		fmt.Println("Populate array", Len, "Time took", time.Since(bench_startTime))
	}()

	rand.Seed(time.Now().UnixNano())

	var elemMap map[int]bool = make(map[int]bool)

	check := func (elem int) bool {
		_, presence := elemMap[elem]
		return presence
	}
	
	for i := range *arr {
		for {
			randInt := rand.Intn(Len)
			if !check(randInt) {
				elemMap[randInt] = true
				(*arr)[i] = randInt
				break
			}
		}
	}
}

