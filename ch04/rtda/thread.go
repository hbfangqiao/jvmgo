package rtda
type Thread struct {
	pc		int//PC寄存器
	stack 		*Stack//JAVA虚拟机栈
}
/*新建栈结构体*/
func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024),
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
