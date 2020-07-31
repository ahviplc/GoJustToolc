package UStringUtil

import (
	"fmt"
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

// Test UStringUtil.SnakeString() 推荐使用SnakeString2() 更加合理
func TestSnakeString(t *testing.T) {
	assert.Equal(t, "to_say", SnakeString("ToSay"), "Should be equal") // 断言为二者相等
	// more test
	data := [][2]string{
		{"XxYy", "xx_yy"},
		{"_XxYy", "_xx_yy"},
		{"TcpRpc", "tcp_rpc"},
		{"ID", "i_d"},
		{"UserID", "user_i_d"},
		{"RPC", "r_p_c"},
		{"TCP_RPC", "t_c_p__r_p_c"},
		{"wakeRPC", "wake_r_p_c"},
		{"_TCP__RPC", "_t_c_p___r_p_c"},
		{"_TcP__RpC_", "_tc_p___rp_c_"},
	}
	for _, p := range data {
		r := SnakeString(p[0])
		assert.Equal(t, p[1], r, p[0])
		r = SnakeString(p[1])
		assert.Equal(t, p[1], r, p[0])
	}
}

// Test UStringUtil.CamelString() 等价于方法 CamelString2()
func TestCamelString(t *testing.T) {
	data := [][2]string{
		{"_", "_"},
		{"xx_yy", "XxYy"},
		{"_xx_yy", "_XxYy"},
		{"id", "Id"},
		{"user_id", "UserId"},
		{"rpc", "Rpc"},
		{"tcp_rpc", "TcpRpc"},
		{"wake_rpc", "WakeRpc"},
		{"_tcp___rpc", "_Tcp__Rpc"},
		{"_tc_p__rp_c__", "_TcP_RpC__"},
	}
	for _, p := range data {
		r := CamelString(p[0])
		assert.Equal(t, p[1], r, p[0])
		r = CamelString(p[1])
		assert.Equal(t, p[1], r, p[0])
	}
}

// Test UStringUtil.CamelStringFirstLower()
func TestCamelStringFirstLower(t *testing.T) {
	fmt.Println(CamelStringFirstLower("tcp_rpc")) // tcpRpc
}

// Test UStringUtil.ToUpper()
func TestToUpper(t *testing.T) {
	assert.Equal(t, "AHVIPLC", ToUpper("ah VIP lc"), "Should be equal") // 断言为二者相等
}

// Test UStringUtil.ToLower()
func TestToLower(t *testing.T) {
	assert.Equal(t, "shviplc", ToLower("SH vip lc"), "Should be equal") // 断言为二者相等
	// assert equality
	assert.Equal(t, "shviplc", ToLower("SH vip lc"), "they should be equal") // 断言为二者相等
	// assert inequality
	assert.NotEqual(t, "shviplc not be equal", ToLower("SH vip lc"), "they should not be equal") // 断言为二者不相等
}

// Test UStringUtil.StringInSlice()
func TestStringInSlice(t *testing.T) {
	assert.True(t, StringInSlice("json", []string{"json", "ini"}), "Should be true")
	assert.True(t, StringInSlice("ini", []string{"json", "ini"}), "Should be true")
}

// Test UStringUtil.TrimRightSpace()
func TestTrimRightSpace(t *testing.T) {
	assert.Equal(t, "shviplc", TrimRightSpace("shviplc\n"), "they should be equal") // 断言为二者相等
}

// Test UStringUtil.MakeRandomString()
func TestMakeRandomString(t *testing.T) {
	fmt.Println(MakeRandomString(5)) // SauLP
	fmt.Println(MakeRandomString(8)) // SauLP7yP
}

// Test UStringUtil.Indent()
func TestIndent(t *testing.T) {
	fmt.Println(Indent("abc", "123")) // 123abc
}

// Test UStringUtil.BytesToString()
func TestBytesToString(t *testing.T) {
	bb := []byte("testing: BytesToString")
	ss := BytesToString(bb)
	t.Logf("type: %T, value: %v", ss, ss)
}

// Test UStringUtil.StringToBytes()
func TestStringToBytes(t *testing.T) {
	s := "testing: StringToBytes"
	b := StringToBytes(s)
	t.Logf("type: %T, value: %v, val-string: %s\n", b, b, b)
	b = append(b, '!')
	t.Logf("after append:\ntype: %T, value: %v, val-string: %s\n", b, b, b)
}

// Test UStringUtil.SnakeString2()
func TestSnakeString2(t *testing.T) {
	data := [][2]string{
		{"XxYy", "xx_yy"},
		{"_XxYy", "_xx_yy"},
		{"TcpRpc", "tcp_rpc"},
		{"ID", "id"},
		{"UserID", "user_id"},
		{"RPC", "rpc"},
		{"TCP_RPC", "tcp_rpc"},
		{"wakeRPC", "wake_rpc"},
		{"_TCP__RPC", "_tcp__rpc"},
		{"_TcP__RpC_", "_tc_p__rp_c_"},
	}
	for _, p := range data {
		r := SnakeString2(p[0])
		assert.Equal(t, p[1], r, p[0])
		r = SnakeString2(p[1])
		assert.Equal(t, p[1], r, p[0])
	}
}

// Test UStringUtil.CamelString2()
func TestCamelString2(t *testing.T) {
	data := [][2]string{
		{"_", "_"},
		{"xx_yy", "XxYy"},
		{"_xx_yy", "_XxYy"},
		{"id", "Id"},
		{"user_id", "UserId"},
		{"rpc", "Rpc"},
		{"tcp_rpc", "TcpRpc"},
		{"wake_rpc", "WakeRpc"},
		{"_tcp___rpc", "_Tcp__Rpc"},
		{"_tc_p__rp_c__", "_TcP_RpC__"},
	}
	for _, p := range data {
		r := CamelString2(p[0])
		assert.Equal(t, p[1], r, p[0])
		r = CamelString2(p[1])
		assert.Equal(t, p[1], r, p[0])
	}
}

// Test UStringUtil.LintCamelString()
func TestLintCamelString(t *testing.T) {
	data := [][2]string{
		{"_", "_"},
		{"xx_yy", "XxYy"},
		{"_xx_yy", "XxYy"},
		{"id", "ID"},
		{"user_id", "UserID"},
		{"rpc", "RPC"},
		{"tcp_rpc", "TCPRPC"},
		{"wake_rpc", "WakeRPC"},
		{"___tcp___rpc", "TCPRPC"},
		{"_tc_p__rp_c__", "TcPRpC"},
	}
	for _, p := range data {
		r := LintCamelString(p[0])
		assert.Equal(t, p[1], r, p[0])
		r = LintCamelString(p[1])
		assert.Equal(t, p[1], r, p[0])
	}
}

// Test UStringUtil.HTMLEntityToUTF8()
func TestHTMLEntityToUTF8(t *testing.T) {
	want := `{"info":[["color","咖啡色|绿色"]]｝`
	got := HTMLEntityToUTF8(`{"info":[["color","&#5496;&#5561;&#8272;&#7c;&#7eff;&#8272;"]]｝`, 16)
	if got != want {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

// Test UStringUtil.HTMLEntityToUTF8()
func TestCodePointToUTF8(t *testing.T) {
	got := CodePointToUTF8(`{"info":[["color","\u5496\u5561\u8272\u7c\u7eff\u8272"]]｝`, 16)
	want := `{"info":[["color","咖啡色|绿色"]]｝`
	if got != want {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

// Test UStringUtil.SpaceInOne()
func TestSpaceInOne(t *testing.T) {
	a := struct {
		input  string
		output string
	}{
		input: `# authenticate method 
		//  comment2	
		/*  some other 
			  comments */
		`,
		output: `# authenticate method
	// comment2
	/* some other
	comments */
	`,
	}
	r := SpaceInOne(a.input)
	if r != a.output {
		t.Fatalf("want: %q, got: %q", a.output, r)
	}
}

// Test UStringUtil.StringMarshalJSON()
func TestStringMarshalJSON(t *testing.T) {
	s := `<>&{}""`
	fmt.Printf("%s\n", StringMarshalJSON(s, true))
	fmt.Printf("%s\n", StringMarshalJSON(s, false))
	// Output:
	// "\u003c\u003e\u0026{}\"\""
	// "<>&{}\"\""
}
