package array

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	var isRepeat, repeatCnt, firstSegment = cutFirstSegment(s)
	if isRepeat {
		var repeatRes = decodeString(firstSegment)
		return buildStr(repeatCnt, repeatRes) + decodeString(s[3+len(firstSegment):])
	} else {
		return firstSegment + decodeString(s[len(firstSegment):])
	}

}

// 返回值：是否是重复数，重复次数，重复字符串
func cutFirstSegment(s string) (bool, int, string) {
	if !isNumber(s[0]) {
		var endIndex = 1
		for endIndex < len(s) && !isNumber(s[endIndex]) {
			endIndex++
		}
		return false, 0, s[:endIndex]
	}
	var endIndex = 1
	var stack []byte
	stack = append(stack, s[endIndex])
	endIndex++
	for len(stack) > 0 {
		if s[endIndex] == '[' {
			stack = append(stack, '[')
		}
		if s[endIndex] == ']' {
			stack = stack[:len(stack)-1]
		}
		endIndex++
	}
	return true, int(s[0] - '0'), s[2 : endIndex-1]
}

func buildStr(repeat int, s string) string {
	var res = ""
	for i := 0; i < repeat; i++ {
		res += s
	}
	return res
}

func isNumber(s byte) bool {
	return s >= '1' && s <= '9'
}
