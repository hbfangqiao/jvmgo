package heap

import "jvmgo/ch06/classfile"
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
