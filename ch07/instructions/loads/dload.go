package loads

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

//从局部变量表中加载double类型的数据到操作数栈中
type DLOAD struct { base.Index8Instruction }

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame,uint(self.Index))
}

type DLOAD_0 struct { base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtda.Frame){
	_dload(frame,0)
}

type DLOAD_1 struct { base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtda.Frame){
	_dload(frame,1)
}
type DLOAD_2 struct { base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtda.Frame){
	_dload(frame,2)
}
type DLOAD_3 struct { base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtda.Frame){
	_dload(frame,3)
}

func _dload(frame *rtda.Frame,index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
