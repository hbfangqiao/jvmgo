package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type ILOAD struct { base.Index8Instruction }
type ILOAD_0 struct { base.NoOperandsInstruction }
type ILOAD_1 struct { base.NoOperandsInstruction }
type ILOAD_2 struct { base.NoOperandsInstruction }
type ILOAD_3 struct { base.NoOperandsInstruction }

/*取frame的中的局部变量表的数，推入操作数栈中*/
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
/*1.从字节码中获取int8操作数
  2.以操作数index取出Slot[index]
  3.推入操作数栈*/
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame,uint(self.Index))
}
/*将本地变量表中的第一个变量，推入操作数栈中*/
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame,0)
}
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame,1)
}
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame,2)
}
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame,3)
}

