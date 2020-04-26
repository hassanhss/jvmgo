package heap

import "jvmgo/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(pool *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.pool = pool
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	d := self.pool.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChange")
	}
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if method.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupMethod(c *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(c, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(c.interfaces, name, descriptor)
	}
	return method
}
