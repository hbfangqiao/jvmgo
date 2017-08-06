package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	//如果是数组类名，描述符就是类名，直接返回即可。
	if className[0] == '[' {
		// array
		return className
	}
	//如果是基本类型名，返回对应的类型描述符，
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// object 否则肯定是普通的类名，前面 加上方括号，结尾加上分号即可得到类型描述符
	return "L" + className + ";"
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	//如果类型描述符以方括号开头，那么肯定是数组，描述符即是类名
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	//如果类型描述符以L开头，那么肯定是类描述符，去掉开头的 L和末尾的分号即是类名
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor)-1]
	}
	//否则判断是否是基本类型的描述符，如 果是，返回基本类型名称
	for className, d := range primitiveTypes {
		if d == descriptor {
			// primitive
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
