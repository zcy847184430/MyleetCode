package leetCode_2019_11_07

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
	https://leetcode-cn.com/problems/move-zeroes/
	2019.11.19
*/

/*
思路：
	1. 两次循环，第一次数零的个数，第二次调整位置
	2. 开新数组，循环一次，0 -》从后往前，非0 ->从前往后
	3. [最优解]:一次遍历，两个索引，其中一个指向下一个非0元素的目标。
*/

//leetCode_2019_11_07/11/17
func moveZeroes1(nums []int)  {
	j := int(0)
	for i := int(0); i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}

//leetCode_2019_11_07/11/17
func moveZeroes2(nums []int) {
	i, j := int(0), int(0)
	for  ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

















































































