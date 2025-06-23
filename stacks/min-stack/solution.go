package stacks

type MinStack struct {
	stack []int
	min   []int
	top   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (ms *MinStack) Push(val int) {
	if ms.top == -1 {
		ms.min = append(ms.min, val)
	} else {
		ms.min = append(ms.min, min(val, (ms.min)[ms.top]))
	}
	ms.stack = append(ms.stack, val)
	ms.top++
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:ms.top]
	ms.min = ms.min[:ms.top]
	ms.top--
}

func (ms *MinStack) Top() int {
	return (ms.stack)[ms.top]
}

func (ms *MinStack) GetMin() int {
	return (ms.min)[ms.top]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
