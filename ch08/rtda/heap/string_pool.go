package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

// todo
// go string -> java.lang.String
func JString(loader *ClassLoader, goStr string) *Object {
	//如果 Java字符串已经在池中，直接返回
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	//否则先把Go字符串（UTF8 格式）转换成Java字符数组（UTF16格式）
	chars := stringToUtf16(goStr)
	//创建一个Java字符串 实例，把它的value变量设置成刚刚转换而来的字符数组
	jChars := &Object{loader.LoadClass("[C"), chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	//将jStr的char[] value 赋值为 jChars
	jStr.SetRefVar("value", "[C", jChars)
	//把 Java字符串放入池中。
	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
