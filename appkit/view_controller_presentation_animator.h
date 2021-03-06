#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSViewControllerPresentationAnimatorAdaptor : NSObject <NSViewControllerPresentationAnimator>
@property (assign) uintptr_t goID;
@end

void* WrapViewControllerPresentationAnimator(uintptr_t goID);
