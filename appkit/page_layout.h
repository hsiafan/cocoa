#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PageLayout_Alloc();

void* C_NSPageLayout_Init(void* ptr);
void* C_NSPageLayout_PageLayout_();
int C_NSPageLayout_RunModal(void* ptr);
int C_NSPageLayout_RunModalWithPrintInfo(void* ptr, void* printInfo);
void C_NSPageLayout_AddAccessoryController(void* ptr, void* accessoryController);
void C_NSPageLayout_RemoveAccessoryController(void* ptr, void* accessoryController);
Array C_NSPageLayout_AccessoryControllers(void* ptr);
void* C_NSPageLayout_PrintInfo(void* ptr);
