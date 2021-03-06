package main

//import "fmt"

type bitmap struct {
	size   int
	values []bool
	right_neighbors []int
	down_neighbors []int
}

func NewBitmap(size int) *bitmap {
	b := new(bitmap)
	b.size = size
	b.values = make([]bool, size*size)
	b.right_neighbors = make([]int, 0)
	b.down_neighbors = make([]int, 0)

	return b
}

func NewBitmapFromInt(size int, val int) *bitmap {
	b := new(bitmap)
	b.size = size
	b.values = make([]bool, size*size)

	base2 := 2
	for i := 0; i < size*size; i++ {
		b.values[i] = val%base2 >= base2/2
		base2 *= 2
	}

	return b
}

func (b *bitmap) Increment() {
	for i := 0; i < len(b.values); i++ {
		b.values[i] = !b.values[i]
		if b.values[i] {
			break
		}
	}
}

func (b bitmap) S() int {
	s := 0

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size-1; j++ {
			tmp := i*b.size + j
			if b.values[tmp] == b.values[tmp+1] {
				s--
			}
			tmp = j*b.size + i
			if b.values[tmp] == b.values[tmp+b.size] {
				s--
			}
		}
	}
	return s
}

func (b bitmap) ToInt() int {
	base2 := 1
	tot := 0
	for _, j := range b.values {
		if j {
			tot += base2
		}
		base2 *= 2
	}
	return tot
}

func (b bitmap) ToString() string {
	str := ""
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.values[i*b.size+j] {
				str += "_"
			} else {
				str += "X"
			}
		}
		str += "\n"
	}
	return str
}

func (b bitmap) copy() bitmap {
	c := NewBitmap(b.size)
	for i, val := range b.values {
		c.values[i] = val
	}
	return *c
}

func (b bitmap) closestChecker() bool {
	sum := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if ((i + j%2) == 0) == b.values[i*b.size+j] {
				sum += 1
			} else {
				sum -= 1
			}
		}
	}
	return sum >= 0
}

func (b bitmap) bufferWithChecker() *bitmap {
	starting := b.closestChecker()
	c := NewBitmap(b.size + 2)
	ij := 0
	for i := 0; i < c.size; i++ {
		for j := 0; j < c.size; j++ {
			// checker border
			if i == 0 || i == c.size-1 || j == 0 || j == c.size-1 {
				if (i+j)%2 == 0 {
					c.values[ij] = starting
				} else {
					c.values[ij] = !starting
				}

				// inner section
			} else {
				c.values[ij] = b.values[(i-1)*b.size+j-1]
			}

			ij++
		}
	}
	return c
}

func (b bitmap) CompareTranslation(c *bitmap, delta_i int, delta_j int) bool {

	var i_lb, i_ub, j_lb, j_ub int
	if delta_i < 0 {
		i_lb = -delta_i
		i_ub = b.size
	} else {
		i_lb = 0
		i_ub = b.size - delta_i
	}
	if delta_j < 0 {
		j_lb = -delta_j
		j_ub = b.size
	} else {
		j_lb = 0
		j_ub = b.size - delta_j
	}

	// check overlap for similarity
	for i_b := i_lb; i_b < i_ub; i_b++ {
		i_c := i_b + delta_i
		for j_b := j_lb; j_b < j_ub; j_b++ {
			j_c := j_b + delta_j
			if b.values[i_b*b.size+j_b] != c.values[i_c*b.size+j_c] {
				return false
			}
		}
	}
	return true
}


func (b *bitmap) fillNeighbors() {

	//get info for right neighbors
	base2 := 1
	right_base := 0
	right_values := make([]int,0)
	for i, j := range b.values {
		if i%b.size > 0 {
			if j {
				right_base += base2
			}
			base2 *= 2
			if i%b.size == b.size-1 {
				right_values = append(right_values, base2)
				base2 *= 2
			}
		}
	}
	b.right_neighbors = getCombinations(right_base, right_values)

	//get info for down neighbors
	base2 = 1
	down_base := 0
	down_values := make([]int,0)
	for i, j := range b.values {
		if i>=b.size{
			if j {
				down_base += base2
			}
			base2 *= 2
		}
	}
	for i:=0; i<b.size; i++ {
		down_values = append(down_values, base2)
		base2 *= 2
	}
	b.down_neighbors = getCombinations(down_base, down_values)

}


func getCombinations(base int, values []int) []int{
	var list []int
	if len(values) == 1 {
		list = make([]int, 2)
		list[0] = base
		list[1] = base+values[0]
	} else {
		list = append(getCombinations(base, values[1:]), getCombinations(base+values[0], values[1:])...)
	}
	return list
}