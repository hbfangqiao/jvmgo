package constants

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/**从运行时常量池中加载常量值，并把它推入操作数栈
ldc和ldc_w指令用 于加载int、float和字符串常量，java.lang.Class实例或者MethodType 和MethodHandle实例
ldc2_w指令用于加载long和double常量。ldc 和ldc_w指令的区别仅在于操作数的宽度
*/
type LDC struct {
	base.Index8Instruction
}

func (self *LDC) Execute(frame *rtda.Frame)  {
	_ldc(frame, self.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (self *LDC_W) Execute(frame *rtda.Frame)  {
	_ldc(frame, self.Index)
}

type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC2_W) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64: stack.PushLong(c.(int64))
	case float64: stack.PushDouble(c.(float64))
	default: panic("java.lang.ClassFormatError")   }
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	default:
		panic("todo: ldc!")

	}
}
