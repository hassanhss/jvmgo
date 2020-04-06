package rtda

type Frame struct {
	lower        *Frame
	localVar     LocalVars
	operandStack *OperandStack
}
