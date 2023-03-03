package main

import (
	"math"
	"sort"
)

func bucketSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	maxNum, minNum := nums[0], nums[0]
	for _, v := range nums {
		maxNum, minNum = max(maxNum, v), min(minNum, v)
	}
	cnt := 10
	gap := (maxNum - minNum) / cnt
	buckets := make([][]int, cnt+1)
	for _, v := range nums {
		index := v / gap
		buckets[index] = append(buckets[index], v)
	}
	index := 0
	for i := range buckets {
		if len(buckets) > 1 {
			sort.Ints(buckets[i])
			for _, v := range buckets[i] {
				nums[index] = v
				index++
			}
		}
	}

}
func RadixSortMSD(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	maxNum := nums[0]
	for _, v := range nums {
		maxNum = max(maxNum, v)
	}
	digits := getDigits(maxNum)
	MSDRadix(nums, digits)
}

func MSDRadix(s []int, radix int) {
	exp, buckets, hasSorted := int(math.Pow10(radix-1)), make([][]int, 10), true
	// hasSorted检验数组内的元素是否已排序好
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			hasSorted = false
			break
		}
	}
	if hasSorted {
		return
	} // 如果元素全部相同，直接返回

	// 将数放入桶内
	for _, v := range s {
		n := v / exp % 10 // exp位的数，先整除再求余
		buckets[n] = append(buckets[n], v)
	}

	// 桶内只有1个数时直接合并，否则递归
	index := 0
	for _, bucket := range buckets {
		if len(bucket) > 1 {
			MSDRadix(bucket, radix-1) // 桶内递归
			for _, v := range bucket {
				s[index] = v
				index++
			}
		} else if len(bucket) == 1 {
			s[index] = bucket[0]
			index++
		}
	}
	return
}

func RadixSortLSD(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	maxNum := nums[0]
	for _, v := range nums {
		maxNum = max(maxNum, v)
	}
	buckets := make([][]int, 10)
	for exp := 1; maxNum/exp > 0; exp *= 10 {
		for _, v := range nums {
			index := v / exp % 10
			buckets[index] = append(buckets[index], v)
		}
		index := 0
		for i := range buckets {
			for len(buckets[i]) > 0 {
				nums[index], buckets[i] = buckets[i][0], buckets[i][1:]
				index++
			}
		}
	}
}

func getDigits(num int) int {
	res := 0
	for num > 0 {
		res++
		num /= 10
	}
	return res
}

func countSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	maxNum, minNum := nums[0], nums[0]
	for _, v := range nums {
		maxNum, minNum = max(maxNum, v), min(minNum, v)
	}
	bucket := make([]int, maxNum-minNum+1)
	for _, v := range nums {
		bucket[v-minNum]++
	}
	index := 0
	for i := range bucket {
		for bucket[i] > 0 {
			nums[index] = i + minNum
			index++
			bucket[i]--
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
