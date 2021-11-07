package main

import "fmt"

func main() {

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	var num int
	length := len(height)

	leftMaxArr := make([]int, length)
	rightMaxArr := make([]int, length)
	leftMaxArr[0] = height[0]
	rightMaxArr[length-1] = height[length-1]

	for i := 1; i < length; i++ {
		if leftMaxArr[i-1] > height[i] {
			leftMaxArr[i] = leftMaxArr[i-1]
		} else {
			leftMaxArr[i] = height[i]
		}
	}
	for i := length - 2; i >= 0; i-- {
		if rightMaxArr[i+1] > height[i] {
			rightMaxArr[i] = rightMaxArr[i+1]
		} else {
			rightMaxArr[i] = height[i]
		}
	}
	for i := 0; i < length; i++ {
		var minNum int

		if leftMaxArr[i] < rightMaxArr[i] {
			minNum = leftMaxArr[i]
		} else {
			minNum = rightMaxArr[i]
		}

		num += minNum - height[i]
	}
	return num
}
