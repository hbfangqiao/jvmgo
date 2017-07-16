package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
type BIPUSH struct { val int8 } //Push byte
type SIPUSH struct { val int16 } //Push short

/*从操作数中获取一个byte型整数*/
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
/*将读取的byte型整数扩展成int型，推入栈顶*/
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
