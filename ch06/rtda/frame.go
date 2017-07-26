package rtda

import "jvmgo/ch06/rtda/heap"

type Frame struct {
	lower		*Frame//实现链表数据结构
	localVars 	LocalVars//局部变量表指针
	operandStack	*OperandStack//操作数栈指针
	//实现跳转指令------------
	thread		*Thread
	nextPC		int
	//----------------------
}

func newFrame(thread *Thread,maxLocals, maxStack uint) *Frame{
	//局部变量表大小maxLocals与操作数栈深度maxStack是由编译器预先计算好的，存储在class文件method_info结构体的Code属性中
	return &Frame{
		thread: 	thread,
		localVars: 	newLocalVars(maxLocals),
		operandStack:	newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}



func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}