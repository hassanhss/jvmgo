package stack

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

// Pop the top operand stack value
type POP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
