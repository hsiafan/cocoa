#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutGroup_Alloc();

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count);
void* C_NSCollectionLayoutGroup_VisualDescription(void* ptr);
Array C_NSCollectionLayoutGroup_Subitems(void* ptr);
void C_NSCollectionLayoutGroup_SetSupplementaryItems(void* ptr, Array value);
void* C_NSCollectionLayoutGroup_InterItemSpacing(void* ptr);
void C_NSCollectionLayoutGroup_SetInterItemSpacing(void* ptr, void* value);
