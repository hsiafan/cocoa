#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSOutlineViewDelegateAdaptor : NSObject <NSOutlineViewDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapOutlineViewDelegate(uintptr_t goID);
