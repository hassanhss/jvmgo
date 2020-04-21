package main

type Object struct {
	Name string
}

func (self Object) M1() {
	self.Name = "m1"
}

func (self *Object) M2() {
	self.Name = "m2"
}
