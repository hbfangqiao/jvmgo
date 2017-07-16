package heap

import "jvmgo/ch06/classfile"
/*字段和方法相同的信息，相当于字段，方法的抽象类*/
type ClassMember struct {
	accessFlags	uint16 //访问标识
	name 		string //字段、方法名
	descriptor 	string //描述符
	class 		*Class //存放Class结构体指针，这样可以通过字段或方法访问到它所属的类
}
/*从class文件中复制MemberInfo的数据*/
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}