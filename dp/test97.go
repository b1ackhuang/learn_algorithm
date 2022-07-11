package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	//memo[i][j]：s1[:i] s2[:j] 能否交错生成s3[i+j]
	var memo = make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		memo[i] = make([]bool, len(s2)+1)
	}
	memo[0][0] = true
	for k := 1; k <= len(s3); k++ {
		for i := 0; i <= len(s1) && i <= k; i++ {
			var j = k - i
			if j <= len(s2) {
				if i > 0 && s3[k-1] == s1[i-1] && memo[i-1][j] {
					memo[i][j] = true
				}
				if j > 0 && s3[k-1] == s2[j-1] && memo[i][j-1] {
					memo[i][j] = true
				}
			}
		}
	}
	fmt.Println(memo)
	return memo[len(s1)][len(s2)]
}
