package UUtils

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test UStringUtil.IsBlank()
func TestIsBlank(t *testing.T) {
	blank := IsBlank("")
	UConsole.Log(blank) // true
	blank2 := IsBlank(" ")
	UConsole.Log(blank2) // true
	blank3 := IsBlank("a")
	UConsole.Log(blank3) // false
	blank4 := IsBlank("abcd")
	UConsole.Log(blank4) // false
	blank5 := IsBlank("ab cd")
	UConsole.Log(blank5)               // true
	UConsole.Log(IsBlank("SH vip lc")) // true
	UConsole.Log(IsBlank("我爱你"))       // false
	UConsole.Log(IsBlank("我爱 你"))      // true
}

// Test UStringUtil.IsBlankIfStr()
func TestIsBlankIfStr(t *testing.T) {
	UConsole.Log(IsBlankIfStr(1))    // false
	UConsole.Log(IsBlankIfStr(true)) // false
	UConsole.Log(IsBlankIfStr("1"))  // false
	UConsole.Log(IsBlankIfStr(" "))  // true
	UConsole.Log(IsBlankIfStr2(" ")) // true
}

// Test UStringUtil.IsNotBlank()
func TestIsNotBlank(t *testing.T) {
	UConsole.Log(IsNotBlank(""))     // fales
	UConsole.Log(IsNotBlank(" "))    // fales
	UConsole.Log(IsNotBlank("我爱 你")) // fales
	UConsole.Log(IsNotBlank("我爱你"))  // true
}

// Test UStringUtil.IsNotBlankIfStr()
func TestIsNotBlankIfStr(t *testing.T) {
	UConsole.Log(IsNotBlankIfStr(1))     // false
	UConsole.Log(IsNotBlankIfStr(false)) // false
	UConsole.Log(IsNotBlankIfStr("1"))   // true
	UConsole.Log(IsNotBlankIfStr(" "))   // false
}

// Test UStringUtil.IsEmpty()
func TestIsEmpty(t *testing.T) {
	UConsole.Log(IsEmpty(""))          // true
	UConsole.Log(IsEmpty(" "))         // false
	UConsole.Log(IsEmpty("SH vip lc")) // false
	UConsole.Log(IsEmpty("我爱你"))       // false
	UConsole.Log(IsEmpty("我爱 你"))      // false
}

// Test UStringUtil.IsEmptyIfStr()
func TestIsEmptyIfStr(t *testing.T) {
	UConsole.Log(IsEmptyIfStr(1))                            // false
	UConsole.Log(IsEmptyIfStr(""))                           // true
	UConsole.Log(IsEmptyIfStr(" "))                          // false
	UConsole.Log(IsEmptyIfStr("SH vip lc"))                  // false
	UConsole.Log(IsEmptyIfStr("我爱你"))                        // false
	UConsole.Log(IsEmptyIfStr("我爱 你"))                       // false
	assert.False(t, IsEmptyIfStr("我爱 你"), "False is false!") // 断言为false
}

// Test UStringUtil.IsNotEmpty()
func TestIsNotEmpty(t *testing.T) {
	UConsole.Log(IsNotEmpty(""))                        // false
	UConsole.Log(IsNotEmpty(" "))                       // true
	UConsole.Log(IsNotEmpty("SH vip lc"))               // true
	UConsole.Log(IsNotEmpty("我爱你"))                     // true
	UConsole.Log(IsNotEmpty("我爱 你"))                    // true
	assert.True(t, IsNotEmpty("我爱 你"), "True is true!") // 断言为true
}

// Test UStringUtil.IsNotEmptyIfStr()
func TestIsNotEmptyIfStr(t *testing.T) {
	UConsole.Log(IsEmptyIfStr(1))              // false
	UConsole.Log(IsNotEmptyIfStr(""))          // false
	UConsole.Log(IsNotEmptyIfStr(" "))         // true
	UConsole.Log(IsNotEmptyIfStr("SH vip lc")) // true
	UConsole.Log(IsNotEmptyIfStr("我爱你"))       // true
	UConsole.Log(IsNotEmptyIfStr("我爱 你"))      // true
}

// Test UStringUtil.SnakeString()
func TestUStringUtilSnakeString(t *testing.T) {
	UConsole.Log(SnakeString("ToSay")) // to_say
}

// Test UStringUtil.ToUpper()
func TestUStringUtilToUpper(t *testing.T) {
	UConsole.Log(ToUpper("ah VIP lc")) // AHVIPLC
}

// Test UStringUtil.ToLower()
func TestUStringUtilToLower(t *testing.T) {
	UConsole.Log(ToLower("SH vip lc")) // shviplc

	// assert equality
	assert.Equal(t, "shviplc", ToLower("SH vip lc"), "they should be equal") // 断言为二者相等

	// assert inequality
	assert.NotEqual(t, "shviplc not be equal", ToLower("SH vip lc"), "they should not be equal") // 断言为二者不相等
}
