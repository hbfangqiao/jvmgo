package rtda

import "jvmgo/ch07/rtda/heap"

type Frame struct {
	lower		*Frame//实现链表数据结构
	localVars 	LocalVars//局部变量表指针
	operandStack	*OperandStack//操作数栈指针
	//实现跳转指令------------
	thread		*Thread
	nextPC		int
	//----------------------
	method 		*heap.Method
}

func newFrame(thread *Thread,method *heap.Method) *Frame{
	//局部变量表大小maxLocals与操作数栈深度maxStack是由编译器预先计算好的，存储在class文件method_info结构体的Code属性中
	return &Frame{
		thread: 	thread,
		method:		method,
		localVars: 	newLocalVars(method.MaxLocals()),
		operandStack:	newOperandStack(method.MaxStack()),
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

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}