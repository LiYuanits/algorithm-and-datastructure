package sort

import "fmt"

// 插入排序
// 在已排序序列中从后向前扫描，找到相应位置并插入

func InsertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {

		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {

			nums[j], nums[j-1] = nums[j-1], nums[j]
		}

	}
	fmt.Printf("插入排序好结果 %v\n", nums)
}
