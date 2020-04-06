package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type MarkerAttribute struct {
}

type DepreatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}
