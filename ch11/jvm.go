package main

import (
	"fmt"
	"jvmgo/ch11/classpath"
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
	"strings"
)

type JVM struct {
	cmd 		*Cmd
	classloader *heap.ClassLoader
	mainThread 	*rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	classloader := heap.NewClassLoader(cp,cmd.verboseClassFlag)
	return &JVM{
		cmd: cmd,
		classloader: classloader,
		mainThread: rtda.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

func (self *JVM) initVM() {
	vmClass := self.classloader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread,self.cmd.verboseClassFlag)
}

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.class,".","/",-1)
	mainClass := self.classloader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
		return
	}
	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread,self.cmd.verboseClassFlag)
}

func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classloader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i,arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classloader,arg)
	}
	return argsArr
}