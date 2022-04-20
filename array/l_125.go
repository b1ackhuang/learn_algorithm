package array

import (
	"math"
	"strings"
)

const Distance = 'a' - 'A'

// "A man, a plan, a canal: Panama"
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var l, r = 0, len(s) - 1
	for l < r {
		if !isLetter(s[l]) && !isDigest(s[l]) {
			l++
			continue
		}
		if !isLetter(s[r]) && !isDigest(s[r]) {
			r++
			continue
		}
		if !isSimilar(s[l], s[r]) {
			return false
		}
	}
	return true
}

func isSimilar(a, b byte) bool {
	if isDigest(a) && isDigest(b) {
		return a == b
	}
	if isDigest(a) || isDigest(b) {
		return false
	}
	var distance = math.Abs(float64(a - b))
	return distance == 0 || distance == float64(Distance)
}

func isLetter(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isDigest(s byte) bool {
	return s >= '0' && s <= '9'
}
