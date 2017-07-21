package main

import (
	"fmt"
	"math"
	//"sort"
	"time"
)

func findBitmaps(m int, s_lim int) map[int]bitmap {

	fmt.Println("Finding bitmaps which satisfy S constraint...")
	start := time.Now()

	maxval := int(math.Pow(2.0, float64(m*m)))

	b_map := make(map[int]bitmap)

	dist := make(map[int]int)
	b := NewBitmapFromInt(m, 0)
	for i := 0; i < maxval; i++ {
		b.Increment()
		s := b.S()
		dist[s] += 1

		if s >= s_lim {
			//fmt.Println(b.ToInt())
			b_map[b.ToInt()] = b.copy()
		}
	}

	/*
	// print in sorted order
	var keys []int
	for k := range dist {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, " : ", dist[k])
	}
	*/

	fmt.Println("  total found: ", len(b_map))
	fmt.Printf("  execution took: %s\n", time.Since(start))
	return b_map
}

func reduceSetViaTranslation(m int, b_map map[int]bitmap) map[int]bitmap{
	fmt.Println("Reducing to bitmaps which have possible neighbors in the list...")
	start := time.Now()

	t_max := (m + 1) / 2

	b_map_reduced := make(map[int]bitmap)

	for key, b := range b_map {
		all_translations_found := true
		TranslationLoop:
			for i := -t_max; i < t_max; i++ {
				for j := -t_max; j < t_max; j++ {
					if i != 0 || j != 0 {
						t_match := false
						for _, c := range b_map {
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
			b_map_reduced[key] = b.copy()
		}
	}

	fmt.Println("  reduced to: ", len(b_map_reduced))
	fmt.Printf("  execution took: %s\n", time.Since(start))
	return b_map_reduced
}

func findNeighborsInList(b_map map[int]bitmap) map[int]bitmap{
	fmt.Println("Finding bitmap neighbors which are in the list...")
	start := time.Now()

	for _, b := range b_map {
		b.fillNeighbors()
		for d_n := range b.down_neighbors {
			if _, found := b_map[d_n]; !found {
				delete(b.down_neighbors, d_n)
			}
		}
		fmt.Println(d.ToInt())
	}

	fmt.Println("  reduced to: ", len(b_map_reduced))
	fmt.Printf("  execution took: %s\n", time.Since(start))
	return b_map_reduced
}


func main() {

	/*
	b := NewBitmapFromInt(3, 85)
	fmt.Println(b.ToString())
	fmt.Println(b.closestChecker())
	c := NewBitmapFromInt(3, 2+8+32)
	fmt.Println(c.ToString())
	fmt.Println(c.closestChecker())

	fmt.Println("matching translation: ", b.CompareTranslation(c, 0, -1))
`	*/

	m := 5
	list := findBitmaps(m, -4)
	reduced_list := reduceSetViaTranslation(m, list)

	for _, d := range reduced_list {
		fmt.Println(d.ToInt())
	}

	/*
	pattern := NewMosaic()
	pattern.set(0,0, *b)
	fmt.Println(pattern.ToString())
	pattern.set(0,0, *c)
	fmt.Println(pattern.ToString())
	pattern.set(1,1, *b)
	fmt.Println(pattern.ToString())
	*/

	/*
	fmt.Println(b.ToString())
	b.fillNeighbors()
	fmt.Println("rn: ",len(b.right_neighbors))

	fmt.Println("b:--------")
	fmt.Println(b.ToString())
	for indx, i := range b.right_neighbors {
		c = NewBitmapFromInt(b.size, i)
		fmt.Println(indx," neighbor:")
		fmt.Println(c.ToString())
	}
	*/


}
