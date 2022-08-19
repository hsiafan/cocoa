#import <CoreGraphics/CGDirectDisplay.h>
#import <CoreGraphics/CGDisplayConfiguration.h>

uint32_t MainDisplayID() {
    return CGMainDisplayID();
}

uint32_t GetOnlineDisplayList(uint32_t maxDisplays, void* onlineDisplays, uint32_t *displayCount) {
    return CGGetOnlineDisplayList(maxDisplays, (CGDirectDisplayID *)onlineDisplays, displayCount);
}

uint32_t GetActiveDisplayList(uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount) {
    return CGGetActiveDisplayList(maxDisplays, (CGDirectDisplayID *)activeDisplays, displayCount);
}

uint32_t GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount) {
    return CGGetDisplaysWithPoint(point, maxDisplays, (CGDirectDisplayID *)activeDisplays, displayCount);
}

uint32_t GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount) {
    return CGGetDisplaysWithRect(rect, maxDisplays, (CGDirectDisplayID *)activeDisplays, displayCount);
}

void* DisplayCreateImage(uint32_t displayID) {
    return CGDisplayCreateImage(displayID);
}

void* DisplayCreateImageForRect(uint32_t displayID, CGRect rect) {
    return CGDisplayCreateImageForRect(displayID, rect);
}

CGSize DisplayScreenSize(uint32_t displayID) {
    return CGDisplayScreenSize(displayID);
}

bool DisplayIsActive(uint32_t displayID) {
    return CGDisplayIsActive(displayID);
}

bool DisplayIsAsleep(uint32_t displayID) {
    return CGDisplayIsAsleep(displayID);
}

bool DisplayIsBuiltin(uint32_t displayID) {
    return CGDisplayIsBuiltin(displayID);
}

bool DisplayIsMain(uint32_t displayID) {
    return CGDisplayIsMain(displayID);
}

bool DisplayIsOnline(uint32_t displayID) {
    return CGDisplayIsOnline(displayID);
}

size_t DisplayPixelsHigh(uint32_t displayID) {
    return CGDisplayPixelsHigh(displayID);
}

size_t DisplayPixelsWide(uint32_t displayID) {
    return DisplayPixelsWide(displayID);
}

CGRect DisplayBounds(uint32_t displayID) {
    return CGDisplayBounds(displayID);
}