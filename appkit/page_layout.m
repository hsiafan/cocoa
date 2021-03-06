#import <Appkit/Appkit.h>
#import "page_layout.h"

void* C_PageLayout_Alloc() {
    return [NSPageLayout alloc];
}

void* C_NSPageLayout_Init(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPageLayout* result_ = [nSPageLayout init];
    return result_;
}

void* C_NSPageLayout_PageLayout_() {
    NSPageLayout* result_ = [NSPageLayout pageLayout];
    return result_;
}

int C_NSPageLayout_RunModal(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSInteger result_ = [nSPageLayout runModal];
    return result_;
}

int C_NSPageLayout_RunModalWithPrintInfo(void* ptr, void* printInfo) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSInteger result_ = [nSPageLayout runModalWithPrintInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void C_NSPageLayout_AddAccessoryController(void* ptr, void* accessoryController) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    [nSPageLayout addAccessoryController:(NSViewController*)accessoryController];
}

void C_NSPageLayout_RemoveAccessoryController(void* ptr, void* accessoryController) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    [nSPageLayout removeAccessoryController:(NSViewController*)accessoryController];
}

Array C_NSPageLayout_AccessoryControllers(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSArray* result_ = [nSPageLayout accessoryControllers];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

void* C_NSPageLayout_PrintInfo(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPrintInfo* result_ = [nSPageLayout printInfo];
    return result_;
}
