package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	//声明字段的类还没有初始化好，也需要先初始化
	class := field.Class()

	//类还没初始化，则先初始化类
	if !class.InitStarted() {
		frame.RevertNextPC()//将指令重新指向当前指令
		base.InitClass(frame.Thread(), class)//初始化类
		return//终止执行当前指令
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//getstatic读取静态变量的值，不用管是否是final
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))

	}
}
