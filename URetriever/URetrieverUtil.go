package URetriever

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/URetriever/mock"
	"github.com/ahviplc/GoJustToolc/URetriever/real"
)

// 获取Retriever类型 使用 Type Switch 来判断
func Inspext(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		UConsole.Log("我是*mock.Retriever类型")
		UConsole.Log("Contents:", v.Contents)
	case *real.Retriever:
		UConsole.Log("我是*real.Retriever类型")
		UConsole.Log("UA:", v.UserAgent)
	}
}
