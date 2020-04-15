package heap

type Object struct {
	class  *Class
	fields Slots
}

func NewObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
