package rtda

type Frame struct {
	lower        *Frame // stack is implemented as linked list
	localVar     LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVar:     newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVar
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
