package heap

import "jvmgo/ch06/classfile"
/*字段信息 存入方法区*/
type Field struct{
	ClassMember
}

func newFields(class *Class,cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field,len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)//复制访问标识，名字，描述符
	}
	return fields
}
