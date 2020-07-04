package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// 通过get请求此url获取其页面body
// 可以通过指针来访问Get 假设Retriever结构体很大的话
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
