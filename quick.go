package main

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	p := nums[0]
	left, right := 0, n-1
	for left < right {
		for left < right && nums[right] >= p {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= p {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = p
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}
