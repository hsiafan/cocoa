#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSWindowDelegateAdaptor : NSObject <NSWindowDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapWindowDelegate(uintptr_t goID);
