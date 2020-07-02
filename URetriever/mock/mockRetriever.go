package mock

import "github.com/ahviplc/GoJustToolc/UConsole"

type Retriever struct {
	Contents string
}

// 统一改成指针接收者
func (r *Retriever) Post(url string, form map[string]string) string {
	returnDesc := "I am Mock Retriever Post 返回的Contents: "
	r.Contents = returnDesc + form["contents"]
	// 判断map中是否存在 slogan 这个key
	if _, ok := form["slogan"]; ok {
		// ok为 true 存在 打印值
		UConsole.Log("存在slogan为: " + form["slogan"])
	} else {
		// ok为 false 不存在 提示不存在
		UConsole.Log("map form 中 key slogan 不存在...")
	}

	return "ok"
}

func (r *Retriever) Get(url string) string {
	UConsole.Log("I am Mock Retriever Get, a fake url response!~LC Your Get url: " + url)
	// 返回描述语

	//var returnDesc string
	//returnDesc = "###"

	returnDesc := "返回描述语:我这个Retriever Get是直接返回Contents的呀 | Contents："
	// 直接返回Contents
	return returnDesc + r.Contents
}
