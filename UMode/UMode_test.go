package UMode

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv(EnvGoJustToolcMode, TestMode)
}

// Test SetMode() also in UConsole/UConsole_test.go:47
func TestSetMode(t *testing.T) {
	assert.Equal(t, TestCode, UMode)
	assert.Equal(t, TestMode, Mode())
	os.Unsetenv(EnvGoJustToolcMode)

	SetMode("")
	assert.Equal(t, DebugCode, UMode)
	assert.Equal(t, DebugMode, Mode())

	SetMode(DebugMode)
	assert.Equal(t, DebugCode, UMode)
	assert.Equal(t, DebugMode, Mode())

	SetMode(ReleaseMode)
	assert.Equal(t, ReleaseCode, UMode)
	assert.Equal(t, ReleaseMode, Mode())

	SetMode(TestMode)
	assert.Equal(t, TestCode, UMode)
	assert.Equal(t, TestMode, Mode())

	assert.Panics(t, func() { SetMode("unknown") })
}
