package USafeMap

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"sync"
)

// 线程安全的map接口
type USafeMap interface {
	// Get接口
	Get(k interface{}) interface{}
	// Set接口
	Set(k interface{}, v interface{})
}

// 线程安全的map接口的具体实现结构
type USafeMapImpl struct {
	USafeMap
	sync.RWMutex
	data map[interface{}]interface{}
}

// New()函数
func New() USafeMap {
	return &USafeMapImpl{
		data: make(map[interface{}]interface{}),
	}
}

// Get的具体实现
func (this *USafeMapImpl) Get(k interface{}) interface{} {
	this.Lock()
	tmp := this.data[k]
	this.Unlock()
	return tmp
}

// Set的具体实现
func (this *USafeMapImpl) Set(k interface{}, v interface{}) {
	this.Lock()
	this.data[k] = v
	this.Unlock()
}

// String() 查看此map的data数据
func (this *USafeMapImpl) String() string {
	UConsole.PrintTypeAndValue(this.data)
	return fmt.Sprint(this.data)
}
