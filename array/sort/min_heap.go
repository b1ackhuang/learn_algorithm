package main

type MinHeap struct {
	data []int
}

func (m *MinHeap) Init(nums []int) {
	m.data = nums //应该要深拷贝
	for i := len(m.data)/2 - 1; i >= 0; i-- {
		m.shiftDown(i)
	}
}

func (m *MinHeap) Insert(val int) {
	m.data = append(m.data, val)
	m.shiftUp()
}

func (m *MinHeap) shiftUp() {
	for i := len(m.data) - 1; i > 0 && m.data[i/2] > m.data[i]; i /= 2 {
		m.swap(i/2, i)
	}
}

func (m *MinHeap) shiftDown(beginIndex int) {
	for i := beginIndex; i*2+1 < len(m.data); {
		var minSonIndex = 2*i + 1
		if 2*i+2 < len(m.data) && m.data[2*i+2] < m.data[2*i+1] {
			minSonIndex = 2*i + 2
		}
		if m.data[minSonIndex] > m.data[i] {
			break
		}
		m.swap(minSonIndex, i)
		i = minSonIndex
	}
}

func (m *MinHeap) Top() int {
	var ret = m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.shiftDown(0)
	return ret
}

func (m *MinHeap) swap(x, y int) {
	var tmp = m.data[x]
	m.data[x] = m.data[y]
	m.data[y] = tmp
}
