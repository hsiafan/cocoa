#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_AnimationContext_Alloc();

void* C_NSAnimationContext_Init(void* ptr);
void C_NSAnimationContext_AnimationContext_BeginGrouping();
void C_NSAnimationContext_AnimationContext_EndGrouping();
void* C_NSAnimationContext_AnimationContext_CurrentContext();
double C_NSAnimationContext_Duration(void* ptr);
void C_NSAnimationContext_SetDuration(void* ptr, double value);
bool C_NSAnimationContext_AllowsImplicitAnimation(void* ptr);
void C_NSAnimationContext_SetAllowsImplicitAnimation(void* ptr, bool value);
