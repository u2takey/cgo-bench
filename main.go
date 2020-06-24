package main

import (
	"fmt"
	"time"
)

/*

void calSum(int c) {
	int sum = 0;
	for(int i=0; i<=c; i++ ){
        sum=sum+i;
    }
}

*/
// #cgo LDFLAGS: -lstdc++
import "C"

func calSum(c int) {
	sum := 0
	for i := 0; i <= c; i++ {
		sum += i
	}
}

func main() {
	cycles := []int{500000, 1000000, 5000000, 10000000}
	counts := []int{10, 50, 100, 500, 1000, 5000, 10000}
	for _, count := range counts {
		for _, cycle := range cycles {
			startCgo := time.Now()
			for i := 0; i < cycle; i = i + 1 {
				C.calSum(C.int(count))
			}
			costCgo := time.Now().Sub(startCgo)

			startGo := time.Now()
			for i := 0; i < cycle; i = i + 1 {
				calSum(count)
			}
			costGo := time.Now().Sub(startGo)

			fmt.Printf("count: %d, cycle: %d, cgo: %s, go: %s, cgo/cycle: %s, go/cycle: %s cgo/go: %.4f \n",
				count, cycle, costCgo, costGo, costCgo/time.Duration(cycle), costGo/time.Duration(cycle), float64(costCgo)/float64(costGo))
		}
	}
}
