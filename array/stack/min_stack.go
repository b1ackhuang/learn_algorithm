package main

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minData) == 0 ||
		val <= this.minData[len(this.minData)-1] {
		this.minData = append(this.minData, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	var delVal = this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if delVal == this.minData[len(this.minData)-1] {
		this.minData = this.minData[:len(this.minData)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minData) > 0 {
		return this.minData[len(this.minData)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
