package model

import "fmt"

const EMPTY int = 0
const BLACK int = 1
const WHITE int = -1

// x-axis is left to right
// y-axis is front to back
// z-axis is bottom to top
type Cube struct {
	// in [z][y][z] manner
	Content [][][]int
	Size int
}

func (cube *Cube) Initialize(size int) {
	if size <= 0 {
		panic("size no more than zero!")
	}

	cube.clear()

	cube.Size = size
	cube.Content = make([][][]int, 0)
	for z := 0; z < size; z++ {
		area := make([][]int, 0)
		for y := 0; y < size; y++ {
			line := make([]int, 0)
			for x := 0; x < size; x++ {
				line = append(line, EMPTY)
			}
			area = append(area, line)
		}
		cube.Content = append(cube.Content, area)
	}
}

func (cube *Cube) clear() {
	cube.Size = 0
	cube.Content = nil
}

func (cube *Cube) Print() {
	if cube.Size <= 0 {
		fmt.Println("empty cube")
		return
	}

	for z := cube.Size - 1; z >= 0; z-- {
		fmt.Printf("----- z %d -----\n", z)

		for y := cube.Size - 1; y >= 0; y-- {
			for x := 0; x < cube.Size; x++ {
				fmt.Printf("%d ", cube.Content[z][y][x])
			}
			fmt.Printf("\n")
		}

		fmt.Printf("\n")
	}
}



