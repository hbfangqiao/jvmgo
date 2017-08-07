package classfile
type ConstantNameAndTypeInfo struct {
	nameIndex		uint16
	descriptorIndex		uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	//字段描述符或方法描述符
	self.descriptorIndex = reader.readUint16()
}
