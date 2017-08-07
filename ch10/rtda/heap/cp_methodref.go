package heap

import "jvmgo/ch10/classfile"

/*本类中所引用的方法，包含所属的类 自己的类名，以及修饰符 存放在constant_pool中，*/
type MethodRef struct {
	MemberRef
	method *Method//缓存解析后的Method
}

func newMethodRef(cp *ConstantPool,refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

/*如果还没有解析过符号引用，调用resolveMethodRef（）方法进 行解析，否则直接返回方法指针*/
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	//方法引用自身所属的类
	d := self.cp.class
	//要调用方法所属的类
	c := self.ResolvedClass()
	//C是接口，则抛出 IncompatibleClassChangeError异常
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//根据1.方法所属的类 2.方法名 3.描述符查找方法
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {//查找不到
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {//本类没有权限访问method
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	//先从class的继承层次中找，
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {//如果找不到，就去class的接口中找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}