#import <stdint.h>
#import <objc/message.h>
#import <objc/objc-runtime.h>
#import <Foundation/NSInvocation.h>
#import <Foundation/NSMethodSignature.h>
#import <Foundation/Foundation.h>
#include "_cgo_export.h"

id __NSMakeSpecialForwardingCaptureBlock(const char *signature, void (^handler)(NSInvocation *inv));

const char *_Block_signature(void *);

@interface BlockMonitor : NSObject
@property (assign) uintptr_t goID;
@end

@implementation BlockMonitor

- (void)dealloc {
	internalDeleteHandle([self goID]);
	[super dealloc];
}
@end

void* wrap_block(uintptr_t goID, const char *typeEncoding) {
    BlockMonitor* m = [[[BlockMonitor alloc] init] autorelease];
    [m setGoID:goID];
    id proxy = __NSMakeSpecialForwardingCaptureBlock(typeEncoding, ^(NSInvocation *invocation) {
        handleBlockInvocation(m.goID, invocation);
    });
    [proxy autorelease];
    return proxy;
}

void callBlock(void *bp, int argc, uintptr_t argsPtr, void* ret) {
    id block = (id)bp;
    void** args = (void**)argsPtr;

    NSMethodSignature* methodSignature = [NSMethodSignature signatureWithObjCTypes:_Block_signature(block)];
    NSInvocation* invocation = [NSInvocation invocationWithMethodSignature:methodSignature];
    [invocation setTarget:block];
    for (int i = 0; i < argc; i++) {
        [invocation setArgument:args[i] atIndex:i+1];
    }
    [invocation invoke];
    if (ret != NULL) {
        [invocation getReturnValue:ret];
    }
}