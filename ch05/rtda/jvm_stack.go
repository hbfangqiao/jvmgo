package rtda
type Stack struct {
	maxSize		uint//最多可以容纳的帧数
	size 		uint//栈的当前大小
	_top 		*Frame//栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
/*把帧推入栈顶*/
func (self *Stack) push(frame *Frame){
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		//将推入帧的lower指向原先的栈顶帧
		frame.lower = self._top
	}
	//将栈顶指针指向新推入的帧
	self._top = frame
	self.size++
}
/*把栈顶帧弹出*/
func (self *Stack) pop() *Frame {
	if self._top == nil{
		panic("jvm stack is empty")
	}
	top := self._top
	//将栈顶指向指向原本栈顶帧的下一个链表（也是一个帧）
	self._top = top.lower
	//将原本栈顶帧的lower指向变为nil
	top.lower = nil
	self.size--
	return top
}
/*只返回栈顶帧，并不弹出*/
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
