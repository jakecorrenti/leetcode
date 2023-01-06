package medium

type MinStack struct {
    stack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack {
        stack: make([]int, 0),
        minStack:  make([]int, 0),
    }
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    if len(this.minStack) < 1 {
        this.minStack = append(this.minStack, val)
    } else if val <= this.minStack[len(this.minStack) - 1] {
        this.minStack = append(this.minStack, val)
    }
}


func (this *MinStack) Pop()  {
    if this.stack[len(this.stack) - 1] == this.minStack[len(this.minStack) - 1] {
        this.minStack = this.minStack[:len(this.minStack) - 1]
    }
    this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
    if len(this.stack) < 1 {
        return 0
    }
    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
    if len(this.minStack) < 1 {
        return 0
    }
    return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
