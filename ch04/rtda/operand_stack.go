package rtda

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxLock uint) *OperandStack {
	if maxLock > 0 {
		return &OperandStack{
			slots: make([]Slot, maxLock),
		}
	}
	return nil
}
