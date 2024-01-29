package kata

type MyQueue []int

func Constructor() MyQueue {
	new := MyQueue{}
	return new
}

func (this *MyQueue) Push(x int) {
	*this = append(*this, x)
}

func (this *MyQueue) Pop() int {
	a := []int{}
	for _, m := range *this {
		a = append(a, m)
	}
	*this = a[1:]
	return a[0]
}

func (this *MyQueue) Peek() int {
	a := []int{}
	for _, m := range *this {
		a = append(a, m)
	}
	return a[0]
}

func (this *MyQueue) Empty() bool {
	return len(*this) == 0
}
