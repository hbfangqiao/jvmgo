package rtda

import "jvmgo/ch06/rtda/heap"

type Thread struct {
	pc		int//PC寄存器 当前方法是java方法则存放当前正在执行的java虚拟机指令的地址
	stack 		*Stack//JAVA虚拟机栈
}
/*新建栈结构体*/
func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024),//该栈最多可以容纳1024个帧
	}
}
func (self *Thread) PC() int {return self.pc}//getter
func (self *Thread) SetPC(pc int) {self.pc = pc}//setter

/*调用stack的push方法*/
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
/*调用stack的pop方法*/
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
/*返回当前帧*/
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self,method)
}