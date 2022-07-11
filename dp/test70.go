package main

func climbStairs(n int) int {
	//if n == 0 {
	//	return 1
	//}
	//if n == 1 {
	//	return 1
	//}
	//return climbStairs(n-1) + climbStairs(n-2)
	var mem = make([]int, n+1)
	mem[0] = 1
	mem[1] = 1
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}
