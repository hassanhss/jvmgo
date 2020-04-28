package loads

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}





func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int,index int32) {
	if index <0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}