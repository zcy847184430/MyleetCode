package leetCode_2019_11_07

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
https://leetcode-cn.com/problems/container-with-most-water/

2019.11.19
*/

/*
	1.嵌套遍历，算出所有的容积.O(N^2)
	2.[最优解]：左右双坐标，夹逼
*/


func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

//暴力解法
func maxArea1(height []int) int {
	result := int(0)
	for i := int(0) ;i < len(height) - 1; i++ {
		for j := i + 1; j < len(height); j++ {
			minHeight := min(height[i], height[j])
			result = max(minHeight * (j - i), result)
		}
	}
	return result
}

//最优解O(N)
func maxArea2(height []int) int {
	result := int(0)
	for i, j := int(0), len(height) - 1; i < j;  {
		result = max(result, min(height[i], height[j]) * (j - i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return result
}

func maxArea3(height []int) int {
	result := int(0)
	for i, j := int(0), len(height) - 1; i < j;  {
		result = max(result, min(height[i], height[j]) * (j - i))
		if height[i] > height[j] {
			for {
				k := height[j]
				j--
				if i > j || height[j] > k {
					break
				}
			}
		} else {
			for {
				k := height[i]
				i++
				if i > j || height[i] > k {
					break
				}
			}
		}

	}
	return result
}





































































