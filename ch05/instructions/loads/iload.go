package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ILOAD struct {
	base.Index8Instruction
}

func _load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
