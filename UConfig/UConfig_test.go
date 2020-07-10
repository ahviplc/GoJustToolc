package UConfig

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test GetJavaHome
func TestGetJavaHome(t *testing.T) {
	UConsole.Log(GetJavaHome(""))
	UConsole.PrintTypeAndValue(GetJavaHome(""))
	// D:\Java\jdk1.8.0_211
	// 类型(Type):string  值(Value):D:\Java\jdk1.8.0_211
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
