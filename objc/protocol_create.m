#import <Foundation/NSInvocation.h>
#import <Foundation/NSMethodSignature.h>
#include "_cgo_export.h"


@interface ProtocolImplBase : NSObject
@property (assign) uintptr_t goID;
@end

@implementation ProtocolImplBase

-(void)forwardInvocation:(NSInvocation *)invocation {
    handleDelegateInvocation([self goID], [invocation selector], invocation);
}

- (BOOL)respondsToSelector:(SEL)aSelector {
	return respondsTo([self goID], aSelector);
}

- (void)dealloc {
	internalDeleteHandle([self goID]);
	[super dealloc];
}
@end


void* New_ProtocolImpl(void* class, uintptr_t goID) {
    ProtocolImplBase* o =(ProtocolImplBase*)[[(Class)class alloc] init];
    [o setGoID:goID];
    return o;
}

uintptr_t ProtocolImpl_GetGoID(void* ptr) {
    ProtocolImplBase* o =(ProtocolImplBase*)ptr;
    return [o goID];
}

