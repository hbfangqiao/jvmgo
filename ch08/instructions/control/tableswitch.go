package control

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/*形如
int chooseNear(int i) {
switch (i) {
	case 0:  return  0;
	case 1:  return  1;
	case 2:  return  2;
	default: return -1;
	}
}
被编译成tableswitch指令
*/
type TABLE_SWITCH struct {
	defaultOffset int32//默认情况下执行跳转所需要的字节码偏移量
	low           int32//low和high记录case的取值范围
	high          int32
	jumpOffsets   []int32//存放high - low + 1个int值 对应各种case情况下，执行跳转所需的字节码偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}