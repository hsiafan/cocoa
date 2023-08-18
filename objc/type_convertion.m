#import <stdint.h>
#import <stdlib.h>
#import <stdbool.h>
#import <CoreFoundation/CFString.h>
#import <CoreFoundation/CFData.h>
#import <CoreFoundation/CFArray.h>
#import <CoreFoundation/CFDictionary.h>


void* to_ns_string(void* str, unsigned long len) {
    return (void*)CFStringCreateWithBytes(kCFAllocatorDefault, (const UInt8 *)str, len, kCFStringEncodingUTF8, false);
}

const char* to_c_string(void* ptr, bool* shouldFree) {
    const char * cstr = CFStringGetCStringPtr(ptr, kCFStringEncodingUTF8);
    if (cstr != NULL) {
        return cstr;
    }
    CFIndex length = CFStringGetLength(ptr);
    CFIndex maxSize = CFStringGetMaximumSizeForEncoding(length, kCFStringEncodingUTF8) + 1;
    char *buffer = (char *)malloc(maxSize);
    if (CFStringGetCString(ptr, buffer, maxSize, kCFStringEncodingUTF8)) {
        *shouldFree = true;
        return buffer;
    }
    free(buffer); // If we failed
    return NULL;
}

void* to_ns_data(void* data, unsigned long len) {
    return (void *)CFDataCreate(kCFAllocatorDefault, (const UInt8 *)data, len);
}

void* to_c_bytes(void* ptr, unsigned long *len) {
    *len = CFDataGetLength(ptr);
    const UInt8 * bytesPtr = CFDataGetBytePtr(ptr);
    return (void*)bytesPtr;
}

void* to_ns_array(void** items, unsigned long len) {
    return (void*)CFArrayCreate(kCFAllocatorDefault, (const void**)items, len, &kCFTypeArrayCallBacks);
}

unsigned long ns_array_len(void* ptr) {
    return CFArrayGetCount(ptr);
}

void ns_array_get(void* ptr, const void** item, unsigned long len) {
    CFArrayGetValues(ptr, CFRangeMake(0, len), item);
}


void* to_ns_dict(const void** keys, const void** values, unsigned long size) {
    return (void*)CFDictionaryCreate(kCFAllocatorDefault, keys, values, size, &kCFTypeDictionaryKeyCallBacks, &kCFTypeDictionaryValueCallBacks);
}

unsigned long ns_dict_size(void* ptr) {
    return CFDictionaryGetCount(ptr);
}

void ns_dict_get(void* ptr, const void** keys, const void** values) {
    CFDictionaryGetKeysAndValues(ptr, keys, values);
}