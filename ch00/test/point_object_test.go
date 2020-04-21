package main

import (
	"fmt"
	"testing"
)

func TestObject_M1(t *testing.T) {
	object := &Object{Name: "m"}
	fmt.Println("M1调用前：", object.Name)
	//当调用 t1.M1() 时相当于 M1(t1) ，实参和行参都是类型 T，可以接受。此时在M1()中的t只是t1的值拷贝，所以M1()的修改影响不到t1。
	object.M1()
	fmt.Println("M1调用后：", object.Name)

	fmt.Println("M2调用前：", object.Name)
	//当调用 t1.M2() => M2(t1)，这是将 T 类型传给了 *T 类型，go可能会取 t1 的地址传进去： M2(&t1)。所以 M2() 的修改可以影响 t1 。
	object.M2()
	fmt.Println("M2调用后：", object.Name)
}
