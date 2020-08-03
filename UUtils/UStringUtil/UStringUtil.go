package UStringUtil

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// 字符串是否为空白 空白的定义如下:
// 1.为null
// 2.为不可见字符 (如空格)
// 3.""
func IsBlank(in string) bool {
	if len(in) == 0 || in == "" {
		return true
	}
	inRune := []rune(in)
	for _, r := range inRune {
		// 是否为空白符号
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

// 如果对象是字符串是否为空白 空白的定义如下:
// 1.为null
// 2.为不可见字符 (如空格)
// 3.""
func IsBlankIfStr(in interface{}) bool {
	_, isFlag := in.(string) // in2, isFlag := in.(string)
	if isFlag {
		return IsBlank(in.(string)) // return IsBlank(in2)
	}
	return false
}

// IsBlankIfStr 写法2 方法效果一样
func IsBlankIfStr2(in interface{}) bool {
	in2, isFlag := in.(string)
	if isFlag {
		return IsBlank(in2)
	}
	return false
}

// 字符串是否为非空白 非空白的定义如下:
// 1.不为null
// 2.不为不可见字符 (如空格)
// 3.不为""
func IsNotBlank(in string) bool {
	return false == IsBlank(in)
}

// 如果对象是字符串是否为非空白 非空白的定义如下:
// 1.不为null
// 2.不为不可见字符 (如空格)
// 3.不为""
func IsNotBlankIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsNotBlank(in.(string))
	}
	return false
}

// 字符串是否为空 空的定义如下:
// 1.为null
// 2.""
func IsEmpty(in string) bool {
	return len(in) == 0 || in == ""
}

// 如果对象是字符串是否为空字符串 空的定义如下:
// 1.为null
// 2.""
func IsEmptyIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsEmpty(in.(string))
	}
	return false
}

// 字符串是否为非空 非空的定义如下:
// 1.不为null
// 2.不为""
func IsNotEmpty(in string) bool {
	return false == IsEmpty(in)
}

// 如果对象是字符串是否为非空字符串 非空的定义如下:
// 1.不为null
// 2.不为""
func IsNotEmptyIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsNotEmpty(in.(string))
	}
	return false
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func CamelStringFirstLower(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		if i == 0 && d >= 'A' && d <= 'Z' {
			d = d + 32
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 字符串去除空格并将所有字母大写
func ToUpper(oldData string) string {
	return strings.ToUpper(strings.Replace(oldData, " ", "", -1))
}

// 字符串去除空格并将所有字母小写
func ToLower(oldData string) string {
	return strings.ToLower(strings.Replace(oldData, " ", "", -1))
}

// int转string
func IntToString(in int) string {
	return strconv.Itoa(in)
}

// string转int
func StringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

// int64转string
func Int64ToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

// string转int64
func StringToInt64(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return out
}

// 判断某个字符串在不在某个字符串切片里 true 在 false 不在
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// 去掉换行,空格,回车,制表
// 去掉【\r \n \t 空格】
func TrimRightSpace(in string) string {
	return strings.TrimRight(string(in), "\r\n\t ")
}

// 生成随机字符串
// length 随机字符串长度
func MakeRandomString(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytesTemp := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytesTemp[r.Intn(len(bytesTemp))])
	}
	return string(result)
}

// Indent inserts prefix at the beginning of each line
func Indent(text, prefix string) string {
	if len(prefix) == 0 {
		return text
	}
	has := strings.HasSuffix(text, "\n")
	text = prefix + strings.Replace(text, "\n", "\n"+prefix, -1)
	if has {
		return text[:len(text)-len(prefix)]
	}
	return text
}

// BytesToString convert []byte type to string type.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes convert string type to []byte type.
// NOTE: panic if modify the member value of the []byte.
func StringToBytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

// SnakeString2 converts the accepted string to a snake string (XxYy to xx_yy)
func SnakeString2(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	for _, d := range StringToBytes(s) {
		if d >= 'A' && d <= 'Z' {
			if j {
				data = append(data, '_')
				j = false
			}
		} else if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(BytesToString(data))
}

// CamelString2 converts the accepted string to a camel string (xx_yy to XxYy)
func CamelString2(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return BytesToString(data[:])
}

// LintCamelString converts the accepted string to a camel string (xx_id to XxID)
// NOTE:
//  support common initialisms
func LintCamelString(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return "_"
	}
	runes := []rune(name)
	var i int
	for k, v := range runes {
		if v != '_' {
			i = k
			runes[k] = unicode.ToUpper(v)
			break
		}
	}
	r := string(toInitialisms(runes[i:]))
	return r
}

func toInitialisms(runes []rune) []rune {
	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return runes
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

var htmlEntityRegexp = regexp.MustCompile(`&#([0-9a-zA-Z]+);*`)

// HTMLEntityToUTF8 converts HTML Unicode to UTF-8.
// e.g.: HTMLEntityToUTF8(`{"info":[["color","&#5496;&#5561;&#8272;&#7c;&#7eff;&#8272;"]]｝`, 16)
// => `{"info":[["color","咖啡色|绿色"]]｝`
func HTMLEntityToUTF8(str string, base int) string {
	a := htmlEntityRegexp.FindAllStringSubmatch(str, -1)
	if len(a) == 0 {
		return str
	}
	oldnew := make([]string, 0, len(a)*2)
	for _, s := range a {
		if i, err := strconv.ParseInt(s[1], base, 32); err == nil {
			oldnew = append(oldnew, s[0], string(i))
		}
	}
	r := strings.NewReplacer(oldnew...)
	return r.Replace(str)
}

// CodePointToUTF8 converts Unicode Code Point to UTF-8.
// e.g.: CodePointToUTF8(`{"info":[["color","\u5496\u5561\u8272\u7c\u7eff\u8272"]]｝`, 16)
// => `{"info":[["color","咖啡色|绿色"]]｝`
func CodePointToUTF8(str string, base int) string {
	i := 0
	if strings.Index(str, `\u`) > 0 {
		i = 1
	}
	strSlice := strings.Split(str, `\u`)
	last := len(strSlice) - 1
	if len(strSlice[last]) > 4 {
		strSlice = append(strSlice, string(strSlice[last][4:]))
		strSlice[last] = string(strSlice[last][:4])
	}
	for ; i <= last; i++ {
		if x, err := strconv.ParseInt(strSlice[i], base, 32); err == nil {
			strSlice[i] = string(x)
		}
	}
	return strings.Join(strSlice, "")
}

var spaceReplacer = strings.NewReplacer(
	"  ", " ",
	"\n\n", "\n",
	"\r\r", "\r",
	"\t\t", "\t",
	"\r\n\r\n", "\r\n",
	" \n", "\n",
	"\t\n", "\n",
	" \t", "\t",
	"\t ", "\t",
	"\v\v", "\v",
	"\f\f", "\f",
	string(0x85)+string(0x85),
	string(0x85),
	string(0xA0)+string(0xA0),
	string(0xA0),
)

// SpaceInOne combines multiple consecutive space characters into one.
func SpaceInOne(s string) string {
	var old string
	for old != s {
		old = s
		s = spaceReplacer.Replace(s)
	}
	return s
}

// StringMarshalJSON converts the string to JSON byte stream.
func StringMarshalJSON(s string, escapeHTML bool) []byte {
	a := StringToBytes(s)
	var buf = bytes.NewBuffer(make([]byte, 0, 64))
	buf.WriteByte('"')
	start := 0
	for i := 0; i < len(a); {
		if b := a[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				buf.Write(a[start:i])
			}
			switch b {
			case '\\', '"':
				buf.WriteByte('\\')
				buf.WriteByte(b)
			case '\n':
				buf.WriteByte('\\')
				buf.WriteByte('n')
			case '\r':
				buf.WriteByte('\\')
				buf.WriteByte('r')
			case '\t':
				buf.WriteByte('\\')
				buf.WriteByte('t')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				buf.WriteString(`\u00`)
				buf.WriteByte(hexSet[b>>4])
				buf.WriteByte(hexSet[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRune(a[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				buf.Write(a[start:i])
			}
			buf.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				buf.Write(a[start:i])
			}
			buf.WriteString(`\u202`)
			buf.WriteByte(hexSet[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(a) {
		buf.Write(a[start:])
	}
	buf.WriteByte('"')
	return buf.Bytes()
}

var hexSet = "0123456789abcdef"

// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
var safeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      true,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      true,
	'=':      true,
	'>':      true,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// htmlSafeSet holds the value true if the ASCII character with the given
// array position can be safely represented inside a JSON string, embedded
// inside of HTML <script> tags, without any additional escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), the backslash character ("\"), HTML opening and closing
// tags ("<" and ">"), and the ampersand ("&").
var htmlSafeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      false,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      false,
	'=':      true,
	'>':      false,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// ----------------------------------------------------------------------------------------

// from:https://github.com/shenghui0779/yiigo/blob/master/strings.go

type HashAlgo string

const (
	AlgoMD5    HashAlgo = "md5"
	AlgoSha1   HashAlgo = "sha1"
	AlgoSha224 HashAlgo = "sha224"
	AlgoSha256 HashAlgo = "sha256"
	AlgoSha384 HashAlgo = "sha384"
	AlgoSha512 HashAlgo = "sha512"
)

// MD5 calculate the md5 hash of a string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// Hash Generate a hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func Hash(algo HashAlgo, s string) string {
	var h hash.Hash

	switch algo {
	case AlgoMD5:
		h = md5.New()
	case AlgoSha1:
		h = sha1.New()
	case AlgoSha224:
		h = sha256.New224()
	case AlgoSha256:
		h = sha256.New()
	case AlgoSha384:
		h = sha512.New384()
	case AlgoSha512:
		h = sha512.New()
	default:
		return s
	}

	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// HMAC Generate a keyed hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func HMAC(algo HashAlgo, s, key string) string {
	var mac hash.Hash

	switch algo {
	case AlgoMD5:
		mac = hmac.New(md5.New, []byte(key))
	case AlgoSha1:
		mac = hmac.New(sha1.New, []byte(key))
	case AlgoSha224:
		mac = hmac.New(sha256.New224, []byte(key))
	case AlgoSha256:
		mac = hmac.New(sha256.New, []byte(key))
	case AlgoSha384:
		mac = hmac.New(sha512.New384, []byte(key))
	case AlgoSha512:
		mac = hmac.New(sha512.New, []byte(key))
	default:
		return s
	}

	mac.Write([]byte(s))

	return hex.EncodeToString(mac.Sum(nil))
}

// AddSlashes returns a string with backslashes added before characters that need to be escaped.
func AddSlashes(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		if ch == '\'' || ch == '"' || ch == '\\' {
			buf.WriteRune('\\')
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}

// StripSlashes returns a string with backslashes stripped off. (\' becomes ' and so on.) Double backslashes (\\) are made into a single backslash (\).
func StripSlashes(s string) string {
	var buf bytes.Buffer

	l, skip := len(s), false

	for i, ch := range s {
		if skip {
			buf.WriteRune(ch)
			skip = false

			continue
		}

		if ch == '\\' {
			if i+1 < l && s[i+1] == '\\' {
				skip = true
			}

			continue
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}

// QuoteMeta returns a version of str with a backslash character (\) before every character that is among these: . \ + * ? [ ^ ] ( $ )
func QuoteMeta(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		switch ch {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}

// ----------------------------------------------------------------------------------------
