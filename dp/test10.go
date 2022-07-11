package main

import "fmt"

func isMatch(s string, p string) bool {
	var canRepeat []bool
	canRepeat, p = simplifyP(p)

	var memo = make([][]bool, len(p))
	for i := 0; i < len(p); i++ {
		memo[i] = make([]bool, len(s))
	}
	memo[0][0] = canRepeat[0]
	memo[0][1] = p[0] == '.' || s[0] == p[0]
	if canRepeat[0] {
		for i := 1; i < len(s); i++ {
			if p[0] == '.' || s[i] == p[0] {
				memo[0][i] = memo[0][i-1]
			}
		}
	}
	for i := 0; i < len(p) && canRepeat[i]; i++ {
		memo[i][0] = true
	}
	fmt.Println(memo)

	for i := 1; i < len(p); i++ {
		for j := 1; j < len(s); j++ {
			//case1: 后面不带*：(aa  ab)，(aa a.)
			if !canRepeat[i] {
				memo[i][j] = (p[i] == '.' || s[j] == p[i]) && memo[i-1][j-1]
				continue
			}

			//case2: 后面带*：(aa a*)，(aa *.)
			memo[i][j] = memo[i][j]   //*a 一个都不取
			for k := 1; k <= j; k++ { //*a 取k个
				if p[i] != '.' && s[j-k+1] != p[i] {
					break
				}
				memo[i][j] = memo[i-1][j-k]
				if memo[i][j] {
					break
				}
			}
		}
	}
	fmt.Println(memo)
	return memo[len(p)-1][len(s)]
}

func simplifyP(p string) ([]bool, string) {
	var canRepeat, bts = make([]bool, 0, len(p)), make([]byte, 0, len(p))
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			continue
		}
		bts = append(bts, p[i])
		if i+1 < len(p) && p[i+1] == '*' {
			canRepeat = append(canRepeat, true)
		} else {
			canRepeat = append(canRepeat, false)
		}
	}
	return canRepeat, string(bts)
}
