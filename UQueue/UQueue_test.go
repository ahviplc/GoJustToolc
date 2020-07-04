package UQueue

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test UQueue
func TestUQueue(t *testing.T) {
	UConsole.PrintAStraightLine()
	q := Queue{1}

	q.Push(2)
	q.Push(3)

	UConsole.Log(q.Pop())
	UConsole.Log(q.Pop())
	UConsole.Log(q.IsEmpty())
	UConsole.Log(q.Pop())
	UConsole.Log(q.IsEmpty())

	q.Push("abc")
	UConsole.Log(q.Pop())
	UConsole.PrintAStraightLine()
}
