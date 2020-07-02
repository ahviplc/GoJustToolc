package UQueue

// 队列（Queue）
type Queue []interface{}

// Push
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pop
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
