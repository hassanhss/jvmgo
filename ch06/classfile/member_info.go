package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	pool            ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	//创建membercount个MemberInfo
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, pool)
	}
	return members
}

func readMember(reader *ClassReader, pool ConstantPool) *MemberInfo {
	return &MemberInfo{
		pool:            pool,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, pool),
	}
}

func (self *MemberInfo) AccessFlag() uint16 {
	return self.accessFlag
}

func (self *MemberInfo) Name() string {
	return self.pool.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.pool.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
