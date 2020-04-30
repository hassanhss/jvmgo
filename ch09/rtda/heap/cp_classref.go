package heap

import "jvmgo/ch09/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(pool *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.pool = pool
	ref.className = classInfo.Name()
	return ref
}
