#import <Appkit/Appkit.h>
#import "text_storage.h"

void* C_TextStorage_Alloc() {
    return [NSTextStorage alloc];
}

void* C_NSTextStorage_InitWithString(void* ptr, void* str) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result_ = [nSTextStorage initWithString:(NSString*)str];
    return result_;
}

void* C_NSTextStorage_InitWithAttributedString(void* ptr, void* attrStr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result_ = [nSTextStorage initWithAttributedString:(NSAttributedString*)attrStr];
    return result_;
}

void* C_NSTextStorage_Init(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result_ = [nSTextStorage init];
    return result_;
}

void C_NSTextStorage_AddLayoutManager(void* ptr, void* aLayoutManager) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage addLayoutManager:(NSLayoutManager*)aLayoutManager];
}

void C_NSTextStorage_RemoveLayoutManager(void* ptr, void* aLayoutManager) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage removeLayoutManager:(NSLayoutManager*)aLayoutManager];
}

void C_NSTextStorage_Edited_Range_ChangeInLength(void* ptr, unsigned int editedMask, NSRange editedRange, int delta) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage edited:editedMask range:editedRange changeInLength:delta];
}

void C_NSTextStorage_ProcessEditing(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage processEditing];
}

void C_NSTextStorage_InvalidateAttributesInRange(void* ptr, NSRange _range) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage invalidateAttributesInRange:_range];
}

void C_NSTextStorage_EnsureAttributesAreFixedInRange(void* ptr, NSRange _range) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage ensureAttributesAreFixedInRange:_range];
}

void* C_NSTextStorage_Delegate(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    id result_ = [nSTextStorage delegate];
    return result_;
}

void C_NSTextStorage_SetDelegate(void* ptr, void* value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage setDelegate:(id)value];
}

Array C_NSTextStorage_LayoutManagers(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result_ = [nSTextStorage layoutManagers];
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

bool C_NSTextStorage_FixesAttributesLazily(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    BOOL result_ = [nSTextStorage fixesAttributesLazily];
    return result_;
}

unsigned int C_NSTextStorage_EditedMask(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorageEditActions result_ = [nSTextStorage editedMask];
    return result_;
}

NSRange C_NSTextStorage_EditedRange(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSRange result_ = [nSTextStorage editedRange];
    return result_;
}

int C_NSTextStorage_ChangeInLength(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSInteger result_ = [nSTextStorage changeInLength];
    return result_;
}

Array C_NSTextStorage_AttributeRuns(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result_ = [nSTextStorage attributeRuns];
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

void C_NSTextStorage_SetAttributeRuns(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setAttributeRuns:objcValue];
}

Array C_NSTextStorage_Paragraphs(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result_ = [nSTextStorage paragraphs];
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

void C_NSTextStorage_SetParagraphs(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setParagraphs:objcValue];
}

Array C_NSTextStorage_Words(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result_ = [nSTextStorage words];
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

void C_NSTextStorage_SetWords(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setWords:objcValue];
}

Array C_NSTextStorage_Characters(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result_ = [nSTextStorage characters];
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

void C_NSTextStorage_SetCharacters(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setCharacters:objcValue];
}

void* C_NSTextStorage_Font(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSFont* result_ = [nSTextStorage font];
    return result_;
}

void C_NSTextStorage_SetFont(void* ptr, void* value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage setFont:(NSFont*)value];
}

void* C_NSTextStorage_ForegroundColor(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSColor* result_ = [nSTextStorage foregroundColor];
    return result_;
}

void C_NSTextStorage_SetForegroundColor(void* ptr, void* value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage setForegroundColor:(NSColor*)value];
}
