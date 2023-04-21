package objc

// #import <stdlib.h>
// #import <stdint.h>
// void* MethodSignature_WithObjCTypes(const char* typeCodes);
// void Invocation_GetArguments(void* ptr, int begin, int argc, uintptr_t argsPtr);
// void Invocation_SetReturnValue(void* ptr, void* loc);
import "C"
import (
	"unsafe"
)

type invocation struct {
	Object
}

func makeInvocation(ptr unsafe.Pointer) invocation {
	return invocation{
		Object: MakeObject(ptr),
	}
}

func (i invocation) getArguments(begin int, args []unsafe.Pointer) {
	C.Invocation_GetArguments(i.Ptr(), C.int(begin), C.int(len(args)), toUIntptr(&args[0]))
}

func (i invocation) setReturnValue(rref unsafe.Pointer) {
	C.Invocation_SetReturnValue(i.Ptr(), rref)
}
