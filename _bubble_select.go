package main

func selectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

func bubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
