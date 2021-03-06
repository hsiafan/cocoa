#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSBrowserDelegateAdaptor : NSObject <NSBrowserDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapBrowserDelegate(uintptr_t goID);
