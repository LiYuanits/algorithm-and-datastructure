package sort

import "fmt"

// 冒泡排序
// 左边跟右边元素比较如果左边的元素大就往右边移动

func BubblingSort(nums []int) {
	l := len(nums)
	// 遍历所有元素
	for i := 0; i < l-1; i++ {
		// 当前元素比较一轮所有元素
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Printf("冒泡排序排序好的结果 %v\n", nums)
}
