#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ImageRep_Alloc();

void* C_NSImageRep_Init(void* ptr);
void* C_NSImageRep_InitWithCoder(void* ptr, void* coder);
Array C_NSImageRep_ImageRepsWithContentsOfFile(void* filename);
Array C_NSImageRep_ImageRepsWithPasteboard(void* pasteboard);
Array C_NSImageRep_ImageRepsWithContentsOfURL(void* url);
void* C_NSImageRep_ImageRepWithContentsOfFile(void* filename);
void* C_NSImageRep_ImageRepWithPasteboard(void* pasteboard);
void* C_NSImageRep_ImageRepWithContentsOfURL(void* url);
bool C_NSImageRep_ImageRep_CanInitWithData(Array data);
bool C_NSImageRep_ImageRep_CanInitWithPasteboard(void* pasteboard);
bool C_NSImageRep_Draw(void* ptr);
bool C_NSImageRep_DrawAtPoint(void* ptr, CGPoint point);
bool C_NSImageRep_DrawInRect(void* ptr, CGRect rect);
Array C_NSImageRep_ImageRep_ImageTypes();
Array C_NSImageRep_ImageRep_ImageUnfilteredTypes();
CGSize C_NSImageRep_Size(void* ptr);
void C_NSImageRep_SetSize(void* ptr, CGSize value);
int C_NSImageRep_BitsPerSample(void* ptr);
void C_NSImageRep_SetBitsPerSample(void* ptr, int value);
void* C_NSImageRep_ColorSpaceName(void* ptr);
void C_NSImageRep_SetColorSpaceName(void* ptr, void* value);
bool C_NSImageRep_HasAlpha(void* ptr);
void C_NSImageRep_SetAlpha(void* ptr, bool value);
bool C_NSImageRep_IsOpaque(void* ptr);
void C_NSImageRep_SetOpaque(void* ptr, bool value);
int C_NSImageRep_PixelsHigh(void* ptr);
void C_NSImageRep_SetPixelsHigh(void* ptr, int value);
int C_NSImageRep_PixelsWide(void* ptr);
void C_NSImageRep_SetPixelsWide(void* ptr, int value);
int C_NSImageRep_LayoutDirection(void* ptr);
void C_NSImageRep_SetLayoutDirection(void* ptr, int value);
