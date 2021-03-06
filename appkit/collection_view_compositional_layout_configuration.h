#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewCompositionalLayoutConfiguration_Alloc();

void* C_NSCollectionViewCompositionalLayoutConfiguration_Init(void* ptr);
int C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(void* ptr, int value);
double C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(void* ptr, double value);
Array C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(void* ptr, Array value);
