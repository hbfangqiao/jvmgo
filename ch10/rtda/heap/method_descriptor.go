package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (self *MethodDescriptor) addParameterType(t string) {
	//获取描述符中 方法参数的个数
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {//cap()获取切片分配空间的大小
		s := make([]string, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}
	self.parameterTypes = append(self.parameterTypes, t)
}