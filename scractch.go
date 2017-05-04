package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	m := 5
	maxval := int(math.Pow(2.0, float64(m*m)))

	start := time.Now()

	dist := make(map[int]int)
	b := NewBitmapFromInt(m,0)
	for i := 0; i < maxval; i++ {
		b.Increment()
		s := b.S()
		dist[s] += 1
	}

	elapsed := time.Since(start)
	defer fmt.Printf("Execution took %s\n", elapsed)

	// print in sorted order
	var keys []int
	for k := range dist {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, " : ", dist[k])
	}
}
