package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type MULTI_ANEW_ARRAY struct {
	index 		uint16
	dimensions  uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {

}