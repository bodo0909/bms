package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func findBitmaps(m int, s_lim int) []bitmap {

	maxval := int(math.Pow(2.0, float64(m*m)))

	b_list := []bitmap{}

	start := time.Now()

	dist := make(map[int]int)
	b := NewBitmapFromInt(m, 0)
	for i := 0; i < maxval; i++ {
		b.Increment()
		s := b.S()
		dist[s] += 1

		if s >= s_lim {
			b_list = append(b_list, b.copy())
		}
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

	fmt.Println("total captured: ", len(b_list))

	return b_list
}

func reduceSetViaTranslation(m int, b_list []bitmap) []bitmap{
	t_max := (m + 1) / 2

	b_list_reduced := []bitmap{}

	for _, b := range b_list {
		all_translations_found := true
		TranslationLoop:
			for i := -t_max; i < t_max; i++ {
				for j := -t_max; j < t_max; j++ {
					if i != 0 || j != 0 {
						t_match := false
						for _, c := range b_list {
							if b.CompareTranslation(&c, i, j) {
								t_match = true
								break
							}
						}
						if !t_match {
							all_translations_found = false
							break TranslationLoop
						}
					}
				}
			}

		if all_translations_found {
			b_list_reduced = append(b_list_reduced, b.copy())
		}
	}

	fmt.Println("reduced to: ", len(b_list_reduced))
	return b_list_reduced
}


func main() {

	b := NewBitmapFromInt(3, 85)
	fmt.Println(b.ToString())
	fmt.Println(b.closestChecker())
	c := NewBitmapFromInt(3, 2+8+32)
	fmt.Println(c.ToString())
	fmt.Println(c.closestChecker())

	fmt.Println("matching translation: ", b.CompareTranslation(c, 0, -1))

	m := 5
	list := findBitmaps(m, -4)
	reduced_list := reduceSetViaTranslation(m, list)

	for _, d := range reduced_list {
		fmt.Println(d.ToString())
	}

	pattern := NewMosaic()
	pattern.set(0,0, *b)
	fmt.Println(pattern.ToString())
	pattern.set(0,0, *c)
	fmt.Println(pattern.ToString())
	pattern.set(0,1, *b)
	fmt.Println(pattern.ToString())

}
