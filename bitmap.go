package main

type bitmap struct {
	size   int
	values []bool
}

func NewBitmap(size int) *bitmap {
	b := new(bitmap)
	b.size = size
	b.values = make([]bool, size*size)

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
			tmp := i*b.size+j
			if b.values[tmp] == b.values[tmp+1] {
				s--
			}
			tmp = j*b.size+i
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
		base2 *=2
	}
	return tot
}


func (b bitmap) ToString() string {
	str := ""
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.values[i*b.size+j]{
				str += "_"
			} else {
				str += "X"
			}
		}
		str += "\n"
	}
	return str
}