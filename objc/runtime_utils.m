#import <stdint.h>
#import <objc/runtime.h>
#import <Foundation/NSObject.h>
#include "_cgo_export.h"


@interface DeallocListener : NSObject
@property (assign) uintptr_t goID;
@end

@implementation DeallocListener

- (void)dealloc {
    runTaskAndDeleteHandle([self goID]);
    [super dealloc];
}
@end

void* C_NewDeallocListener(uintptr_t id) {
    DeallocListener* listener = [[DeallocListener alloc] init];
	[listener setGoID:id];
	return listener;
}