package UConfig

import (
	"encoding/json"
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/UUtils/UStringUtil"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"os"
	"runtime"
)

// UConfig.go说明
// 配置文件读取解析工具

// 所有支持的配置文件类型
// SupportedExts are universally supported extensions.
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"}

// 获取JAVA_HOME环境变量路径
// envKey 默认值 "JAVA_HOME"
// 如果环境变量key错误或者无JAVA环境变量 则输出字符串string初始值(空字符串) ""
func GetJavaHome(envKey string) string {
	if UStringUtil.IsEmpty(envKey) {
		// 如果key为空 则赋默认值 "JAVA_HOME"
		envKey = "JAVA_HOME"
	}
	getenv := os.Getenv(envKey)
	return getenv
}

// 获取 GetUserHomeDir
func GetUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// 初始化配置文件读取解析三方库viper
// InitUConfig()
// configPath 配置文件路径 如果输入为"" 默认为当前路径 .
// configName 配置文件名
// configType 配置文件类型 支持【JSON, TOML, YAML, HCL, envfile and Java properties config files】 这里的工具类只封装了前三个
func InitUConfig(configPath string, configName string, configType string) *viper.Viper {
	v := viper.New()
	// 设置读取的配置文件路径
	if UStringUtil.IsEmpty(configPath) {
		configPath = "."
	}
	v.AddConfigPath(configPath)
	// 设置读取的配置文件名
	v.SetConfigName(configName)
	// 设置读取的配置文件类型
	v.SetConfigType(configType)

	// 读取
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		UConsole.DebugPrintError(err)
	}
	return v
}

// Json类型配置文件读取解析
// v Viper结构指针
// indentFlag 是否缩进 true 缩进 false 不缩进
// path
// ===============================================================
// path 详细说明
// Get searches json for the specified path.
// A path is in dot syntax, such as "name.last" or "age".
// When the value is found it's returned immediately.
//
// A path is a series of keys searated by a dot.
// A key may contain special wildcard characters '*' and '?'.
// To access an array value use the index as the key.
// To get the number of elements in an array or to access a child path, use
// the '#' character.
// The dot and wildcard character can be escaped with '\'.
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
//
// This function expects that the json is well-formed, and does not validate.
// Invalid json will not panic, but it may return back unexpected results.
// If you are consuming JSON from an unpredictable source then you may want to
// use the Valid function first.
// ===============================================================
func GetJson(v *viper.Viper, path string, indentFlag bool) string {
	allSettingsMap := v.AllSettings()
	var jsonTemp []byte
	var errTemp error
	if indentFlag {
		// 相比 Marshal 方法 MarshalIndent 方法对Json多了一些格式处理
		jsonTemp, errTemp = json.MarshalIndent(allSettingsMap, "", "\t") //格式化编码
	} else {
		jsonTemp, errTemp = json.Marshal(allSettingsMap) //格式化编码
	}

	if errTemp != nil {
		fmt.Printf("err:%s\n", errTemp)
		UConsole.DebugPrintError(errTemp)
	}
	return gjson.Get(string(jsonTemp), UStringUtil.ToLower(path)).String()
}

// Yaml类型配置文件读取解析
// v Viper结构指针
// indentFlag 是否缩进 true 缩进 false 不缩进
// path 详细说明同GetJson()一样
func GetYaml(v *viper.Viper, path string, indentFlag bool) string {
	allSettingsMap := v.AllSettings()
	var jsonTemp []byte
	var errTemp error
	if indentFlag {
		// 相比 Marshal 方法 MarshalIndent 方法对Json多了一些格式处理
		jsonTemp, errTemp = json.MarshalIndent(allSettingsMap, "", "\t") //格式化编码
	} else {
		jsonTemp, errTemp = json.Marshal(allSettingsMap) //格式化编码
	}

	if errTemp != nil {
		fmt.Printf("err:%s\n", errTemp)
		UConsole.DebugPrintError(errTemp)
	}
	return gjson.Get(string(jsonTemp), UStringUtil.ToLower(path)).String()
}

// Toml类型配置文件读取解析
// v Viper结构指针
// indentFlag 是否缩进 true 缩进 false 不缩进
// path 详细说明同GetJson()一样
func GetToml(v *viper.Viper, path string, indentFlag bool) string {
	allSettingsMap := v.AllSettings()
	var jsonTemp []byte
	var errTemp error
	if indentFlag {
		// 相比 Marshal 方法 MarshalIndent 方法对Json多了一些格式处理
		jsonTemp, errTemp = json.MarshalIndent(allSettingsMap, "", "\t") //格式化编码
	} else {
		jsonTemp, errTemp = json.Marshal(allSettingsMap) //格式化编码
	}

	if errTemp != nil {
		fmt.Printf("err:%s\n", errTemp)
		UConsole.DebugPrintError(errTemp)
	}
	return gjson.Get(string(jsonTemp), UStringUtil.ToLower(path)).String()
}

// 全部类型配置文件读取解析 支持【JSON, TOML, YAML, HCL, envfile and Java properties config files】
// 【"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"】
// 目前只有JSON, TOML, YAML这三个有具体封装方法【GetJson GetYaml GetToml】 其他的调用GetAll()即可 当然JSON, TOML, YAML这三个也可调用
// v Viper结构指针
// indentFlag 是否缩进 true 缩进 false 不缩进
// path 详细说明同GetJson()一样
func GetAll(v *viper.Viper, path string, indentFlag bool) string {
	allSettingsMap := v.AllSettings()
	var jsonTemp []byte
	var errTemp error
	if indentFlag {
		// 相比 Marshal 方法 MarshalIndent 方法对Json多了一些格式处理
		jsonTemp, errTemp = json.MarshalIndent(allSettingsMap, "", "\t") //格式化编码
	} else {
		jsonTemp, errTemp = json.Marshal(allSettingsMap)
	}

	if errTemp != nil {
		fmt.Printf("err:%s\n", errTemp)
		UConsole.DebugPrintError(errTemp)
	}
	return gjson.Get(string(jsonTemp), UStringUtil.ToLower(path)).String()
}

// 判断是否支持某类型的配置文件读取解析 true 是支持 false 不支持
// configType 要判断是否倍支持的配置文件类型 支持的配置文件类型:【"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"】
func IsSupportedConfigType(configType string) bool {
	if UStringUtil.StringInSlice(configType, SupportedExts) {
		return true
	}
	return false
}
