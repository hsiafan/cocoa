#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_IndexPath_Alloc();

void* C_NSIndexPath_InitWithIndex(void* ptr, unsigned int index);
void* C_NSIndexPath_Init(void* ptr);
void* C_NSIndexPath_IndexPathWithIndex(unsigned int index);
void* C_NSIndexPath_IndexPathForItem_InSection(int item, int section);
void* C_NSIndexPath_IndexPathByAddingIndex(void* ptr, unsigned int index);
void* C_NSIndexPath_IndexPathByRemovingLastIndex(void* ptr);
int C_NSIndexPath_Compare(void* ptr, void* otherObject);
unsigned int C_NSIndexPath_IndexAtPosition(void* ptr, unsigned int position);
int C_NSIndexPath_Section(void* ptr);
int C_NSIndexPath_Item(void* ptr);
unsigned int C_NSIndexPath_Length(void* ptr);
