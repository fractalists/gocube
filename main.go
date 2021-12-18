package main

import (
	"fmt"
	"scorefour/model"
)

const SIZE = 4
const BLOCK_EMPTY = 0
const BLOCK_BLACK = 1
const BLOCK_WHITE = -1

var THIRTEEN_DIRECTION = map[string][][]int{
	"UP_TO_DOWN":[][]int{[]int{1, 0, 0}, []int{-1, 0, 0}},
	"IN_TO_OUT":[][]int{[]int{0, 1, 0}, []int{0, -1, 0}},
	"LEFT_TO_RIGHT":[][]int{[]int{0, 0, -1}, []int{0, 0, 1}},
	"UP-LEFT_TO_DOWN-RIGHT":[][]int{[]int{1, 0, -1}, []int{-1, 0, 1}},
	"UP-RIGHT_TO_DOWN-LEFT":[][]int{[]int{1, 0, 1}, []int{-1, 0, -1}},
	"IN-LEFT_TO_OUT-RIGHT":[][]int{[]int{0, 1, -1}, []int{0, -1, 1}},
	"IN-RIGHT_TO_OUT-LEFT":[][]int{[]int{0, 1, 1}, []int{0, -1, -1}},
	"UP-IN_TO_DOWN-OUT":[][]int{[]int{1, 1, 0}, []int{-1, -1, 0}},
	"UP-OUT_TO_DOWN-IN":[][]int{[]int{1, -1, 0}, []int{-1, 1, 0}},
	"UP-OUT-LEFT_TO_DOWN-IN-RIGHT":[][]int{[]int{1, -1, -1}, []int{-1, 1, 1}},
	"UP-OUT-RIGHT_TO_DOWN-IN-LEFT":[][]int{[]int{1, -1, 1}, []int{-1, 1, -1}},
	"UP-IN-LEFT_TO_DOWN-OUT-RIGHT":[][]int{[]int{1, 1, -1}, []int{-1, -1, 1}},
	"UP-IN-RIGHT_TO_DOWN-OUT-LEFT":[][]int{[]int{1, 1, 1}, []int{-1, -1, -1}},
}

func main() {
	cube := model.Cube{}
	cube.Initialize(SIZE)

	for level := 0; level < SIZE; level++ {
		for height := 0; height < SIZE; height++ {
			for length := 0; length < SIZE; length++ {
				calcAndSetValue(&cube.Content, []int{level, height, length})
			}
		}
	}
	//
	//calcAndSetValue(&cube, []int{0, 0, 0})

	cube.Print()
}


func calcAndSetValue(cube *[][][]int, coordinate []int) {
	count := 0

	for _, directionList := range THIRTEEN_DIRECTION {
		directionOne := directionList[0]
		directionTwo := directionList[1]

		coordinateOne := moveAlongTheDirection(coordinate, directionOne)
		coordinateTwo := moveAlongTheDirection(coordinate, directionTwo)

		distance := distance(coordinateOne, coordinateTwo)
		if distance == SIZE - 1 {
			count++
		}

		//if distance == SIZE - 1 {
		// fmt.Printf("%s ok \n", directionName)
		// count++
		//} else {
		// fmt.Printf("%s not ok \n", directionName)
		//}
	}

	(*cube)[coordinate[0]][coordinate[1]][coordinate[2]] = count
}

func moveAlongTheDirection(coordinate []int, direction []int) []int {
	level := coordinate[0]
	height := coordinate[1]
	length := coordinate[2]

	if level < 0 || level >= SIZE || height < 0 || height >= SIZE || length < 0 || length >= SIZE {
		panic(fmt.Sprintf("invalid coordinate: %d %d %d", level, height, length))
	}

	for level >= 0 && level < SIZE && height >= 0 && height < SIZE && length >= 0 && length < SIZE {
		level += direction[0]
		height += direction[1]
		length += direction[2]
	}

	return []int{level - direction[0], height - direction[1], length - direction[2]}
}

func distance(a, b []int) int {
	return max([]int{abs(a[0], b[0]), abs(a[1], b[1]), abs(a[2], b[2])})
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}

	return b - a
}

func max(nums []int) int {
	if len(nums) == 0 {
		panic("max input nums is empty")
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}