package UConfig

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Test GetJavaHome
func TestGetJavaHome(t *testing.T) {
	UConsole.Log(GetJavaHome(""))
	UConsole.PrintTypeAndValue(GetJavaHome(""))
	// D:\Java\jdk1.8.0_211
	// 类型(Type):string  值(Value):D:\Java\jdk1.8.0_211
}

// Test GetUserHomeDir
func TestGetUserHomeDir(t *testing.T) {
	UConsole.Log(GetUserHomeDir()) // C:\Users\Administrator
}

// Test GetJson
// ===============================================================
// path 填写格式示例
//  {
//    "name": {"first": "Tom", "last": "Anderson"},
//    "age":37,
//    "children": ["Sara","Alex","Jack"],
//    "friends": [
//      {"first": "James", "last": "Murphy"},
//      {"first": "Roger", "last": "Craig"}
//    ]
//  }
//  "name.last"          >> "Anderson"
//  "age"                >> 37
//  "children"           >> ["Sara","Alex","Jack"]
//  "children.#"         >> 3
//  "children.1"         >> "Alex"
//  "child*.2"           >> "Jack"
//  "c?ildren.0"         >> "Sara"
//  "friends.#.first"    >> ["James","Roger"]
// ===============================================================
func TestGetJson(t *testing.T) {
	v := InitUConfig(".", "demo", "json")
	UConsole.Log(GetJson(v, "GoToolcUrl", false))
	UConsole.Log(GetJson(v, "author", false))
	UConsole.Log(GetJson(v, "contents.post1", false))
	UConsole.Log(GetJson(v, "contents2.#.post1", false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetJson(v, "contents2", false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetJson(v, "contents2", true))
	UConsole.PrintAStraightLine()

	// https://github.com/ahviplc/GoJustToolc
	// LC<ahlc@sina.cn>
	// 1
	// ["3-1","3-2"]
	// --------------------------------------------------------------------------------------------------------------
	// [{"post1":"3-1","post2":"4-1"},{"post1":"3-2","post2":"4-2"}]
	// --------------------------------------------------------------------------------------------------------------
	//    [
	//	     {
	//		  "post1": "3-1",
	//		  "post2": "4-1"
	//	     },
	//	     {
	//		  "post1": "3-2",
	//		  "post2": "4-2"
	//	     }
	//    ]
	// --------------------------------------------------------------------------------------------------------------
}

// Test GetYaml
func TestGetYaml(t *testing.T) {
	v := InitUConfig(".", "demo2", "yaml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetYaml(v, "userName", false))
	UConsole.Log(GetYaml(v, "address", false))
	UConsole.Log(GetYaml(v, "company.name", false))
	UConsole.Log(GetYaml(v, "company.department", false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetYaml(v, "company.department", true))
	UConsole.PrintAStraightLine()

	// LC
	// 上海
	// LC公司
	// ["技术部"]
	// --------------------------------------------------------------------------------------------------------------
	// [
	// "技术部"
	// ]
	// --------------------------------------------------------------------------------------------------------------
}

// Test GetYaml
func TestGetToml(t *testing.T) {
	v := InitUConfig(".", "demo3", "toml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetToml(v, "server.Address", false))
	UConsole.Log(GetToml(v, "logger.Level", false))
	UConsole.Log(GetToml(v, "logger.path", false))
	UConsole.Log(GetToml(v, "viewer.DefaultFile ", false))
	UConsole.Log(GetToml(v, "database.link", false))
	UConsole.Log(GetToml(v, "swagger.user", false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetToml(v, "database.logger", false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetToml(v, "database.logger", true))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetToml(v, "database.logger.path", false))
	UConsole.Log(GetToml(v, "database.logger.path", true))
	UConsole.PrintAStraightLine()

	// :8199
	// all
	// tmp/log/GoJustToolc-demos
	// index.html
	// mysql:root:root@tcp(192.168.192.168:3306)/GoJustToolc-demos-by-lc
	// ahviplc
	//--------------------------------------------------------------------------------------------------------------
	// {"level":"all","path":"/tmp/log/GoJustToolc-demos/sql","stdout":true}
	// --------------------------------------------------------------------------------------------------------------
	// {
	//  "level": "all",
	//  "path": "/tmp/log/GoJustToolc-demos/sql",
	//  "stdout": true
	// }
	//--------------------------------------------------------------------------------------------------------------
	/// tmp/log/GoJustToolc-demos/sql
	/// tmp/log/GoJustToolc-demos/sql
	// --------------------------------------------------------------------------------------------------------------
}

// Test GetAll
// GetJson GetYaml GetToml当然都可以使用GetAll方法替代
func TestGetAll(t *testing.T) {
	v := InitUConfig(".", "demo", "json")
	UConsole.Log(GetAll(v, "GoToolcUrl", false))
	v = InitUConfig(".", "demo2", "yaml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetAll(v, "userName", false))
	v = InitUConfig(".", "demo3", "toml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetAll(v, "server.Address", false))

	// https://github.com/ahviplc/GoJustToolc
	// LC
	// :8199
}

// Test GetAllSettingsMap
func TestGetAllSettingsMap(t *testing.T) {
	v := InitUConfig(".", "demo", "json")
	allSettingsMap := GetAllSettingsMap(v)
	fmt.Println(v.AllSettings())
	// 打印
	UConsole.PrintAStraightLine()
	fmt.Println(v.AllKeys())
	UConsole.PrintAStraightLine()
	//fmt.Printf("GoToolcUrl:%s author:%s \n", v.Get("GoToolcUrl"), v.Get("author"))
	for k, v := range v.AllKeys() {
		fmt.Println(k, " -> ", v)
	}
	UConsole.PrintAStraightLine()
	UConsole.PrintAStraightLine()
	for k, v := range v.AllSettings() {
		fmt.Println(k, " -> ", v)
	}
	UConsole.PrintAStraightLine()
	UConsole.PrintAStraightLine()
	// allSettingsMap := v.AllSettings()
	UConsole.Log("循环遍历key为contents的子Map,结果如下:")
	// 使用【.(map[string]interface{})】将【allSettingsMap["contents"]】强制转成map
	// 使用类型断言 if contentsMap, ok := allSettingsMap["contents"].(map[int]interface{})
	// 如果 allSettingsMap["contents"]是(map[int]interface{})类型的话,ok就是true,contentsMap就是(map[int]interface{})类型allSettingsMap["contents"]的值.
	// 否则ok为false，contentsMap就是(map[int]interface{})类型的初始化 nil
	if contentsMap, ok := allSettingsMap["contents"].(map[string]interface{}); ok {
		UConsole.Log("类型转换成功:", ok)
		for k, v := range contentsMap {
			fmt.Println(k, " -> ", v)
		}
	} else {
		UConsole.Log("无效的类型转换:", ok) // 如果无效类型转换 则 contentsMap 会是初始值 nil
	}
	UConsole.PrintAStraightLine()
	UConsole.PrintTypeAndValue(allSettingsMap["contents2"])
	fmt.Println(reflect.TypeOf(allSettingsMap["contents2"])) // []interface {}
	c2 := allSettingsMap["contents2"]
	for k, v := range c2.([]interface{}) {
		fmt.Println(k, " -> ", v)
	}
	UConsole.PrintAStraightLine()
	if contents2Map, ok := allSettingsMap["contents2"].([]interface{}); ok {
		UConsole.Log("类型转换成功:", ok)
		for k, v := range contents2Map {
			fmt.Println(k, " -> ", v)
		}
	} else {
		UConsole.Log("无效的类型转换:", ok) // 如果无效类型转换 则 contents2Map 会是初始值 nil
	}
	UConsole.PrintAStraightLine()
	if contents2Map, ok := allSettingsMap["contents2"].([]interface{})[0].(map[string]interface{}); ok {
		UConsole.Log("类型转换成功:", ok)
		UConsole.PrintTypeAndValue(contents2Map) // 类型(Type):map[string]interface {}  值(Value):map[post1:3 post2:4]
		for k, v := range contents2Map {
			fmt.Println(k, " -> ", v)
		}
	} else {
		UConsole.Log("无效的类型转换:", ok) // 如果无效类型转换 则 contents2Map 会是初始值 nil
	}
	UConsole.PrintAStraightLine()

	// 判断key是否存在
	keyTemp := "contents3" // 不存在
	if i, ok := allSettingsMap[keyTemp]; ok {
		UConsole.Log(keyTemp, "-key-存在")
		UConsole.PrintTypeAndValue(i)
	} else {
		UConsole.Log(keyTemp, "-key-不存在")
	}
}

// Test GetAllSettingsJson
func TestGetAllSettingsJson(t *testing.T) {
	v := InitUConfig(".", "demo", "json")
	// Ugly
	UConsole.Log(GetAllSettingsJson(v, false))
	UConsole.PrintAStraightLine()
	// Pretty
	UConsole.Log(GetAllSettingsJson(v, true))
	UConsole.PrintAStraightLine()
	v = InitUConfig(".", "demo2", "yaml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetAllSettingsJson(v, false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetAllSettingsJson(v, true))
	UConsole.PrintAStraightLine()
	v = InitUConfig(".", "demo3", "toml") // 这里的configName如果也是demo 它回去读取 demo.json
	UConsole.Log(GetAllSettingsJson(v, false))
	UConsole.PrintAStraightLine()
	UConsole.Log(GetAllSettingsJson(v, true))
	UConsole.PrintAStraightLine()
}

// Test IsSupportedConfigType
func TestIsSupportedConfigType(t *testing.T) {
	// assert equality
	assert.Equal(t, true, IsSupportedConfigType("json"), "Should be true")    // 断言为二者相等
	assert.Equal(t, false, IsSupportedConfigType("json2"), "Should be false") // 断言为二者相等
	assert.True(t, IsSupportedConfigType("json"), "Should be true")
	assert.False(t, IsSupportedConfigType("json2"), "Should be false")
}
