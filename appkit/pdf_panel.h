#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PDFPanel_Alloc();

void* C_NSPDFPanel_Init(void* ptr);
void* C_NSPDFPanel_PDFPanel_Panel();
void* C_NSPDFPanel_AccessoryController(void* ptr);
void C_NSPDFPanel_SetAccessoryController(void* ptr, void* value);
int C_NSPDFPanel_Options(void* ptr);
void C_NSPDFPanel_SetOptions(void* ptr, int value);
void* C_NSPDFPanel_DefaultFileName(void* ptr);
void C_NSPDFPanel_SetDefaultFileName(void* ptr, void* value);
