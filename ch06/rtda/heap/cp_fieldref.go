package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field//缓存解析后的字段指针
}
/*创建FieldRef实例*/
func newFieldRef(cp *ConstantPool,refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
