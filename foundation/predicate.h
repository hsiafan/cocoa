#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Predicate_Alloc();

void* C_NSPredicate_Init(void* ptr);
void* C_NSPredicate_PredicateWithFormat_ArgumentArray(void* predicateFormat, Array arguments);
void* C_NSPredicate_PredicateWithValue(bool value);
void* C_NSPredicate_PredicateFromMetadataQueryString(void* queryString);
bool C_NSPredicate_EvaluateWithObject(void* ptr, void* object);
void C_NSPredicate_AllowEvaluation(void* ptr);
void* C_NSPredicate_PredicateFormat(void* ptr);
