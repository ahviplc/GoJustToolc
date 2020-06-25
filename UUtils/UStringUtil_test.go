package UUtils

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test UStringUtil.ToUpper()
func TestUStringUtilToUpper(t *testing.T) {
	UConsole.Log(ToUpper("ah VIP lc"))
}

// Test UStringUtil.ToLower()
func TestUStringUtilToLower(t *testing.T) {
	UConsole.Log(ToLower("SH vip lc"))
}
