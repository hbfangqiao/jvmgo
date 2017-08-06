package references

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

// Get length of array
/*arraylength指令只需要一个操作数，即从操作数栈顶弹出的数组引用*/
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

/*Execute() 方法把数组长度推入操作数栈顶*/
func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
