package heap

import "jvmgo/ch06/classfile"

/*本类中所引用的方法，包含所属的类 自己的类名，以及修饰符 存放在constant_pool中，*/
type MethodRef struct {
	MemberRef
	Method *Method//缓存解析后的Method
}

func newMethodRef(cp *ConstantPool,refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
