#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_DraggingImageComponent_Alloc();

void* C_NSDraggingImageComponent_InitWithKey(void* ptr, void* key);
void* C_NSDraggingImageComponent_DraggingImageComponentWithKey(void* key);
void* C_NSDraggingImageComponent_Key(void* ptr);
void C_NSDraggingImageComponent_SetKey(void* ptr, void* value);
void* C_NSDraggingImageComponent_Contents(void* ptr);
void C_NSDraggingImageComponent_SetContents(void* ptr, void* value);
CGRect C_NSDraggingImageComponent_Frame(void* ptr);
void C_NSDraggingImageComponent_SetFrame(void* ptr, CGRect value);
