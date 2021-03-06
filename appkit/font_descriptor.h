#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_FontDescriptor_Alloc();

void* C_NSFontDescriptor_FontDescriptorWithDesign(void* ptr, void* design);
void* C_NSFontDescriptor_Init(void* ptr);
void* C_NSFontDescriptor_FontDescriptorWithName_Matrix(void* fontName, void* matrix);
void* C_NSFontDescriptor_FontDescriptorWithName_Size(void* fontName, double size);
void* C_NSFontDescriptor_FontDescriptorWithFace(void* ptr, void* newFace);
void* C_NSFontDescriptor_FontDescriptorWithFamily(void* ptr, void* newFamily);
void* C_NSFontDescriptor_FontDescriptorWithMatrix(void* ptr, void* matrix);
void* C_NSFontDescriptor_FontDescriptorWithSize(void* ptr, double newPointSize);
void* C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(void* ptr, uint32_t symbolicTraits);
Array C_NSFontDescriptor_MatchingFontDescriptorsWithMandatoryKeys(void* ptr, void* mandatoryKeys);
void* C_NSFontDescriptor_MatchingFontDescriptorWithMandatoryKeys(void* ptr, void* mandatoryKeys);
void* C_NSFontDescriptor_ObjectForKey(void* ptr, void* attribute);
void* C_NSFontDescriptor_Matrix(void* ptr);
double C_NSFontDescriptor_PointSize(void* ptr);
void* C_NSFontDescriptor_PostscriptName(void* ptr);
uint32_t C_NSFontDescriptor_SymbolicTraits(void* ptr);
bool C_NSFontDescriptor_RequiresFontAssetRequest(void* ptr);
