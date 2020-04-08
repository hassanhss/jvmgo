package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
	//empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {

}
