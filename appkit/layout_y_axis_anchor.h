#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_LayoutYAxisAnchor_Alloc();

void* C_NSLayoutYAxisAnchor_Init(void* ptr);
void* C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor);
