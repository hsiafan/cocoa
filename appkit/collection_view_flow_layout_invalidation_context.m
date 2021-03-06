#import <Appkit/Appkit.h>
#import "collection_view_flow_layout_invalidation_context.h"

void* C_CollectionViewFlowLayoutInvalidationContext_Alloc() {
    return [NSCollectionViewFlowLayoutInvalidationContext alloc];
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_Init(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [nSCollectionViewFlowLayoutInvalidationContext init];
    return result_;
}

bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayoutInvalidationContext invalidateFlowLayoutAttributes];
    return result_;
}

void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(void* ptr, bool value) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    [nSCollectionViewFlowLayoutInvalidationContext setInvalidateFlowLayoutAttributes:value];
}

bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayoutInvalidationContext invalidateFlowLayoutDelegateMetrics];
    return result_;
}

void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(void* ptr, bool value) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    [nSCollectionViewFlowLayoutInvalidationContext setInvalidateFlowLayoutDelegateMetrics:value];
}
