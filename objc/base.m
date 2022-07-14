#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>
#import <objc/runtime.h>
#include "_cgo_export.h"
#import <objc/runtime.h>
#import <Block.h>

void Run_WithAutoreleasePool(uintptr_t ptr) {
    @autoreleasepool {
        runTaskAndDeleteHandle(ptr);
    }
}

void* IMP_ImplementationWithBlock(void* bp) {
    return imp_implementationWithBlock((id)bp);
}

void* IMP_GetBlock(void* ptr) {
    return imp_getBlock((IMP)ptr);
}

bool IMP_RemoveBlock(void* ptr) {
    return imp_removeBlock((IMP)ptr);
}

void* Objc_GetClass(const char* name) {
    return objc_getClass(name);
}

void* Objc_AllocateClassPair(void* superClass, const char* name, size_t extraBytes) {
    return objc_allocateClassPair((Class)superClass, name, extraBytes);
}

void Objc_RegisterClassPair(void* class) {
    return objc_registerClassPair((Class)class);
}

void* Class_CreateInstance(void* cls, unsigned idxIvars) {
    return class_createInstance((Class)cls, idxIvars);
}

const char* Class_GetName(void *cls) {
    return class_getName((Class)cls);
}

void Class_SetVersion(void* cls, int version) {
    return class_setVersion((Class)cls, version);
}
int Class_GetVersion(void* cls) {
    return class_getVersion((Class)cls);
}

void* Class_GetSuperClass(void* cls) {
    return class_getSuperclass((Class)cls);
}

bool Class_RespondsToSelectorAddMethod(void* cls, void* sel) {
    return class_respondsToSelector((Class)cls, (SEL)sel);
}

bool Class_AddMethod(void* cls, void* sel, void* imp, const char* types) {
    return class_addMethod((Class)cls, (SEL)sel, (IMP)imp, types);
}

void* Class_ReplaceMethod(void* cls, void* sel, void* imp, const char* types) {
    return class_replaceMethod((Class)cls, (SEL)sel, (IMP)imp, types);
}

void* Class_GetMethodImplementation(void* cls, void* sel) {
    return class_getMethodImplementation((Class)cls, (SEL)sel);
}

void* Class_GetMethodImplementationStret(void* cls, void* sel) {
    return class_getMethodImplementation_stret((Class)cls, (SEL)sel);
}

void* Class_GetInstanceMethod(void* cls, void* sel) {
    return class_getInstanceMethod((Class)cls, (SEL)sel);
}

void* Class_GetClassMethod(void* cls, void* sel) {
    return class_getClassMethod((Class)cls, (SEL)sel);
}

void* Class_CopyMethodList(void* cls, unsigned int *outCount) {
    return class_copyMethodList((Class)cls, outCount);
}

void* Class_GetProperty(void* cls, const char *name) {
    return class_getProperty((Class)cls, name);
}

void* Class_CopyPropertyList(void* cls, unsigned int *outCount) {
    return class_copyPropertyList((Class)cls, outCount);
}

 bool Class_AddProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount) {
    return class_addProperty((Class)cls, name, (const objc_property_attribute_t *)attributes, attributeCount);
 }

 void Class_ReplaceProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount) {
    class_replaceProperty((Class)cls, name, (const objc_property_attribute_t *)attributes, attributeCount);
 }

bool Class_AddProtocol(void* cls, void* protocol) {
    return class_addProtocol((Class)cls, (Protocol*)protocol);
}

void* Objc_GetProtocol(const char* name) {
    return objc_getProtocol(name);
}

void* Objc_AllocateProtocol(const char* name) {
    return objc_allocateProtocol(name);
}

const char* Protocol_GetName(void *protocol) {
    return protocol_getName((Protocol*)protocol);
}

struct objc_method_description Protocol_GetMethodDescription(void* protocol, void* sel, bool required, bool instanceMethod) {
    return protocol_getMethodDescription((Protocol*)protocol, (SEL)sel, required, instanceMethod);
}

struct objc_method_description* Protocol_CopyMethodDescriptionList(void* protocol, bool required, bool instanceMethod, unsigned int *outCount) {
    return protocol_copyMethodDescriptionList((Protocol*)protocol, required, instanceMethod, outCount);
}

void* Protocol_CopyProtocolList(void* protocol, unsigned int *outCount) {
    return protocol_copyProtocolList((Protocol*)protocol, outCount);
}

void* Protocol_GetProperty(void* protocol, const char *name, bool required, bool isInstanceProperty) {
    return protocol_getProperty((Protocol*)protocol, name, required, isInstanceProperty);
}

void* Protocol_CopyPropertyList(void* protocol, unsigned int *outCount) {
    return protocol_copyPropertyList((Protocol*)protocol, outCount);
}

void* Method_GetName(void* method) {
    return method_getName((Method)method);
}

const char* Method_GetTypeEncoding(void* method) {
    return method_getTypeEncoding((Method)method);
}

void* Method_GetImplementation(void* method) {
    return method_getImplementation((Method)method);
}

void* Method_SetImplementation(void* method, void* imp) {
    return method_setImplementation((Method)method, (IMP)imp);
}


const char* Property_GetName(void* property) {
    return property_getName((objc_property_t)property);
}

const char* Property_GetAttributes(void* property) {
    return property_getAttributes((objc_property_t)property);
}

char* Property_CopyAttributeValue(void* property, const char* name) {
    return property_copyAttributeValue((objc_property_t)property, name);
}

objc_property_attribute_t * Property_CopyAttributeList(void* property, unsigned int *outCount) {
    return  property_copyAttributeList((objc_property_t)property, outCount);
}

void* block_copy(void *ptr) {
    return Block_copy(ptr);
}

void block_release(void *ptr) {
    Block_release(ptr);
}

void* Selector_SEL_RegisterName(const char* name) {
    return (void*)sel_registerName(name);
}

const char* Selector_SEL_GetName(void* ptr) {
    return sel_getName((SEL)ptr);
}

void Objc_SetAssociatedObject(void* ptr, const void* key, void* valuePtr, uintptr_t policy) {
    objc_setAssociatedObject((id)ptr, key, (id)valuePtr, policy);
}

void* Objc_GetAssociatedObject(void* ptr, const void* key) {
    return objc_getAssociatedObject((id)ptr, key);
}

void Objc_RemoveAssociatedObjects(void* ptr) {
    return objc_removeAssociatedObjects((id)ptr);
}