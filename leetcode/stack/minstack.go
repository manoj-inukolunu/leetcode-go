package main

import (
	"fmt"
)

type MinStack struct {
	arr []pair
}

type pair struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]pair, 0)}
}

func (this *MinStack) Push(val int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, pair{val, val})
		return
	}
	top := this.peek()
	if val < top.min {
		this.arr = append(this.arr, pair{val, val})
	} else {
		this.arr = append(this.arr, pair{val, top.min})
	}
}
func (this *MinStack) peek() pair {
	return this.arr[len(this.arr)-1]
}
func (this *MinStack) Pop() {
	this.arr = this.arr[0 : len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1].val
}

func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1].min
}

func main() {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-1)
	fmt.Println(minstack.arr)
	// minstack.Pop()
	// fmt.Println(minstack.arr)

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
