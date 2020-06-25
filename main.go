package main

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/UUtils"
	"github.com/tidwall/gjson"
)

// some test main
func main() {
	// ToUpper ToLower
	fmt.Println(UUtils.ToUpper("ah vip lc"))
	fmt.Println(UUtils.ToLower("SH vip LC"))
	UConsole.Log(UUtils.ToUpper("ah vip lc"))
	UConsole.Log(UUtils.ToLower("SH vip LC"))

	// UConsole
	UConsole.Log("我就是一个长长的字符串")
	UConsole.Log("string1", "string2")
	UConsole.Log(110, 120)
	UConsole.Log("我是字符串", 2)
	UConsole.Log(1.1, 2.2, 3.3)
	UConsole.Log(1.11, 2.22, 3.33)
	UConsole.Log(1.11, 2.222, 3.3333)

	// gjson
	const json = `{"name":{"first":"ahviplc","last":"shviplc"},"age":18}`
	value := gjson.Get(json, "name.first")
	UConsole.Log(value.String())
	UConsole.Log(gjson.Get(json, "name.last").String())
	UConsole.Log(gjson.Get(json, "age").String())
}
