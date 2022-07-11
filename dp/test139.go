package main

import "math"

func WordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}
	//mem[i]表示[0,i]区间内是否能被dict组合出来
	//状态转移：memo[i+1] = (memo[i+1 - len(wordDict[j])] && wordDict[j] == s[i+1 - len(wordDict[j]):i+1]), j:[]
	var memo = make([]bool, len(s))
	var matchDict []string
	var maxWordLen = 0
	for i := 0; i < len(wordDict); i++ {
		if len(wordDict[i]) <= len(s) {
			if s[:len(wordDict[i])] == wordDict[i] {
				memo[len(wordDict[i])-1] = true
			}
			matchDict = append(matchDict, wordDict[i])
			maxWordLen = int(math.Max(float64(maxWordLen), float64(len(wordDict[i]))))
		}

	}
	wordDict = matchDict

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) &&
				memo[i-len(wordDict[j])] &&
				s[i-len(wordDict[j])+1:i+1] == wordDict[j] {
				memo[i] = true
				break
			}
		}
	}

	return memo[len(s)-1]
}
