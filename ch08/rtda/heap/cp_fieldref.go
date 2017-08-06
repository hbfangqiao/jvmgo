package heap

import "jvmgo/ch08/classfile"

type FieldRef struct {
	MemberRef
	field *Field //缓存解析后的字段指针
}
/*创建FieldRef实例*/
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	//虚拟机规范5.4.3.2节给出了解析步骤
	//访问类C的某个字段，首先要解 析符号引用得到类C，然后根据字段名和描述符查找字段。如果字段查找失败，
	//则虚拟机抛出NoSuchFieldError异常。如果查找成功， 但D没有足够的权限访问该字段，则虚拟机抛出
	//IllegalAccessError异常
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {//首先在C的字段中查找。
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range  c.interfaces {//如果找不到就在C的直接接口递归查找
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {//在超类中递归查找
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}