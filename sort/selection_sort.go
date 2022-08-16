package sort

import "fmt"

// 选择排序
// 每次从带排序的数组中选出一个最大或者最小的元素放在排序的前面

func SelectionSort(nums []int) {
	n := len(nums)
	// 遍历元素
	for i := 0; i < n; i++ {
		// 寻找最小的元素
		minIndex := i
		for j := i; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 把最小的元素放到最前面，把当前[i]位置的元素放到 最小元素的原来位置
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}
	fmt.Printf("选择排序排序好的结果 %v\n", nums)
}

// 选择排序第二种写法
func SelectSort2(arr []int) {

	for i := 0; i < len(arr)-1; i++ {

		//第二层for遍历比第一层加1，如果第二层for遍历到的值小许第一次就替换位置
		for j := i + 1; j <= len(arr)-1; j++ {

			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}

		}

	}
	fmt.Printf("选择排序排序好的结果 %v\n", arr)
}
