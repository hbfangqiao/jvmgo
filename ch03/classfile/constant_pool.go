package classfile

type ConstantPool []ConstantInfo
type ConstantInfo interface {
	//读取常量信息，由具体的结构体实现
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo{
	//读取出tag值
	tag := reader.readUint8()
	//创建具体常量
	c := newConstantInfo(tag, cp)
	//调用readInfo读取常量信息
	c.readInfo(reader)
	return c
}
/*根据tag值创建具体的常量*/
func newConstantInfo(tag uint8,cp ConstantPool) ConstantInfo{
	switch tag {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return  &ConstantFloatInfo{}
	case CONSTANT_Long: return  &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}
	case CONSTANT_Utf8: return &ConstantUtf8Info{}
	case CONSTANT_String: return &ConstantStringInfo{}
	case CONSTANT_Class: return  &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp :cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp :cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp :cp}}
	case CONSTANT_NameAndType:return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:return &ConstantInvokeDynamicInfo{}
	default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

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
