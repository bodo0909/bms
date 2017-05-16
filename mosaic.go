package main

import "strconv"

type mosaic struct {
	size_i   int
	size_j   int
	bitmaps [][]bitmap
}

func NewMosaic() *mosaic {
	m := new(mosaic)
	m.size_i = 0
	m.size_j = 0
	m.bitmaps = make([][]bitmap, 0)

	return m
}

func (m *mosaic) set(I int, J int, b bitmap) {
	//expand if necessary
	if J >= m.size_j {
		for i := range m.bitmaps{
			m.bitmaps[i] = append(m.bitmaps[i], make([]bitmap, J-m.size_j+1)...)
		}
		m.size_j = J+1
	}
	if I >= m.size_i {
		for i := m.size_i; i < I; i++ {
			m.bitmaps = append(m.bitmaps, make([]bitmap, m.size_j))
		}
		m.size_i = I+1
	}

	//set value
	m.bitmaps[I][J] = b
}

func (m *mosaic) ToString() string {
	str := "~~~\n"
	for i := range m.bitmaps {
		for j := range m.bitmaps[i] {
			str += strconv.Itoa(m.bitmaps[i][j].ToInt())
			str += "\t"
		}
		str += "\n"
	}
	return str
}
