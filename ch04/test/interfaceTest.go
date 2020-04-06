package test

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Walk(step int)
	Walk1(step int)
}

type People struct {
	Name     string
	CurrStep int
}

func (c People) PeInfo() {
	fmt.Println("people name:", c.Name, "currstep", c.CurrStep)
}

func (c *People) SetName(name string) {
	c.Name = name
	fmt.Println("people new name:", name)
}

func (c People) Walk(step int) {
	c.PeInfo()
	c.CurrStep += step
	fmt.Println("walk", step)
}

func (c *People) Walk1(step int) {
	c.PeInfo()
	c.CurrStep += step
	fmt.Println("walk", step)
}

func PrintMethodSet(i interface{}) {
	rt := reflect.TypeOf(i)
	fmt.Println("type is ", rt)
	for i, n := 0, rt.NumMethod(); i < n; i++ {
		fmt.Println("it has method:", rt.Method(i).Name)
	}
}

func LetWalk(a Animal, s int) {
	a.Walk(s)
}

func LetWalk1(p People, s int) {
	p.Walk(s)
}

func main() {
	//https://www.jianshu.com/p/e644efb9515f
	var p = People{Name: "hassan"}
	PrintMethodSet(p)

	//下面两句等效
	p.PeInfo()
	/*
	   指针调用值方法，指针包含所有方法集，指针调用值方法时，会通过指针创建值副本，然后再调用值方法;
	   效果等同于：
	   q := *(&p)
	   q.PeInfo()
	*/
	(&p).PeInfo()
	People.PeInfo(p)
	// People.PeInfo(&p) //语句错误，这里需要显示传递指针
	//下面三句等效
	p.SetName("xiaoming") //p是值类型，不包含指针方法，但是go 隐式调用了(&p).SetName("jiang")语法糖，所以可以调用
	(&p).SetName("xiaojiang")
	(*People).SetName(&p, "liu")
}
