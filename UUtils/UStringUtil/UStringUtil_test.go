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

// ----------------------------------------------------------------------------------------

// Test UStringUtil.MD5()
func TestMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "iiinsomnia"},
			want: "483367436bc9a6c5256bfc29a24f955e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test UStringUtil.SHA1()
func TestSHA1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "iiinsomnia"},
			want: "7a4082bd79f2086af2c2b792c5e0ad06e729b9c4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test UStringUtil.Hash()
func TestHash(t *testing.T) {
	type args struct {
		algo HashAlgo
		s    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{algo: AlgoMD5, s: "admin"},
			want: "21232f297a57a5a743894a0e4a801fc3",
		},
		{
			name: "sha1",
			args: args{algo: AlgoSha1, s: "admin"},
			want: "d033e22ae348aeb5660fc2140aec35850c4da997",
		},
		{
			name: "sha224",
			args: args{algo: AlgoSha224, s: "admin"},
			want: "58acb7acccce58ffa8b953b12b5a7702bd42dae441c1ad85057fa70b",
		},
		{
			name: "sha256",
			args: args{algo: AlgoSha256, s: "admin"},
			want: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
		},
		{
			name: "sha384",
			args: args{algo: AlgoSha384, s: "admin"},
			want: "9ca694a90285c034432c9550421b7b9dbd5c0f4b6673f05f6dbce58052ba20e4248041956ee8c9a2ec9f10290cdc0782",
		},
		{
			name: "sha512",
			args: args{algo: AlgoSha512, s: "admin"},
			want: "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.algo, tt.args.s); got != tt.want {
				t.Errorf("Hash(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// Test UStringUtil.HMAC()
func TestHMAC(t *testing.T) {
	hmac := HMAC("md5", "LC", "123456")
	assert.Equal(t, "36fb7c781b238655b0fe8fc4d692f596", hmac, "Should be equal") // 断言为二者相等
}

// Test UStringUtil.AddSlashes()
func TestAddSlashes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "Is your name O'Reilly?"},
			want: `Is your name O\'Reilly?`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSlashes(tt.args.s); got != tt.want {
				t.Errorf("AddSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test UStringUtil.StripSlashes()
func TestStripSlashes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: `Is your name O\'reilly?`},
			want: "Is your name O'reilly?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripSlashes(tt.args.s); got != tt.want {
				t.Errorf("StripSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test UStringUtil.QuoteMeta()
func TestQuoteMeta(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "Hello world. (can you hear me?)"},
			want: `Hello world\. \(can you hear me\?\)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuoteMeta(tt.args.s); got != tt.want {
				t.Errorf("QuoteMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ----------------------------------------------------------------------------------------
