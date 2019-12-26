package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	result := make([][]int, len(image))
	if len(image) == 0 {
		return result
	}

	for i := 0; i < len(image); i++ {
		result[i] = make([]int, len(image[i]))
		copy(result[i], image[i])
	}

	fmt.Printf("result: %v\n", result)
	queue := make([][]int, 0, len(image)*len(image[0]))

	startingImage := image[sr][sc]

	if startingImage == newColor {
		return result
	}

	queue = append(queue, []int{sr, sc})

	for len(queue) != 0 {

		points := queue[0]
		queue = queue[1:]

		result[points[0]][points[1]] = newColor

		if points[0] > 0 && result[points[0]-1][points[1]] == startingImage {
			queue = append(queue, []int{points[0]-1, points[1]})
		}

		if points[0] < len(result)-1 && result[points[0]+1][points[1]] == startingImage {
			queue = append(queue, []int{points[0]+1, points[1]})
		}

		if points[1] > 0 && result[points[0]][points[1]-1] == startingImage {
			queue = append(queue, []int{points[0], points[1]-1})
		}

		if points[1] < len(image[0])-1 && result[points[0]][points[1]+1] == startingImage {
			queue = append(queue, []int{points[0], points[1]+1})
		}
	}

	return result
}

func main() {
	fmt.Printf("image %v\n", floodFill([][]int{{0,0,0}, {0,1,1}}, 1,1,2))
}
