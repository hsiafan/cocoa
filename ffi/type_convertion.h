#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>

typedef struct {
    unsigned long len;
    const void* data;
} data;

typedef struct {
    unsigned long len;
    const void* data;
} array;

typedef struct {
    unsigned long len;
    const void* key_data;
    const void* value_data;
} dict;