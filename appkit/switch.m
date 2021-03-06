#import <Appkit/Appkit.h>
#import "switch.h"

void* C_Switch_Alloc() {
    return [NSSwitch alloc];
}

void* C_NSSwitch_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result_ = [nSSwitch initWithFrame:frameRect];
    return result_;
}

void* C_NSSwitch_InitWithCoder(void* ptr, void* coder) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result_ = [nSSwitch initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSwitch_Init(void* ptr) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result_ = [nSSwitch init];
    return result_;
}

int C_NSSwitch_State(void* ptr) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSControlStateValue result_ = [nSSwitch state];
    return result_;
}

void C_NSSwitch_SetState(void* ptr, int value) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    [nSSwitch setState:value];
}
