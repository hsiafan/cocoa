#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Image_Alloc();

void* C_NSImage_InitByReferencingFile(void* ptr, void* fileName);
void* C_NSImage_InitByReferencingURL(void* ptr, void* url);
void* C_NSImage_InitWithContentsOfFile(void* ptr, void* fileName);
void* C_NSImage_InitWithContentsOfURL(void* ptr, void* url);
void* C_NSImage_InitWithData(void* ptr, Array data);
void* C_NSImage_InitWithDataIgnoringOrientation(void* ptr, Array data);
void* C_NSImage_InitWithPasteboard(void* ptr, void* pasteboard);
void* C_NSImage_InitWithCoder(void* ptr, void* coder);
void* C_NSImage_InitWithSize(void* ptr, CGSize size);
void* C_NSImage_Init(void* ptr);
void* C_NSImage_ImageNamed(void* name);
bool C_NSImage_SetName(void* ptr, void* _string);
void* C_NSImage_Name(void* ptr);
void* C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(void* symbolName, void* description);
void* C_NSImage_ImageWithSymbolConfiguration(void* ptr, void* configuration);
bool C_NSImage_Image_CanInitWithPasteboard(void* pasteboard);
void C_NSImage_AddRepresentation(void* ptr, void* imageRep);
void C_NSImage_AddRepresentations(void* ptr, Array imageReps);
void C_NSImage_RemoveRepresentation(void* ptr, void* imageRep);
void C_NSImage_DrawInRect(void* ptr, CGRect rect);
void C_NSImage_DrawAtPoint_FromRect_Operation_Fraction(void* ptr, CGPoint point, CGRect fromRect, unsigned int op, double delta);
void C_NSImage_DrawInRect_FromRect_Operation_Fraction(void* ptr, CGRect rect, CGRect fromRect, unsigned int op, double delta);
bool C_NSImage_DrawRepresentation_InRect(void* ptr, void* imageRep, CGRect rect);
void C_NSImage_LockFocus(void* ptr);
void C_NSImage_LockFocusFlipped(void* ptr, bool flipped);
void C_NSImage_UnlockFocus(void* ptr);
void C_NSImage_Recache(void* ptr);
Array C_NSImage_TIFFRepresentationUsingCompression_Factor(void* ptr, unsigned int comp, float factor);
void C_NSImage_CancelIncrementalLoad(void* ptr);
void* C_NSImage_LayerContentsForContentsScale(void* ptr, double layerContentsScale);
double C_NSImage_RecommendedLayerContentsScale(void* ptr, double preferredContentsScale);
void* C_NSImage_Delegate(void* ptr);
void C_NSImage_SetDelegate(void* ptr, void* value);
CGSize C_NSImage_Size(void* ptr);
void C_NSImage_SetSize(void* ptr, CGSize value);
bool C_NSImage_IsTemplate(void* ptr);
void C_NSImage_SetTemplate(void* ptr, bool value);
Array C_NSImage_ImageTypes();
Array C_NSImage_ImageUnfilteredTypes();
Array C_NSImage_Representations(void* ptr);
bool C_NSImage_PrefersColorMatch(void* ptr);
void C_NSImage_SetPrefersColorMatch(void* ptr, bool value);
bool C_NSImage_UsesEPSOnResolutionMismatch(void* ptr);
void C_NSImage_SetUsesEPSOnResolutionMismatch(void* ptr, bool value);
bool C_NSImage_MatchesOnMultipleResolution(void* ptr);
void C_NSImage_SetMatchesOnMultipleResolution(void* ptr, bool value);
bool C_NSImage_IsValid(void* ptr);
void* C_NSImage_BackgroundColor(void* ptr);
void C_NSImage_SetBackgroundColor(void* ptr, void* value);
NSEdgeInsets C_NSImage_CapInsets(void* ptr);
void C_NSImage_SetCapInsets(void* ptr, NSEdgeInsets value);
int C_NSImage_ResizingMode(void* ptr);
void C_NSImage_SetResizingMode(void* ptr, int value);
CGRect C_NSImage_AlignmentRect(void* ptr);
void C_NSImage_SetAlignmentRect(void* ptr, CGRect value);
unsigned int C_NSImage_CacheMode(void* ptr);
void C_NSImage_SetCacheMode(void* ptr, unsigned int value);
Array C_NSImage_TIFFRepresentation(void* ptr);
void* C_NSImage_AccessibilityDescription(void* ptr);
void C_NSImage_SetAccessibilityDescription(void* ptr, void* value);
bool C_NSImage_MatchesOnlyOnBestFittingAxis(void* ptr);
void C_NSImage_SetMatchesOnlyOnBestFittingAxis(void* ptr, bool value);
