package control
import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type GOTO struct{ base.BranchInstruction }
/*goto指令进行无条件跳转*/
func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame,self.Offset)
}