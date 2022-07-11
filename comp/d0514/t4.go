package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(x, y int) bool {
		if people[x][0] < people[y][0] {
			return true
		}
		if people[x][0] > people[y][0] {
			return false
		}
		return people[x][1] < people[y][1]
	})
	var res [][]int = make([][]int, 2)
	for len(people) > 0 {
		var target = people[0]
		people = people[1:]
		var insertIndex = target[1] + 1
		var tmp = make([][]int, len(res)-1)
		tmp = append(tmp, res[:insertIndex]...)
		tmp = append(tmp, target)
		tmp = append(tmp, res[insertIndex:]...)
		res = tmp
	}
	return res
}
