package base

import "jvmgo/ch06/rtda"

type Instruction interface {
	/*从字节码中提取操作数*/
	FetchOperands(reader *BytecodeReader)
	/*执行指令逻辑*/
	Execute(frame *rtda.Frame)
}

/*没有操作数的指令*/
type NoOperandsInstruction struct {}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader){}

/*跳转指令*/
type BranchInstruction struct {
	//跳转偏移量
	Offset int
}
/*从字节码中读取一个uint16整数，转成int后赋值给Offset字段*/
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*存储和加载类指令，根据索引存取局部变量表*/
type Index8Instruction struct {
	Index uint//局部变量表索引
}
/*从字节码中读取int8整数，转成uint后赋值给Index字段*/
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/*需要访问运行时常量池的指令，索引由两字节操作数给出*/
type Index16Instruction struct {
	Index uint
}
/*从字节码中读取uint16整数，转成uint后赋值给Index字段*/
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

