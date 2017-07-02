package classfile

type MemberInfo struct {
	cp		ConstantPool//存储常量池指针
	accessFlags	uint16//访问控制符
	nameIndex	uint16
	descriptorIndex	uint16
	attributes	[]AttributeInfo
}
/*读取字段或者方法表*/
func readMembers(reader *ClassReader,cp ConstantPool)[]*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo,memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}
func readMember(reader *ClassReader,cp ConstantPool) *MemberInfo{
	return &MemberInfo{
		cp:			cp,
		accessFlags:		reader.readUint16(),
		nameIndex:		reader.readUint16(),
		descriptorIndex:	reader.readUint16(),
		attributes:		readAttributes(reader,cp),
	}
}
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
/*根据name索引在常量池中查找字段或者方法名*/
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
/*从常量池查找字段或方法描述符*/
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
