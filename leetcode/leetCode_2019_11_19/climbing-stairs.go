package leetCode_2019_11_19
/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。
	https://leetcode-cn.com/problems/climbing-stairs/
*/

/*
思路：
	1：1
	2: 1.2 ->2
	3. f(1) + f(2)
	可得，结果为斐波那契数列
*/

func climbStairs(n int) int {
	s := make([]int, n+2)
	s[1] = 1
	s[2] = 2
	for i := int(3); i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}