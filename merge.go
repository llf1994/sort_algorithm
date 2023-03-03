package main

func mergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	return merge(mergeSort(nums[:n/2]), mergeSort(nums[n/2:]))
}

func merge(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	ans := make([]int, n1+n2)
	i, j := 0, 0
	for index := 0; index < n1+n2; index++ {
		if i < n1 && j < n2 {
			if nums1[i] < nums2[j] {
				ans[index] = nums1[i]
				i++
			} else {
				ans[index] = nums2[j]
				j++
			}
		} else if i < n1 {
			ans[index] = nums1[i]
			i++
		} else if j < n2 {
			ans[index] = nums2[j]
			j++
		}
	}
	return ans
}
