package objc

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import <objc/runtime.h>
//
// void Run_WithAutoreleasePool(uintptr_t ptr);
//
// void* IMP_ImplementationWithBlock(void* bp);
// void* IMP_GetBlock(void* ptr);
// bool IMP_RemoveBlock(void* ptr);
//
// void* Objc_GetClass(const char* name);
// void* Objc_AllocateClassPair(void* superClass, const char* name, size_t extraBytes);
// void Objc_RegisterClassPair(void* class);
//
// void* Class_CreateInstance(void* cls, unsigned idxIvars);
// const char* Class_GetName(void *cls);
// void Class_SetVersion(void* cls, int);
// int Class_GetVersion(void* cls);
// void* Class_GetSuperClass(void* cls);
// bool Class_RespondsToSelectorAddMethod(void* cls, void* sel);
// bool Class_AddProtocol(void* cls, void* protocol);
// bool Class_AddMethod(void* cls, void* sel, void* imp, const char* types);
// void* Class_ReplaceMethod(void* cls, void* sel, void* imp, const char* types);
// void* Class_GetMethodImplementation(void* cls, void* sel);
// void* Class_GetMethodImplementationStret(void* cls, void* sel);
// void* Class_GetInstanceMethod(void* cls, void* sel);
// void* Class_GetClassMethod(void* cls, void* sel);
// void* Class_CopyMethodList(void* cls, unsigned int *outCount);
// void* Class_GetProperty(void* cls, const char *name);
// void* Class_CopyPropertyList(void* cls, unsigned int *outCount);
// bool Class_AddProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount);
// void Class_ReplaceProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount);
//
// void* Objc_GetProtocol(const char* name);
// void* Objc_AllocateProtocol(const char* name);
// const char* Protocol_GetName(void *protocol);
// struct objc_method_description Protocol_GetMethodDescription(void* protocol, void* sel, bool required, bool instanceMethod);
// struct objc_method_description* Protocol_CopyMethodDescriptionList(void* protocol, bool required, bool instanceMethod, unsigned int *outCount);
// void* Protocol_CopyProtocolList(void* protocol, unsigned int *outCount);
// void* Protocol_GetProperty(void* protocol, const char *name, bool required, bool isInstanceProperty);
// void* Protocol_CopyPropertyList(void* protocol, unsigned int *outCount);
//
// void* Method_GetName(void* method);
// const char* Method_GetTypeEncoding(void* method);
// void* Method_GetImplementation(void* method);
// void* Method_SetImplementation(void* method, void* imp);
//
// const char* Property_GetName(void* property);
// const char* Property_GetAttributes(void* property);
// char* Property_CopyAttributeValue(void* property, const char* name);
// objc_property_attribute_t * Property_CopyAttributeList(void* property, unsigned int *outCount);
//
// void* block_copy(void *ptr);
// void block_release(void *ptr);
//
// void* Selector_SEL_RegisterName(const char* name);
// const char* Selector_SEL_GetName(void* ptr);
//
// void Objc_SetAssociatedObject(void* ptr, const void* key, void* valuePtr, uintptr_t policy);
// void* Objc_GetAssociatedObject(void* ptr, const void* key);
// void Objc_RemoveAssociatedObjects(void* ptr);
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/hsiafan/cocoa/internal"
)

//https://developer.apple.com/documentation/objectivec/objective-c_runtime?language=objc

// IMP is function pointer
type IMP struct {
	ptr unsafe.Pointer
}

func (i IMP) Ptr() unsafe.Pointer {
	return i.ptr
}

func MakeIMP(ptr unsafe.Pointer) IMP {
	return IMP{ptr}
}

func IMPWithBlock(b Block) IMP {
	return IMP{
		ptr: C.IMP_ImplementationWithBlock(b.ptr),
	}
}

func (i IMP) GetBlock() Block {
	return Block{
		ptr: C.IMP_GetBlock(i.ptr),
	}
}

func (i IMP) RemoveBlock() bool {
	return bool(C.IMP_RemoveBlock(i.ptr))
}

// make generated code happy
type IClass interface {
	Holder
	CreateInstance(idxIvars uint) Object
	GetName() string
	SetVersion(version int)
	GetVersion()
	GetSuperClass() Class
	RespondsToSelector(sel Selector) bool
	AddMethod(sel Selector, imp IMP, types string) bool
	ReplaceMethod(sel Selector, imp IMP, types string) IMP
	GetMethodImplementation(sel Selector) IMP
	GetMethodImplementationStret(sel Selector)
	GetInstanceMethod(sel Selector) Method
	GetClassMethod(sel Selector) Method
	CopyMethodList() []Method
}

// Class is objc Class
type Class struct {
	ptr unsafe.Pointer
}

func (c Class) Ptr() unsafe.Pointer {
	return c.ptr
}

// GetClass get and return an objc Class by name
func GetClass(name string) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_GetClass(nameStr),
	}
}

func AllocateClassPair(superClass Class, name string, extraBytes uint) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_AllocateClassPair(superClass.ptr, nameStr, C.size_t(extraBytes)),
	}
}

func RegisterClassPair(class Class) {
	C.Objc_RegisterClassPair(class.ptr)
}

func (c Class) CreateInstance(idxIvars uint) Object {
	ptr := C.Class_CreateInstance(c.ptr, C.uint(idxIvars))
	return MakeObject(ptr)
}

func (c Class) GetName() string {
	cname := C.Class_GetName(c.ptr)
	name := C.GoString(cname)
	return name
}

// GetClass returns class's meta class.
func (c Class) GetClass() Class {
	return MakeObject(c.ptr).GetClass()
}

func (c Class) SetVersion(version int) {
	C.Class_SetVersion(c.ptr, C.int(version))
}

func (c Class) GetVersion() int {
	return int(C.Class_GetVersion(c.ptr))
}

func (c Class) GetSuperClass() Class {
	ptr := C.Class_GetSuperClass(c.ptr)
	return Class{ptr}
}

func (c Class) RespondsToSelector(sel Selector) bool {
	r := C.Class_RespondsToSelectorAddMethod(c.ptr, sel.ptr)
	return bool(r)
}

func (c Class) AddMethod(sel Selector, imp IMP, types string) bool {
	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))
	r := C.Class_AddMethod(c.ptr, sel.ptr, imp.ptr, ctypes)
	return bool(r)
}

func (c Class) ReplaceMethod(sel Selector, imp IMP, types string) IMP {
	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))
	r := C.Class_ReplaceMethod(c.ptr, sel.ptr, imp.ptr, ctypes)
	return IMP{ptr: r}
}

func (c Class) GetMethodImplementation(sel Selector) IMP {
	r := C.Class_GetMethodImplementation(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) GetMethodImplementationStret(sel Selector) IMP {
	r := C.Class_GetMethodImplementationStret(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) GetInstanceMethod(sel Selector) Method {
	r := C.Class_GetInstanceMethod(c.ptr, sel.ptr)
	return Method{ptr: r}
}

func (c Class) GetClassMethod(sel Selector) Method {
	r := C.Class_GetClassMethod(c.ptr, sel.ptr)
	return Method{ptr: r}
}

func (c Class) CopyMethodList() []Method {
	var count C.uint
	rp := C.Class_CopyMethodList(c.ptr, &count)
	return convertToSliceAndFreePointer[Method](rp, int(count))
}

func (c Class) AddProtocol(protocol Protocol) bool {
	r := C.Class_AddProtocol(c.ptr, protocol.ptr)
	return bool(r)
}

func (c Class) GetProperty(name string) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Property{
		ptr: C.Class_GetProperty(c.Ptr(), cname),
	}
}

func (c Class) CopyPropertyList() []Property {
	var outCount C.uint
	pp := C.Class_CopyPropertyList(c.Ptr(), &outCount)
	return convertToSliceAndFreePointer[Property](pp, int(outCount))
}

func (c Class) AddProperty0(property Property) bool {
	cname := C.Property_GetName(property.Ptr())
	var outCount C.uint
	pap := unsafe.Pointer(C.Property_CopyAttributeList(property.Ptr(), &outCount))
	if pap != nil {
		defer C.free(pap)
	}
	r := C.Class_AddProperty(c.Ptr(), cname, pap, outCount)
	return bool(r)
}

func (c Class) AddProperty(name string, attributes []PropertyAttribute) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cattributes := convertToObjcPropertyAttributes(attributes)
	defer func() {
		for _, ca := range cattributes {
			C.free(unsafe.Pointer(ca.name))
			C.free(unsafe.Pointer(ca.value))
		}
	}()
	r := C.Class_AddProperty(c.Ptr(), cname, unsafe.Pointer(&cattributes[0]), C.uint(len(attributes)))
	return bool(r)
}

func (c Class) ReplaceProperty(name string, attributes []PropertyAttribute) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cattributes := convertToObjcPropertyAttributes(attributes)
	defer func() {
		for _, ca := range cattributes {
			C.free(unsafe.Pointer(ca.name))
			C.free(unsafe.Pointer(ca.value))
		}
	}()
	C.Class_ReplaceProperty(c.Ptr(), cname, unsafe.Pointer(&cattributes[0]), C.uint(len(attributes)))
}

func convertToObjcPropertyAttributes(attributes []PropertyAttribute) []*C.objc_property_attribute_t {
	aps := make([]*C.objc_property_attribute_t, len(attributes))
	for i := 0; i < len(attributes); i++ {
		a := attributes[i]
		aps[i] = &C.objc_property_attribute_t{
			name:  C.CString(a.Name),
			value: C.CString(a.Value),
		}
	}
	return aps
}

type Protocol struct {
	ptr unsafe.Pointer
}

func (p Protocol) Ptr() unsafe.Pointer {
	return p.ptr
}

func GetProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Protocol{
		ptr: C.Objc_GetProtocol(cname),
	}
}

func AllocateProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Protocol{
		ptr: C.Objc_AllocateProtocol(cname),
	}
}

func (p Protocol) GetName() string {
	cname := C.Protocol_GetName(p.ptr)
	name := C.GoString(cname)
	return name
}

func (p Protocol) GetMethodDescription(sel Selector, required bool, instanceMethod bool) MethodDescription {
	md := C.Protocol_GetMethodDescription(p.ptr, sel.ptr, C.bool(required), C.bool(instanceMethod))
	return MethodDescription{
		Name:  Selector{unsafe.Pointer(md.name)},
		Types: C.GoString(md.types),
	}
}

func (p Protocol) CopyMethodDescriptionList(required bool, instanceMethod bool) []MethodDescription {
	var count C.uint
	mdp := C.Protocol_CopyMethodDescriptionList(p.ptr, C.bool(required), C.bool(instanceMethod), &count)
	if mdp == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(mdp))
	mds := unsafe.Slice(mdp, int(count))
	r := make([]MethodDescription, int(count))
	for i := 0; i < int(count); i++ {
		r[i] = MethodDescription{
			Name:  Selector{unsafe.Pointer(mds[i].name)},
			Types: C.GoString(mds[i].types),
		}
	}
	return r
}

func (p Protocol) CopyProtocolList() []Protocol {
	var count C.uint
	pp := C.Protocol_CopyProtocolList(p.Ptr(), &count)
	return convertToSliceAndFreePointer[Protocol](pp, int(count))
}

func (p Protocol) GetProperty(name string, required bool, isInstanceProperty bool) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Property{
		ptr: C.Protocol_GetProperty(p.Ptr(), cname, C.bool(required), C.bool(isInstanceProperty)),
	}
}

func (p Protocol) CopyPropertyList() []Property {
	var outCount C.uint
	pp := C.Protocol_CopyPropertyList(p.Ptr(), &outCount)
	return convertToSliceAndFreePointer[Property](pp, int(outCount))
}

type MethodDescription struct {
	Name  Selector
	Types string
}

type Property struct {
	ptr unsafe.Pointer
}

func (p Property) Ptr() unsafe.Pointer {
	return p.ptr
}

type PropertyAttribute struct {
	Name  string
	Value string
}

const (
	PropertyAttributeNameNonatomic = "N"
	PropertyAttributeNameStrong    = "&"
	PropertyAttributeNameRetain    = PropertyAttributeNameStrong
	PropertyAttributeNameWeak      = "W"
	PropertyAttributeNameReadonly  = "R"
	PropertyAttributeNameGetter    = "G"
	PropertyAttributeNameSetter    = "S"
	PropertyAttributeNameIvar      = "V"
	PropertyAttributeNameType      = "T"
)

func (p Property) GetName() string {
	cs := C.Property_GetName(p.Ptr())
	return C.GoString(cs)
}

func (p Property) GetAttributes() string {
	cs := C.Property_GetAttributes(p.Ptr())
	return C.GoString(cs)
}

func (p Property) CopyAttributeValue(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cs := C.Property_CopyAttributeValue(p.Ptr(), cname)
	if cs != nil {
		defer C.free(unsafe.Pointer(cs))
	}
	return C.GoString(cs)
}

func (p Property) CopyAttributeList() []PropertyAttribute {
	var outCount C.uint
	pap := C.Property_CopyAttributeList(p.Ptr(), &outCount)
	if outCount == 0 {
		return nil
	}
	defer C.free(unsafe.Pointer(pap))
	pas := unsafe.Slice(pap, int(outCount))
	slice := make([]PropertyAttribute, len(pas))
	for i := 0; i < len(pas); i++ {
		pa := pas[i]
		slice[i] = PropertyAttribute{
			Name:  C.GoString(pa.name),
			Value: C.GoString(pa.value),
		}
	}
	return slice
}

type Method struct {
	ptr unsafe.Pointer
}

func (m Method) Ptr() unsafe.Pointer {
	return m.ptr
}

func (m Method) GetName() Selector {
	return Selector{
		ptr: C.Method_GetName(m.ptr),
	}
}

func (m Method) GetTypeEncoding() string {
	r := C.Method_GetTypeEncoding(m.ptr)
	return C.GoString(r)
}

func (m Method) GetImplementation() IMP {
	return IMP{
		ptr: C.Method_GetImplementation(m.ptr),
	}
}

func (m Method) SetImplementation(imp IMP) IMP {
	return IMP{
		ptr: C.Method_SetImplementation(m.ptr, imp.ptr),
	}
}

type Ivar struct {
	ptr unsafe.Pointer
}

func (i Ivar) Ptr() unsafe.Pointer {
	return i.Ptr()
}

type Category struct {
	ptr unsafe.Pointer
}

func (c Category) Ptr() unsafe.Pointer {
	return c.Ptr()
}

// Selector for select and hold method
type Selector struct {
	ptr unsafe.Pointer
}

func MakeSelector(ptr unsafe.Pointer) Selector {
	return Selector{ptr}
}

var selectorCache = internal.SyncCache[string, Selector]{
	Compute: func(selName string) Selector {
		return SelectorRegisterName(selName)
	},
}

// GetSelector return a method selector by the name. The selector is cached at go side.
func GetSelector(selName string) Selector {
	return selectorCache.Load(selName)
}

// SelectorRegisterName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value
func SelectorRegisterName(name string) Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Selector{C.Selector_SEL_RegisterName(cname)}
}

func (s Selector) Ptr() unsafe.Pointer {
	return s.ptr
}

// GetName return selector name
func (s Selector) GetName() string {
	cstr := C.Selector_SEL_GetName(s.ptr)
	return C.GoString(cstr)
}

type Block struct {
	ptr unsafe.Pointer
}

func MakeBlock(ptr unsafe.Pointer) Block {
	return Block{ptr: ptr}
}

func (b Block) Ptr() unsafe.Pointer {
	return b.ptr
}

func (b Block) Copy() Block {
	return Block{C.block_copy(b.Ptr())}
}

func (b Block) Release() {
	C.block_release(b.Ptr())
}

type AssociationPolicy uintptr

const (
	ASSOCIATION_ASSIGN           = 0     // Specifies a weak reference to the associated object.
	ASSOCIATION_RETAIN_NONATOMIC = 1     // Specifies a strong reference to the associated object. The association is not made atomically.
	ASSOCIATION_COPY_NONATOMIC   = 3     // Specifies that the associated object is copied. The association is not made atomically.
	ASSOCIATION_RETAIN           = 01401 // Specifies a strong reference to the associated object. The association is made atomically.
	ASSOCIATION_COPY             = 01403 //Specifies that the associated object is copied. The association is made atomically.
)

func SetAssociatedObject(o IObject, key unsafe.Pointer, value IObject, policy AssociationPolicy) {
	C.Objc_SetAssociatedObject(o.Ptr(), key, value.Ptr(), C.uintptr_t(policy))
}

func GetAssociatedObject(o IObject, key unsafe.Pointer) Object {
	return MakeObject(C.Objc_GetAssociatedObject(o.Ptr(), key))
}

func RemoveAssociatedObjects(o IObject) {
	C.Objc_RemoveAssociatedObjects(o.Ptr())
}

// WithAutoreleasePool run code in a new auto release pool.
func WithAutoreleasePool(task func()) {
	id := cgo.NewHandle(task)
	C.Run_WithAutoreleasePool(C.uintptr_t(id))
}

//export runTaskAndDeleteHandle
func runTaskAndDeleteHandle(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}

func convertToSliceAndFreePointer[T Holder](p unsafe.Pointer, count int) []T {
	if p == nil {
		return nil
	}
	defer C.free(p)
	ps := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(p)), count)
	slice := make([]T, count)
	for i := 0; i < int(count); i++ {
		slice[i] = internal.ForceCast[unsafe.Pointer, T](ps[i])
	}
	return slice
}
