package main

import "fmt"

func main() {

	arr := []int{1, 9, 8}
	fmt.Println(QuickSort(arr))
}

func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	splitData := arr[0]
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //跟我一样大的数据
	mid = append(mid, splitData) //将第一数据加入进来

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, hight = QuickSort(low), QuickSort(hight)

	myArr := append(append(low, mid...), hight...)
	return myArr
}

// 官方题解
//func QuickSort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//	splitdata := arr[0]          //第一个数据
//	low := make([]int, 0, 0)     //比我小的数据
//	hight := make([]int, 0, 0)   //比我大的数据
//	mid := make([]int, 0, 0)     //与我一样大的数据
//	mid = append(mid, splitdata) //加入一个
//	for i := 1; i < len(arr); i++ {
//		if arr[i] < splitdata {
//			low = append(low, arr[i])
//		} else if arr[i] > splitdata {
//			hight = append(hight, arr[i])
//		} else {
//			mid = append(mid, arr[i])
//		}
//	}
//	low, hight = QuickSort(low), QuickSort(hight)
//	myarr := append(append(low, mid...), hight...)
//	return myarr
//}
