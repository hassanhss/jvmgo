package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (self *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handlerUncaughtException(thread,ex)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread,ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() -1
		handelrPC := frame.Method().FindExceptionHandler(ex.Class(),pc)
		if handelrPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handelrPC)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}


func handlerUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	
}