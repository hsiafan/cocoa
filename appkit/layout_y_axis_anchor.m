#import <Appkit/Appkit.h>
#import "layout_y_axis_anchor.h"

void* C_LayoutYAxisAnchor_Alloc() {
    return [NSLayoutYAxisAnchor alloc];
}

void* C_NSLayoutYAxisAnchor_Init(void* ptr) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutYAxisAnchor init];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintLessThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutDimension* result_ = [nSLayoutYAxisAnchor anchorWithOffsetToAnchor:(NSLayoutYAxisAnchor*)otherAnchor];
    return result_;
}
