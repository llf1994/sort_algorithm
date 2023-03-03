package main

import "fmt"

func main() {
	array := []int{33, 733, 733, 44322, 733, 47565, 14, 64645254, 67, 9, 50, 64, 36, 2, 5, 36, 74}
	fmt.Println(len(array))
	heapSort(array)
	fmt.Println(array)
}

func heapSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	//建立最大堆
	for i := n/2 - 1; i >= 0; i-- {
		sinkDown(nums, i, n-1)
	}
	//不断上浮最大数
	for i := n - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sinkDown(nums, 0, i-1)
	}
}

func sinkDown(nums []int, dad, end int) {
	son := 2*dad + 1
	for son <= end {
		if son+1 <= end && nums[son] < nums[son+1] {
			son++
		}
		if nums[son] <= nums[dad] {
			return
		}
		nums[son], nums[dad] = nums[dad], nums[son]
		dad, son = son, 2*son+1
	}
}
