#import <CoreGraphics/CGWindow.h>

//CFArrayRef CGWindowListCreate(CGWindowListOption option, CGWindowID relativeToWindow);
void* WindowListCreate(uint32_t option, uint32_t relativeToWindow) {
    return (void*)CGWindowListCreate(option, relativeToWindow);
}

long ArrayCount(void* array) {
    return CFArrayGetCount((CFArrayRef)array);
}

void ArrayGetValues(void* array, long len, void* values) {
    return CFArrayGetValues((CFArrayRef)array, CFRangeMake(0, len), (const void**)values);
}

void* ArrayCreate(void *values, long numValues) {
    return (void *)CFArrayCreate(NULL, (const void **)values, numValues, NULL);
}

void* WindowListCreateImage(CGRect screenBounds, uint32_t listOption, uint32_t windowID, uint32_t imageOption) {
    return (void*)CGWindowListCreateImage(screenBounds, listOption, windowID, imageOption);
}

void* WindowListCreateImageFromArray(CGRect screenBounds, void* windowArray, uint32_t imageOption) {
    return (void*)CGWindowListCreateImageFromArray(screenBounds, (CFArrayRef)windowArray, imageOption);
}