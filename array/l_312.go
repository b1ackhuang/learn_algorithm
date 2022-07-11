package array

//
//const (
//	Used    = '1'
//	NotUsed = '0'
//)
//
//var mapMemo map[string]int
//
//func MaxCoins(nums []int) int {
//	mapMemo = make(map[string]int)
//	var used = make([]byte, len(nums))
//	for i := 0; i < len(nums); i++ {
//		used[i] = NotUsed
//	}
//	return dfs(nums, 0, used)
//}
//
//func dfs(nums []int, level int, used []byte) int {
//	if level == len(nums) { // 所有枪打完
//		return 0
//	}
//
//	if tmpRes, ok := mapMemo[string(used)]; ok {
//		return tmpRes
//	}
//
//	var lIndex = make([]int, len(nums))
//	lIndex[0] = -1
//	for i := 1; i < len(nums); i++ {
//		if used[i-1] == NotUsed {
//			lIndex[i] = i - 1
//		} else {
//			lIndex[i] = lIndex[i-1]
//		}
//	}
//	var rIndex = make([]int, len(nums))
//	rIndex[len(nums)-1] = -1
//	for i := len(nums) - 2; i >= 0; i-- {
//		if used[i+1] == NotUsed {
//			rIndex[i] = i + 1
//		} else {
//			rIndex[i] = rIndex[i+1]
//		}
//	}
//
//	//第level次尝试开枪
//	var res = 0
//	for i := 0; i < len(nums); i++ {
//		if used[i] == Used {
//			continue
//		}
//		used[i] = Used
//		var curSum = nums[i]
//		if lIndex[i] != -1 {
//			curSum *= nums[lIndex[i]]
//		}
//		if rIndex[i] != -1 {
//			curSum *= nums[rIndex[i]]
//		}
//		curSum += dfs(nums, level+1, used)
//		used[i] = NotUsed
//		if curSum > res {
//			res = curSum
//		}
//	}
//	mapMemo[string(used)] = res
//	return res
//}
//
//func countBits(n int) []int {
//	var bitMask = 1
//	var bitCnt = make([]int, 32)
//	for i := 0; i <= n; i++ {
//		for j := 0; j < 32; j++ {
//			if i&(bitMask<<j) > 0 {
//				bitCnt[j] += 1
//			}
//		}
//	}
//	return bitCnt
//}
//
////func MaxCoins(nums []int) int {
////	var memo = make([][]int, len(nums))
////	for i := 0; i < len(nums); i++ {
////		memo[i] = make([]int, len(nums))
////	}
////
////	// 第i次戳第j个气球有最大值时, 戳气球路径
////	var prePathUsed = make([][]map[int]bool, len(nums))
////	for i := 0; i < len(nums); i++ {
////		prePathUsed[i] = make([]map[int]bool, len(nums))
////	}
////
////	for i := 0; i < len(nums); i++ {
////		prePathUsed[0][i] = make(map[int]bool)
////		prePathUsed[0][i][i] = true
////		memo[0][i] = nums[i]
////		if i > 0 {
////			memo[0][i] *= nums[i-1]
////		}
////		if i < len(nums)-1 {
////			memo[0][i] *= nums[i+1]
////		}
////	}
////
////	var res = 0
////	for i := 1; i < len(nums); i++ {
////		var maxCurSum = 0
////		var maxKIndex = -1
////		for j := 0; j < len(nums); j++ { //尝试第i次戳第j个气球
////			for k := 0; k < len(nums); k++ { //查看上一次戳气球的最大值
////				if k == j { //两次不能戳同一个地方
////					continue
////				}
////				var tmpSum = nums[j]
////				var tmpPathUsed = prePathUsed[i-1][k]
////				//左边界
////				var l = j - 1
////				for l >= 0 && tmpPathUsed[l] {
////					l--
////				}
////				if l >= 0 {
////					tmpSum *= nums[l]
////				}
////				//右边界
////				var r = j + 1
////				for r < len(nums) && tmpPathUsed[r] {
////					r++
////				}
////				if r < len(nums) {
////					tmpSum *= nums[r]
////				}
////				tmpSum += memo[i-1][k]
////				if maxCurSum <= tmpSum {
////					maxCurSum = tmpSum
////					maxKIndex = k
////				}
////			}
////			prePathUsed[i][j] = make(map[int]bool)
////			prePathUsed[i][j][j] = true
////			for tmpPreIndex := range prePathUsed[i-1][maxKIndex] {
////				prePathUsed[i][j][tmpPreIndex] = true
////			}
////		}
////		if res < maxCurSum {
////			res = maxCurSum
////		}
////	}
////	return res
////
////}
