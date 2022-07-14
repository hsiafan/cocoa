#import <objc/message.h>
#import <objc/objc-runtime.h>
#import <Foundation/NSInvocation.h>
#import <Foundation/NSMethodSignature.h>

void callMethod(void *op, void* sp, int argc, uintptr_t argsPtr, void* ret) {
    id o = (id)op;
    SEL sel = (SEL)sp;
    void** args = (void**)argsPtr;

    NSMethodSignature* methodSignature = [o methodSignatureForSelector:sel];
    if (methodSignature == NULL) {
        methodSignature = [[o class] instanceMethodSignatureForSelector:sel];
    }
    NSInvocation* invocation = [NSInvocation invocationWithMethodSignature:methodSignature];
    [invocation setTarget:o];
    [invocation setSelector:sel];
    for (int i = 0; i < argc; i++) {
        [invocation setArgument:args[i] atIndex:i+2];
    }
    [invocation invoke];
    if (ret != NULL) {
        [invocation getReturnValue:ret];
    }
}
