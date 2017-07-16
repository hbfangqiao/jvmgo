package extended
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//goto_w指令和goto指令的唯一区别就是索引从2字节变成了4字节
type GOTO_W struct { offset int }

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame,self.offset)
}
