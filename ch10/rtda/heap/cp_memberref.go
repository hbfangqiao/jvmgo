package heap

import "jvmgo/ch10/classfile"
/*存放 字段和方法符号引用共有的信息*/
type MemberRef struct {
	SymRef
	name       string
	descriptor string //字段也要存放描述符，同一个类中不能有相同的name只是java语法的限制
}
/*从class文件内存储的字段或方法常量中提取数据*/
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo){
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

//getter
func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}