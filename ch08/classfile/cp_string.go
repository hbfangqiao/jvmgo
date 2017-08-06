package classfile

/*表示java.lang.String字面量，不存放字符串数据，只存了常量池索引*/
type ConstantStringInfo struct {
	cp		ConstantPool
	stringIndex	uint16
}
/*读取常量池索引*/
func (self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string{
	return self.cp.getUtf8(self.stringIndex)
}

