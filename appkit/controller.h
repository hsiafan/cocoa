#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Controller_Alloc();

void* C_NSController_Init(void* ptr);
void* C_NSController_InitWithCoder(void* ptr, void* coder);
void C_NSController_ObjectDidBeginEditing(void* ptr, void* editor);
void C_NSController_ObjectDidEndEditing(void* ptr, void* editor);
bool C_NSController_CommitEditing(void* ptr);
void C_NSController_DiscardEditing(void* ptr);
bool C_NSController_IsEditing(void* ptr);
