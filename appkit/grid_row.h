#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_GridRow_Alloc();

void* C_NSGridRow_Init(void* ptr);
void* C_NSGridRow_CellAtIndex(void* ptr, int index);
void C_NSGridRow_MergeCellsInRange(void* ptr, NSRange _range);
int C_NSGridRow_NumberOfCells(void* ptr);
bool C_NSGridRow_IsHidden(void* ptr);
void C_NSGridRow_SetHidden(void* ptr, bool value);
double C_NSGridRow_TopPadding(void* ptr);
void C_NSGridRow_SetTopPadding(void* ptr, double value);
double C_NSGridRow_BottomPadding(void* ptr);
void C_NSGridRow_SetBottomPadding(void* ptr, double value);
double C_NSGridRow_Height(void* ptr);
void C_NSGridRow_SetHeight(void* ptr, double value);
int C_NSGridRow_RowAlignment(void* ptr);
void C_NSGridRow_SetRowAlignment(void* ptr, int value);
int C_NSGridRow_YPlacement(void* ptr);
void C_NSGridRow_SetYPlacement(void* ptr, int value);
void* C_NSGridRow_GridView(void* ptr);
