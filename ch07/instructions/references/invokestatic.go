package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	pool := frame.Method().Class().ConstantPool()
	methodRef := pool.GetConstant(self.Index).(*heap.MethodRef)
	resoledMethod := methodRef.ResolvedMethod()
	if !resoledMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resoledMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame,resoledMethod)
}