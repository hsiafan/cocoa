#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_AffineTransform_Alloc();

void* C_NSAffineTransform_InitWithTransform(void* ptr, void* transform);
void* C_NSAffineTransform_AffineTransform_Transform();
void C_NSAffineTransform_RotateByDegrees(void* ptr, double angle);
void C_NSAffineTransform_RotateByRadians(void* ptr, double angle);
void C_NSAffineTransform_ScaleBy(void* ptr, double scale);
void C_NSAffineTransform_ScaleXBy_YBy(void* ptr, double scaleX, double scaleY);
void C_NSAffineTransform_TranslateXBy_YBy(void* ptr, double deltaX, double deltaY);
void C_NSAffineTransform_AppendTransform(void* ptr, void* transform);
void C_NSAffineTransform_PrependTransform(void* ptr, void* transform);
void C_NSAffineTransform_Invert(void* ptr);
CGPoint C_NSAffineTransform_TransformPoint(void* ptr, CGPoint aPoint);
CGSize C_NSAffineTransform_TransformSize(void* ptr, CGSize aSize);
void C_NSAffineTransform_Set(void* ptr);
void C_NSAffineTransform_Concat(void* ptr);
NSAffineTransformStruct C_NSAffineTransform_TransformStruct(void* ptr);
void C_NSAffineTransform_SetTransformStruct(void* ptr, NSAffineTransformStruct value);
