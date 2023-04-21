#import <Foundation/NSMethodSignature.h>
#import <Foundation/NSInvocation.h>

void* MethodSignature_WithObjCTypes(const char* typeCodes) {
    return [NSMethodSignature signatureWithObjCTypes:typeCodes];
}

void Invocation_GetArguments(void* ptr, int begin, int argc, uintptr_t argsPtr) {
    void** args = (void**)argsPtr;
    for (int i = 0; i < argc; i++) {
        [(NSInvocation *)ptr getArgument:args[i] atIndex:i+begin];
    }
}

void Invocation_SetReturnValue(void* ptr, void* loc) {
    [(NSInvocation *)ptr setReturnValue:loc];
}
