package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type DSUB struct { base.NoOperandsInstruction }
/*double相减 栈顶的是减数*/
func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

type FSUB struct { base.NoOperandsInstruction }
/*float相减 栈顶是减数*/
func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

type ISUB struct { base.NoOperandsInstruction }
/*Int相减 栈顶是减数*/
func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

type LSUB struct { base.NoOperandsInstruction }
/*Long相减 栈顶是减数*/
func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
