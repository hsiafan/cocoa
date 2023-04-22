#import <objc/NSObject.h>
#include "_cgo_export.h"


@interface ProtocolImplBase : NSObject
@property (assign) uintptr_t goID;
@end

@implementation ProtocolImplBase

- (BOOL)respondsToSelector:(SEL)aSelector {
	return respondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteHandle([self goID]);
	[super dealloc];
}
@end


void* New_ProtocolImpl(void* class, uintptr_t goID) {
    ProtocolImplBase* o =(ProtocolImplBase*)[[(Class)class alloc] init];
    [o setGoID:goID];
    return o;
}

