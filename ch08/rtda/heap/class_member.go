package heap

import "jvmgo/ch08/classfile"
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

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsPrivate() {//默认的访问权限
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}