package main

import "fmt"

func main() {

}

func restoreIpAddresses(s string) []string {
	dfs(s, 0, []int{})
	return res
}

var res []string

func dfs(s string, level int, path []int) {
	if len(path) == 3 {
		fmt.Println(path)
		if isLegalAddrPart(s[path[2]:]) {
			var tmpAddr = s[0:path[0]] + "." + s[path[0]:path[1]] + "." + s[path[1]:path[2]] + "." + s[path[2]:]
			res = append(res, tmpAddr)
		}
		return
	}
	var begin = 0
	if len(path) > 0 {
		begin = path[len(path)-1]
	}
	for cur := begin + 1; cur < len(s); cur++ {
		var tmpStr = s[begin:cur]
		if !isLegalAddrPart(tmpStr) {
			break
		}
		dfs(s, level+1, append(path, cur))
		path = path[0 : len(path)-1]
	}
}

// 每个整数位于 0 到 255 之间组成，且不能含有前导 0
func isLegalAddrPart(s string) bool {
	if len(s) == 0 {
		return false
	}
	var val = 0
	for i := len(s) - 1; i >= 0; i-- {
		val = val*10 + int(s[i]-'0')
		if val > 255 {
			return false
		}
	}
	//
	if val == 0 {
		return s == "0"
	}
	return s[0] == '0'
}
