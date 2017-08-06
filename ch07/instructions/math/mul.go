package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type DMUL struct{ base.NoOperandsInstruction }
/*double相乘*/
func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

type FMUL struct{ base.NoOperandsInstruction }
/*float 相乘*/
func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

type IMUL struct{ base.NoOperandsInstruction }
/*int 相乘*/
func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

type LMUL struct{ base.NoOperandsInstruction }
/*long 相乘*/
func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}