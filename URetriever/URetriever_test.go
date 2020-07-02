package URetriever

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/URetriever/mock"
	"github.com/ahviplc/GoJustToolc/URetriever/real"
	"testing"
	"time"
)

// Test URetriever.Download()
func TestDownload(t *testing.T) {
	var r Retriever
	r = &mock.Retriever{Contents: "https://time.is/Beijing"}
	UConsole.PrintAStraightLine()
	fmt.Printf("%T | %v", r, r)
	UConsole.Log(Download(r, "https://time.is"))
	UConsole.PrintAStraightLine()
}

// Test URetriever.Download()
func TestDownloadV2(t *testing.T) {
	const UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Safari/605.1.15"
	var r Retriever
	// 不使用指针获取Get的话 直接使用下行代码接口.
	// r = real.Retriever{UA, time.Minute}
	// 如果使用指针 还是使用上行代码 则会报错误:Retriever does not implement Retriever (Get method has pointer receiver)
	// 正确使用应该是取地址 加个 &
	r = &real.Retriever{UA, time.Minute}
	UConsole.PrintAStraightLine()
	fmt.Printf("%T %v", r, r)
	UConsole.Log(Download(r, "https://time.is"))
	UConsole.PrintAStraightLine()
}

// Test 获取Retriever类型
func TestGetType(t *testing.T) {
	// 使用 switch
	UConsole.PrintAStraightLine()
	const UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Safari/605.1.15"
	var r Retriever
	r = &real.Retriever{UA, time.Minute} // 备注: 直接使用 := 下面的类型断言 会报错: invalid type assertion: r.(*real.Retriever) (non-interface type *real.Retriever on left)
	Inspext(r)
	UConsole.PrintAStraightLine()
	var r2 Retriever
	r2 = &mock.Retriever{Contents: "https://time.is/Beijing"}
	Inspext(r2)
	UConsole.PrintAStraightLine()

	// 使用 类型断言 Type Assertion 来判断
	realRetriever := r.(*real.Retriever) // 如果指针 * 不加 会报错: real.Retriever does not implement Retriever (Get method has pointer receiver)
	fmt.Println(realRetriever.TimeOut)   // 1m0s
	UConsole.PrintAStraightLine()

	// 如果 real 断言了 mock 则会报错: panic: interface conversion: URetriever.Retriever is *real.Retriever, not mock.Retriever
	// 解决方法: 加 ok 返回
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		UConsole.Log("not a mock retriever")
	}
	UConsole.PrintAStraightLine()
}

// Test URetriever.Session()
func TestSession(t *testing.T) {
	UConsole.PrintAStraightLine()
	retriever := mock.Retriever{Contents: "https://time.is/Tokyo"}
	UConsole.Log(Session(&retriever, "https://time.is/Shanghai"))
	UConsole.PrintAStraightLine()

	// 隐藏掉 Session 方法下 s.Post()方法 输出
	// -----------------------------------------------------------------------------------------------------
	// I am Mock Retriever Get, a fake url response!~LC Your Get url: https://time.is/Shanghai
	// 返回描述语:我这个Retriever Get是直接返回Contents的呀 | Contents：https://time.is/Tokyo
	// -----------------------------------------------------------------------------------------------------

	// 放开 Session 方法下 s.Post()方法 输出 其Contents可以看到被改变了
	// -----------------------------------------------------------------------------------------------------
	// map form 中 key slogan 不存在...
	// I am Mock Retriever Get, a fake url response!~LC Your Get url: https://time.is/Shanghai
	// 返回描述语:我这个Retriever Get是直接返回Contents的呀 | Contents：I am Mock Retriever Post 返回的Contents: another faked contents. Url: https://time.is/Shanghai
	// -----------------------------------------------------------------------------------------------------
}

// Test URetriever.Post()
func TestPost(t *testing.T) {
	UConsole.PrintAStraightLine()
	retriever := mock.Retriever{Contents: "https://time.is/Anhui_Sheng"}
	UConsole.Log(Post(&retriever, "https://time.is/Anhui_Sheng"))
	UConsole.PrintAStraightLine()

	// -----------------------------------------------------------------------------------------------------
	// 存在slogan为: Just do it.
	// Post ok
	// -----------------------------------------------------------------------------------------------------
}
