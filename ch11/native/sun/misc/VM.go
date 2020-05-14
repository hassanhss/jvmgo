package misc

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/native"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) { // hack: just make VM.savedProps nonempty
	classLoader := frame.Method().Class().Loader()
	jISysClass := classLoader.LoadClass("java/lang/System")
	//initSysClass := jISysClass.Get
}
