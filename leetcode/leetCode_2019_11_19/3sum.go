package leetCode_2019_11_19
/*
	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
	使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
	https://leetcode-cn.com/problems/3sum/
*/
/*
	1. 暴力求解，O(N^3)
	2. 左右下标推进
*/

func threeSum1(nums []int) [][]int {
	result := make([][]int, 0)
	for i := int(0); i < len(nums) - 2; i++ {
		for j := i + 1; j < len(nums) - 1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					result = append(result, temp)
				}
			}
		}
	}

	return result
}