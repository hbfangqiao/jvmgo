package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type ISHL struct { base.NoOperandsInstruction }
/*实现Int <<*/
func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()//v2指出要移位
	v1 := stack.PopInt()//要操作的数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)


}

type ISHR struct { base.NoOperandsInstruction }
/*实现 Int >>*/
func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

type IUSHR struct { base.NoOperandsInstruction }
/*实现java中的Int >>>运算，无符号位移高位补0*/
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	//先把v1转成无符号整数，位移操作之后再转回有符号整数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

type LSHL struct { base.NoOperandsInstruction }
/*实现java中Long << 运算*/
func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint(v2) & 0x3f
	result  := v1 << s
	stack.PushLong(result)
}

type LSHR struct { base.NoOperandsInstruction }
/*Long的>> 运算*/
func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f//取前6位
	result := v1 >> s
	stack.PushLong(result)
}
type LUSHR struct { base.NoOperandsInstruction }
/*Long 的逻辑右移*/
func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint(v1) >> s)
	stack.PushLong(result)
}