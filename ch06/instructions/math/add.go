package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type DADD struct{ base.NoOperandsInstruction }
//double相加
func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

type FADD struct{ base.NoOperandsInstruction }
//float相加
func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

type IADD struct { base.NoOperandsInstruction }
//int相加
func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

type LADD struct { base.NoOperandsInstruction }
/*long 相加*/
func (self *LADD) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
