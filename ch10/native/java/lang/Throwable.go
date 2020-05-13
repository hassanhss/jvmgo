package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
)

func init() {
	native.Register("java/lang/Throwable","fillInStackTrace","(I)Ljava/lang/Throwable;",fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {

}
