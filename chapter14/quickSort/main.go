package main

func sortArray(nums []int) []int {
	length := len(nums)
	r := length - 1
	l := 0

	quickSort(nums, l, r) //快排，你真是一位老朋友
	return nums
}

// 没有ctrl+shift+enter，没有.sout，再也没有
func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	l := start
	r := end
	benchMark := nums[l]

	for l < r {
		for l < r && nums[r] >= benchMark {
			r--
		}
		for l < r && nums[l] <= benchMark {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[start] = nums[start], nums[l]

	quickSort(nums, start, l-1)
	quickSort(nums, l+1, end)

}
