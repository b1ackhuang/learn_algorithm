package array

func TopKFrequent(nums []int, k int) []int {
	//排序map
	var mapFreq = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mapFreq[nums[i]] += 1
	}
	var freq = make([]int, 0, len(mapFreq))
	var freqVal = make([]int, 0, len(mapFreq))
	for tmpFreqVal, tmpFreq := range mapFreq {
		freqVal = append(freqVal, tmpFreqVal)
		freq = append(freq, tmpFreq)
	}

	var minHeap = MinHeap{cap: k}
	for i := 0; i < len(freq); i++ {
		minHeap.Add(freq[i], freqVal[i])
	}
	var res = make([]int, k)
	for k > 0 {
		res[k-1] = minHeap.PopTop()
		k -= 1
	}
	return res
}

type MinHeap struct {
	cap      int
	dataFreq []int
	data     []int
}

func (m *MinHeap) Add(freq, val int) {
	//空堆
	if len(m.dataFreq) == 0 {
		m.dataFreq = append(m.dataFreq, freq)
		m.data = append(m.data, val)
		return
	}
	//堆未满
	if len(m.dataFreq) < m.cap {
		m.addIgnoreCap(freq, val)
		return
	}
	//堆已满
	if freq <= m.dataFreq[0] {
		return
	}
	m.addIgnoreCap(freq, val)
	m.PopTop()
}

func (m *MinHeap) PopTop() int {
	var res = m.data[0]
	m.swap(0, len(m.dataFreq)-1)
	m.data = m.data[:len(m.dataFreq)-1]
	m.dataFreq = m.dataFreq[:len(m.dataFreq)-1]
	m.shiftDown(0)
	return res
}

func (m *MinHeap) addIgnoreCap(freq, val int) {
	m.dataFreq = append(m.dataFreq, freq)
	m.data = append(m.data, val)
	m.shiftUp(len(m.dataFreq) - 1)
	return
}

func (m *MinHeap) shiftUp(begin int) {
	for begin > 0 {
		var parentIndex = (begin+1)/2 - 1
		if m.dataFreq[begin] < m.dataFreq[parentIndex] {
			m.swap(begin, parentIndex)
			begin = parentIndex
		} else {
			break
		}
	}
}

func (m *MinHeap) shiftDown(begin int) {
	for begin < len(m.dataFreq) {
		var lSon, rSon = 2*begin + 1, 2*begin + 2
		if lSon >= len(m.dataFreq) {
			break
		}
		var minSonIndex, minSonVal = lSon, m.dataFreq[lSon]
		if rSon < len(m.dataFreq) && m.dataFreq[rSon] < minSonVal {
			minSonIndex = rSon
			minSonVal = m.dataFreq[rSon]
		}
		if m.dataFreq[begin] <= minSonVal {
			break
		}
		m.swap(begin, minSonIndex)
		begin = minSonIndex
	}
}

func (m *MinHeap) swap(x, y int) {
	m.data[x], m.data[y] = m.data[y], m.data[x]
	m.dataFreq[x], m.dataFreq[y] = m.dataFreq[y], m.dataFreq[x]
}
