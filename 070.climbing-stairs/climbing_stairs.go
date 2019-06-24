package main

import "fmt"

/**
动态规划法
第 ii 阶可以由以下两种方法得到：
	在第 (i-1) 阶后向上爬一阶。
	在第 (i-2) 阶后向上爬 2 阶。

所以到达第 ii 阶的方法总数就是到第 (i−1) 阶和第 (i−2) 阶的方法数之和。

令 dp[i] 表示能到达第 ii 阶的方法总数：

dp[i]=dp[i-1]+dp[i-2]

 */

func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i:= 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

/**
斐波那契数
算法

在上述方法中，我们使用 dpdp 数组，其中 dp[i]=dp[i-1]+dp[i-2]。可以很容易通过分析得出 dp[i] 其实就是第 ii 个斐波那契数。

Fib(n)=Fib(n-1)+Fib(n-2)

现在我们必须找出以 1 和 2 作为第一项和第二项的斐波那契数列中的第 n 个数，也就是说 Fib(1)=1 且 Fib(2)=2。


*/
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	sum := 0

	jump1 := 1
	jump2 := 2

	for i:= 3; i <= n; i++{
		sum = jump1 + jump2
		jump1 = jump2
		jump2 = sum
	}

	return sum
}

func main()  {
	fmt.Println(climbStairs1(4))
	fmt.Println(climbStairs2(4))
}