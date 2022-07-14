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

typedef struct {
    void *isa;
    int flags;
    int reserved;
    void *invoke;
} Block;

enum {
    // Set to true on blocks that have captures (and thus are not true
    // global blocks) but are known not to escape for various other
    // reasons. For backward compatibility with old runtimes, whenever
    // BLOCK_IS_NOESCAPE is set, BLOCK_IS_GLOBAL is set too. Copying a
    // non-escaping block returns the original block and releasing such a
    // block is a no-op, which is exactly how global blocks are handled.
    BLOCK_IS_NOESCAPE      =  (1 << 23),

    BLOCK_HAS_COPY_DISPOSE =  (1 << 25),
    BLOCK_HAS_CTOR =          (1 << 26), // helpers have C++ code
    BLOCK_IS_GLOBAL =         (1 << 28),
    BLOCK_HAS_STRET =         (1 << 29), // IFF BLOCK_HAS_SIGNATURE
    BLOCK_HAS_SIGNATURE =     (1 << 30),
};
