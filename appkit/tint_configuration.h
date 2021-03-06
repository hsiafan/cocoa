#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TintConfiguration_Alloc();

void* C_NSTintConfiguration_Init(void* ptr);
void* C_NSTintConfiguration_TintConfigurationWithFixedColor(void* color);
void* C_NSTintConfiguration_TintConfigurationWithPreferredColor(void* color);
bool C_NSTintConfiguration_AdaptsToUserAccentColor(void* ptr);
void* C_NSTintConfiguration_DefaultTintConfiguration();
void* C_NSTintConfiguration_MonochromeTintConfiguration();
void* C_NSTintConfiguration_BaseTintColor(void* ptr);
void* C_NSTintConfiguration_EquivalentContentTintColor(void* ptr);
