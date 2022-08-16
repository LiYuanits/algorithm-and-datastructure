package main

import (
	"algorithm-and-datastructure/sort"
	"fmt"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 9, 3, 7}

	fmt.Printf("待排序的数据 %v\n", nums)

	sort.QuickSort(nums, 0, len(nums)-1)

}

func searchInsert(nums []int, target int) int {
	var ret int
	count := len(nums)
	for key, value := range nums {
		if value == target {
			return key
		} else if value < target && key == count-1 {
			ret = key + 1
			break
		} else if value > target {
			ret = key
			break
		}
	}
	return ret
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// n1有值，n2没值
	if m > 0 && n == 0 {
		return
	}

	// n2有值
	if m == 0 && n > 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
			m++
		}
		return
	}
	// 头跟中间插入元素 要把元素往后面移动

	for _, v2 := range nums2 {

		for k, v3 := range nums1 {
			// 头跟中间插入
			if v3 == v2 || v3 > v2 {
				for i := m - 1; i >= k; i-- {
					nums1[i+1] = nums1[i]
				}
				nums1[k] = v2
				m++
				break
			}

			// 尾巴插入
			if v3 < v2 && k == m-1 {
				nums1[k+1] = v2
				m++
				break
			}
		}
	}
	fmt.Println(nums1)
}
