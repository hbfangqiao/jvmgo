package classfile

type ConstantPool []ConstantInfo





func readConstantPool(reader *ClassReader) ConstantPool{
	//表头给出的常量池大小比实际大1，如果有CONSTANT_Long_info，CONSTANT_Double_info，则实际常量数比n-1还要少
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1;i < cpCount; i++ {//索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++
		}
	}
	return cp

}
/*按照索引查找常量*/
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
/*从常量池查找字段或方法的名字或描述符*/
func (self ConstantPool) getNameAndType(index uint16) (string,string){
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}
/*从常量池中查找类名*/
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
/*从常量池查找UTF-8字符串*/
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
