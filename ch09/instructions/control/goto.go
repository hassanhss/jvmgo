package control

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

//branch control
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
