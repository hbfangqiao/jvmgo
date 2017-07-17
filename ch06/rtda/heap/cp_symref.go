package heap
/*符号引用  由于ClassRef，FieldRef，MethodRef,InterfaceMethodRef 的符号引用有一些共性，使用继承*/
type SymRef struct {
	cp        *ConstantPool //所在的运行时常量池指针 通过符号引用可访问运行时常量池，进一步又可以访问到类数据。
	className string        //类的完全限定名
	class     *Class        //缓存解析后的类的结构体指针，这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值
}