package rtda

type Frame struct {
	lower		*Frame//实现链表数据结构
	localVars 	LocalVars//局部变量表指针
	operandStack	*OperandStack//操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame{
	//局部变量表大小maxLocals与操作数栈深度maxStack是由编译器预先计算好的，存储在class文件method_info结构体的Code属性中
	return &Frame{
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