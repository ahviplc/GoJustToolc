package USafeMap

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"sync"
	"testing"
	"time"
)

// Test USafeMap
func TestUSafeMap(t *testing.T) {
	// 新建一个USafeMap结构对象
	m := New()
	// 打印一下类型与值
	UConsole.PrintAStraightLine()
	UConsole.PrintTypeAndValue(m)
	UConsole.PrintAStraightLine()

	var wg sync.WaitGroup

	// 写协程
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(a int) {
			m.Set(a, fmt.Sprintf("我是#%d", a))
			time.Sleep(time.Millisecond * 100)
			wg.Done()
		}(i)
	}

	// 读协程
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(a int) {
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("%d:%s\n", a, m.Get(a))
			wg.Done()
		}(i)
	}

	wg.Wait()

	// 打印一下类型与值
	UConsole.PrintAStraightLine()
	UConsole.PrintTypeAndValue(m)
	UConsole.PrintAStraightLine()

	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):map[interface {}]interface {}  值(Value):map[]
	// 类型(Type):*USafeMap.USafeMapImpl  值(Value):map[]
	// -----------------------------------------------------------------------------------------------------
	// 4:我是#4
	// 5:我是#5
	// 0:我是#0
	// 2:我是#2
	// 1:我是#1
	// 3:我是#3
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):map[interface {}]interface {}  值(Value):map[0:我是#0 1     :我是#1 2:我是#2 3:我是#3 4:我是#4 5:我是#5]
	// 类型(Type):*USafeMap.USafeMapImpl  值(Value):map[0:我是#0 1:我是#1        2:我是#2 3:我是#3 4:我是#4 5:我是#5]
	// -----------------------------------------------------------------------------------------------------
}
