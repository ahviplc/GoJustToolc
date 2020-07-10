package UStringUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 运行错误 import cycle not allowed in test (这和错误是因为 UConsole 中要导入使用了 UStringUtil )
// 需要解决那就是在本测试中不使用使用 UConsole 中的任何方法 使用断言assert进行测试

// Test UStringUtil.IsBlank()
func TestIsBlank(t *testing.T) {
	blank := IsBlank("")
	assert.True(t, blank, "Should be true")
	blank2 := IsBlank(" ")
	assert.True(t, blank2, "Should be true")
	blank3 := IsBlank("a")
	assert.False(t, blank3, "Should be false")
	blank4 := IsBlank("abcd")
	assert.False(t, blank4, "Should be false")
	blank5 := IsBlank("ab cd")
	assert.True(t, blank5, "Should be true")
	assert.True(t, IsBlank("SH vip lc"), "Should be true")
	assert.False(t, IsBlank("我爱你"), "Should be false")
	assert.True(t, IsBlank("我爱 你"), "Should be true")
}

// Test UStringUtil.IsBlankIfStr()
func TestIsBlankIfStr(t *testing.T) {
	assert.False(t, IsBlankIfStr(1), "Should be false")
	assert.False(t, IsBlankIfStr(true), "Should be false")
	assert.False(t, IsBlankIfStr("1"), "Should be false")
	assert.True(t, IsBlankIfStr(" "), "Should be true")
	assert.True(t, IsBlankIfStr2(" "), "Should be true")
}

// Test UStringUtil.IsNotBlank()
func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""), "Should be false")
	assert.False(t, IsNotBlank(" "), "Should be false")
	assert.False(t, IsNotBlank("我爱 你"), "Should be false")
	assert.True(t, IsNotBlank("我爱你"), "Should be true")
}

// Test UStringUtil.IsNotBlankIfStr()
func TestIsNotBlankIfStr(t *testing.T) {
	assert.False(t, IsNotBlankIfStr(1), "Should be false")
	assert.False(t, IsNotBlankIfStr(false), "Should be false")
	assert.True(t, IsNotBlankIfStr("1"), "Should be true")
	assert.False(t, IsNotBlankIfStr(" "), "Should be false")
}

// Test UStringUtil.IsEmpty()
func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""), "Should be true")
	assert.False(t, IsEmpty(" "), "Should be false")
	assert.False(t, IsEmpty("SH vip lc"), "Should be false")
	assert.False(t, IsEmpty("我爱你"), "Should be false")
	assert.False(t, IsEmpty("我爱 你"), "Should be false")
}

// Test UStringUtil.IsEmptyIfStr()
func TestIsEmptyIfStr(t *testing.T) {
	assert.False(t, IsEmptyIfStr(1), "Should be false")
	assert.True(t, IsEmptyIfStr(""), "Should be true")
	assert.False(t, IsEmptyIfStr(" "), "Should be false")
	assert.False(t, IsEmptyIfStr("SH vip lc"), "Should be false")
	assert.False(t, IsEmptyIfStr("我爱你"), "Should be false")
	assert.False(t, IsEmptyIfStr("我爱 你"), "Should be false") // 断言为false
}

// Test UStringUtil.IsNotEmpty()
func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""), "Should be false")
	assert.True(t, IsNotEmpty(" "), "Should be true")
	assert.True(t, IsNotEmpty("SH vip lc"), "Should be true")
	assert.True(t, IsNotEmpty("我爱你"), "Should be true")
	assert.True(t, IsNotEmpty("我爱 你"), "Should be true") // 断言为true
}

// Test UStringUtil.IsNotEmptyIfStr()
func TestIsNotEmptyIfStr(t *testing.T) {
	assert.False(t, IsEmptyIfStr(1), "Should be false")
	assert.False(t, IsNotEmptyIfStr(""), "Should be false")
	assert.True(t, IsNotEmptyIfStr(" "), "Should be true")
	assert.True(t, IsNotEmptyIfStr("SH vip lc"), "Should be true")
	assert.True(t, IsNotEmptyIfStr("我爱你"), "Should be true")
	assert.True(t, IsNotEmptyIfStr("我爱 你"), "Should be true")
}

// Test UStringUtil.SnakeString()
func TestUStringUtilSnakeString(t *testing.T) {
	assert.Equal(t, "to_say", SnakeString("ToSay"), "Should be equal") // 断言为二者相等
}

// Test UStringUtil.ToUpper()
func TestUStringUtilToUpper(t *testing.T) {
	assert.Equal(t, "AHVIPLC", ToUpper("ah VIP lc"), "Should be equal") // 断言为二者相等
}

// Test UStringUtil.ToLower()
func TestUStringUtilToLower(t *testing.T) {
	assert.Equal(t, "shviplc", ToLower("SH vip lc"), "Should be equal") // 断言为二者相等
	// assert equality
	assert.Equal(t, "shviplc", ToLower("SH vip lc"), "they should be equal") // 断言为二者相等
	// assert inequality
	assert.NotEqual(t, "shviplc not be equal", ToLower("SH vip lc"), "they should not be equal") // 断言为二者不相等
}

//Test StringInSlice
func TestStringInSlice(t *testing.T) {
	assert.True(t, StringInSlice("json", []string{"json", "ini"}), "Should be true")
	assert.True(t, StringInSlice("ini", []string{"json", "ini"}), "Should be true")
}
