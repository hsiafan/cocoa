#import <Appkit/Appkit.h>
#import "table_cell_view.h"

void* C_TableCellView_Alloc() {
    return [NSTableCellView alloc];
}

void* C_NSTableCellView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSTableCellView* result_ = [nSTableCellView initWithFrame:frameRect];
    return result_;
}

void* C_NSTableCellView_InitWithCoder(void* ptr, void* coder) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSTableCellView* result_ = [nSTableCellView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableCellView_Init(void* ptr) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSTableCellView* result_ = [nSTableCellView init];
    return result_;
}

void* C_NSTableCellView_ObjectValue(void* ptr) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    id result_ = [nSTableCellView objectValue];
    return result_;
}

void C_NSTableCellView_SetObjectValue(void* ptr, void* value) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    [nSTableCellView setObjectValue:(id)value];
}

int C_NSTableCellView_BackgroundStyle(void* ptr) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSBackgroundStyle result_ = [nSTableCellView backgroundStyle];
    return result_;
}

void C_NSTableCellView_SetBackgroundStyle(void* ptr, int value) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    [nSTableCellView setBackgroundStyle:value];
}

int C_NSTableCellView_RowSizeStyle(void* ptr) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSTableViewRowSizeStyle result_ = [nSTableCellView rowSizeStyle];
    return result_;
}

void C_NSTableCellView_SetRowSizeStyle(void* ptr, int value) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    [nSTableCellView setRowSizeStyle:value];
}

Array C_NSTableCellView_DraggingImageComponents(void* ptr) {
    NSTableCellView* nSTableCellView = (NSTableCellView*)ptr;
    NSArray* result_ = [nSTableCellView draggingImageComponents];
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
