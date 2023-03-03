package main

//四层循环（内两层为插入排序，外两层为划分不同的子序列）
func shellSort(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		shell(nums, gap)
	}
}

func shell(nums []int, gap int) {
	n := len(nums)
	for i := 0; i < gap; i++ {
		for j := i; j < n; j += gap {
			for k := j; k >= gap && nums[k] < nums[k-gap]; k -= gap {
				nums[k], nums[k-gap] = nums[k-gap], nums[k]
			}
		}
	}
}

func insertSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i; j >= 1 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
