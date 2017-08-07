package math

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type DNEG struct{ base.NoOperandsInstruction }
/*double取反*/
func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

type FNEG struct { base.NoOperandsInstruction }
/*float取反*/
func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type INEG struct{ base.NoOperandsInstruction }
/*int取反*/
func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNEG struct{ base.NoOperandsInstruction }
/*long取反*/
func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
