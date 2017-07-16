package stack

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }
/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
弹出int、float等占用一个操作数栈位置的变量
*/
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }
/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
弹出double等占用两个操作数栈位置的变量
*/
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}