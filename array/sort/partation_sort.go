package main

func partitionSort(nums []int) {
	//var a = byte(5
	//ints := nums[5:]

}

func containWord(ms, mt map[byte]int) bool {
	for k, tCnt := range mt {
		if sCnt, ok := ms[k]; ok {
			if sCnt < tCnt {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func isHappy(n int) bool {
	var mapHistory = make(map[int]int)
	for _, ok := mapHistory[n]; !ok; _, ok = mapHistory[n] {
		mapHistory[1] = 1
	}
	return false
}
