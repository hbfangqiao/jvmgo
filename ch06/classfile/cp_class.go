package classfile

/*constantPool中，类，超类或接口的符号引用，根据每个常量的u1 constant_info 来创建*/
type ConstantClassInfo struct {
	cp		ConstantPool
	nameIndex	uint16
}
/*读取常量池索引*/
func (self *ConstantClassInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
}
/*通过索引从常量池中读取符号*/
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}