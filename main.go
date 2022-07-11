package main

import (
	"fmt"
	"learn_algorithm/array"
)

func main() {
	//"applepenapple", wordDict = ["apple", "pen"]
	//s = "leetcode", wordDict = ["leet", "code"]
	//s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
	//"abcd" ["a","abc","b","cd"]
	//fmt.Println(dp.WordBreak("abcd", []string{"a", "abc", "b", "cd"}))

	//nums = [1,1,1,1,1], target = 3
	//fmt.Println(dp.FindTargetSumWays([]int{1, 2, 1}, 0))

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//var lru = _map.Constructor(2)
	//lru.Put(1, 0)
	//lru.Put(2, 2)
	//lru.Get(1)
	//
	//lru.Put(3, 3)
	//lru.Get(2)
	//lru.Put(4, 4)
	//lru.Get(1)
	//lru.Get(3)
	//lru.Get(4)

	//fmt.Println(array.MaxCoins([]int{3, 1, 5, 8}))

	//fmt.Println(array.TopKFrequent([]int{1, 1, 1, 2, 2, 2, 2, 3, 4}, 3))

	//fmt.Println(array.decodeString("abc3[cd]xyz"))

	fmt.Println(array.SubarraySum([]int{1, 1, 1}, 2))

}
