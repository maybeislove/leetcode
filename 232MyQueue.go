package main

func main() {
	q := Constructor1()
	q.Push(1)
	q.Push(3)
	q.Push(5)
	q.Pop()
	q.Pop()
	q.Pop()

}

type MyQueue struct {
	in, out []int
}

func Constructor1() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.in2out()
	}
	x := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.out) != 0 {
		return this.out[len(this.out)-1]
	}
	return this.in[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

func (q *MyQueue) in2out() {
	for len(q.in) > 0 {
		q.out = append(q.out, q.in[len(q.in)-1])
		q.in = q.in[:len(q.in)-1]
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
