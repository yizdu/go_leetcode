package leetcode

func bt46(list, nums []int, length int, result *[][]int) {
	if len(list) >= length {
		*result = append(*result, append([]int(nil), list...))
		return
	}
	for i := 0; i < len(nums); i++ {
		list = append(list, nums[i])
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt46(list, newNums, length, result)
		list = list[:len(list)-1]
	}
}

func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	bt46([]int{}, nums, n, &result)
	return result

}