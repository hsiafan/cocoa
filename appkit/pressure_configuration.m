#import <Appkit/Appkit.h>
#import "pressure_configuration.h"

void* C_PressureConfiguration_Alloc() {
    return [NSPressureConfiguration alloc];
}

void* C_NSPressureConfiguration_InitWithPressureBehavior(void* ptr, int pressureBehavior) {
    NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
    NSPressureConfiguration* result_ = [nSPressureConfiguration initWithPressureBehavior:pressureBehavior];
    return result_;
}

void* C_NSPressureConfiguration_Init(void* ptr) {
    NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
    NSPressureConfiguration* result_ = [nSPressureConfiguration init];
    return result_;
}

void C_NSPressureConfiguration_Set(void* ptr) {
    NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
    [nSPressureConfiguration set];
}

int C_NSPressureConfiguration_PressureBehavior(void* ptr) {
    NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
    NSPressureBehavior result_ = [nSPressureConfiguration pressureBehavior];
    return result_;
}
