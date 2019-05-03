package stack

// 232. 用栈实现队列
// https://leetcode-cn.com/problems/implement-queue-using-stacks/submissions/

type MyQueue struct {
	inStack, outStack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.checkOutStack()
	return this.outStack.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.checkOutStack()
	return this.outStack.Top().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.inStack.IsEmpty() && this.outStack.IsEmpty()
}

func (this *MyQueue) checkOutStack() {
	if this.outStack.IsEmpty() {
		for !this.inStack.IsEmpty() {
			this.outStack.Push(this.inStack.Pop())
		}
	}
}
