package math
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
/*iinc指令给局部变量表中的int变量增加常量值，局部变量表索 引和常量值都由指令的操作数提供。*/
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())//从字节码中读取操作数
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index,val)
}