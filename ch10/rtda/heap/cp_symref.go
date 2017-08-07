package heap
/*符号引用  由于ClassRef，FieldRef，MethodRef,InterfaceMethodRef 的符号引用有一些共性，使用继承*/
type SymRef struct {
	cp        *ConstantPool //所在的运行时常量池指针 通过符号引用可访问运行时常量池，进一步又可以访问到类数据。
	className string        //类的完全限定名
	class     *Class        //缓存解析后的类的结构体指针，这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值
}

func (self *SymRef) ResolvedClass() *Class {
	//如果class没有被解析，则解析自身（）所属类的引用
	if self.class == nil {
		self.resolvedClassRef()
	}
	return self.class
}
/**/
func (self *SymRef) resolvedClassRef() {
	//d通过符号引用N 引用类c 要解析N 先用d的类加载器加载C，然后检查d是否有权限访问c
	d := self.cp.class //这个Ref所属于的类
	c := d.loader.LoadClass(self.className) //类所指向的类加载器，加载这个类
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}