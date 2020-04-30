package comparisons

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

//branch if reference comparison succeeds
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}
