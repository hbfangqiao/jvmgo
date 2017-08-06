package extended
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type IFNULL struct{ base.BranchInstruction }
/*引用是空则跳转*/
func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
type IFNONNULL struct{ base.BranchInstruction }
//引用非空则跳转
func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}