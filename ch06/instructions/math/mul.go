package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

//multiply double
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

//multiply float
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

//multiply int
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 * v1
	stack.PushInt(result)
}

//multiple Long
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 * v1
	stack.PushLong(result)
}
