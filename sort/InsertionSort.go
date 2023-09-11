package sort

// type T sort.Interface
// 插入排序：nums[0..i-1]是有序的，对nums[i]，将其插入合适的位置
// 然后nums[0..i]即为有序的
func InsertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := i
		for ; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}
