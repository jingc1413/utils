package utils

// 先转为字符串再转为 []byte, 可选指定默认值
func GetBytes(v interface{}, defaultVal ...[]byte) []byte {
	switch b := v.(type) {
	default:
		bs := S2B(MustString(v))
		if len(bs) == 0 && len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return bs
	case []byte:
		return b
	}
}

// Immutable, 可选指定默认值
func GetSafeBytes(b []byte, defaultVal ...[]byte) []byte {
	b = CopyBytes(b)
	if len(b) == 0 && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return b
}

// Immutable, 可选指定默认值
func GetSafeS2B(s string, defaultVal ...[]byte) []byte {
	b := []byte(s)
	if len(b) == 0 && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return b
}

// Immutable, []byte to []byte
func CopyBytes(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

// Immutable, string to []byte
func CopyS2B(s string) []byte {
	// []byte(s)
	return CopyBytes(S2B(s))
}

// 拼接 []byte
func JoinBytes(b ...[]byte) []byte {
	var res []byte
	for _, v := range b {
		res = append(res, v...)
	}
	return res
}
