#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorWell_Alloc();

void* C_NSColorWell_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSColorWell_InitWithCoder(void* ptr, void* coder);
void* C_NSColorWell_Init(void* ptr);
void C_NSColorWell_TakeColorFrom(void* ptr, void* sender);
void C_NSColorWell_Activate(void* ptr, bool exclusive);
void C_NSColorWell_Deactivate(void* ptr);
void C_NSColorWell_DrawWellInside(void* ptr, CGRect insideRect);
void* C_NSColorWell_Color(void* ptr);
void C_NSColorWell_SetColor(void* ptr, void* value);
bool C_NSColorWell_IsActive(void* ptr);
bool C_NSColorWell_IsBordered(void* ptr);
void C_NSColorWell_SetBordered(void* ptr, bool value);
