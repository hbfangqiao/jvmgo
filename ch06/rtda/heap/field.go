package heap

import "jvmgo/ch06/classfile"
/*字段信息 存入方法区*/
type Field struct {
	ClassMember
	constValueIndex uint //static final修饰的字段在编译期就知道值了，该字段指向该值在常量池的Index
	slotId          uint //给字段一个序号，让Class知道自身的字段所对应的Field
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)//复制访问标识，名字，描述符
		fields[i].copyAttributes(cfField)//从字段表中读取constValueIndex
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}