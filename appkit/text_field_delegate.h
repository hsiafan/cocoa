#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextFieldDelegateAdaptor : NSObject <NSTextFieldDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTextFieldDelegate(uintptr_t goID);
