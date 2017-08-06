package math

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type DDIV struct { base.NoOperandsInstruction }
/*double相除 榜眼除以状元*/
func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}


type FDIV struct { base.NoOperandsInstruction }
/*float相除*/
func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

type IDIV struct{ base.NoOperandsInstruction }
//int 相除
func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushInt(result)
}

type LDIV struct { base.NoOperandsInstruction }
/*Long相除*/
func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}

