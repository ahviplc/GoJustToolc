package URetriever

// 接口变量里面有
// 1.接口变量自带指针
// 2.接口变量同样采用值传递,几乎不需要使用接口的指针
// 3.指针的接收者只能以指针的方式使用;值接收者都可以

// 查看接口变量
// 1.表示任何类型: interface{]
// 2.Type Assertion
// 3.Type Switch

// 定义Retriever接口
type Retriever interface {
	Get(url string) string
}

// 定义Poster接口
type Poster interface {
	Post(url string, form map[string]string) string
}

// Download方法
func Download(r Retriever, url string) string {
	return r.Get(url)
}

// Post方法
func Post(poster Poster, url string) string {
	poster.Post(url,
		map[string]string{
			"name":   "LC",
			"slogan": "Just do it.",
		})
	return "Post ok"
}

// 定义RetrieverPoster接口
type RetrieverPoster interface {
	Retriever
	Poster
}

// Session方法 RetrieverPoster接口 同时可以使用Get和Post
func Session(s RetrieverPoster, url string) string {
	s.Post(url, map[string]string{
		"contents": "another faked contents. Url: " + url,
	})
	return s.Get(url)
}
