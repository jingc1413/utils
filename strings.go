package utils

// 获取字符串结果, 可选指定默认值
func GetString(v interface{}, defaultVal ...string) string {
	s := MustString(v)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// Immutable, 可选指定默认值
func GetSafeString(s string, defaultVal ...string) string {
	s = CopyString(s)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// Immutable, 可选指定默认值
func GetSafeB2S(b []byte, defaultVal ...string) string {
	s := string(b)
	if s == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return s
}

// Immutable, string to string
// e.g. fiberParam := utils.CopyString(c.Params("test"))
// e.g. utils.CopyString(s[500:1000]) // 可以让 s 被 GC 回收
func CopyString(s string) string {
	return B2S([]byte(s))
}

// Immutable, []byte to string
func CopyB2S(b []byte) string {
	// string(b)
	return B2S(CopyBytes(b))
}

// 拼接字符串, 返回 bytes from bytes.Join()
func AddStringBytes(s ...string) []byte {
	switch len(s) {
	case 0:
		return []byte{}
	case 1:
		return []byte(s[0])
	}

	n := 0
	for _, v := range s {
		n += len(v)
	}

	b := make([]byte, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], v)
	}

	return b
}

// 拼接字符串
func AddString(s ...string) string {
	return B2S(AddStringBytes(s...))
}

// 搜索字符串位置(左, 第一个)
func SearchString(ss []string, s string) int {
	for i, v := range ss {
		if s == v {
			return i
		}
	}
	return -1
}

// 检查字符串是否存在于 slice
func InStrings(ss []string, s string) bool {
	return SearchString(ss, s) != -1
}
