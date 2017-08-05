package heap

import "jvmgo/ch07/classfile"
/*常量池中会存储本类所引用的类的类名，将所引用的类解析出来缓存*/
type ClassRef struct {
	SymRef
}
/*根据class文件中存储的类常量创建ClassRef实例*/
func newClassRef(cp *ConstantPool,classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
