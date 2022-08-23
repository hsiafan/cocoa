package ffi

// #import <stdint.h>
//
// void* New_ProtocolImpl(void* class, uintptr_t goID);
// uintptr_t ProtocolImpl_GetGoID(void* ptr);
import "C"
import (
	"reflect"
	"runtime/cgo"
	"strings"
	"sync"
	"unsafe"

	"github.com/hsiafan/cocoa/internal"

	"github.com/hsiafan/cocoa/objc"
)

var classCache = map[string]*classInfo{} //go protocol interface name to ClassInfo
var classLock sync.Mutex
var baseClass = objc.GetClass("ProtocolImplBase")

type classInfo struct {
	class       objc.Class
	methodInfos map[string]*methodInfo // selector name to MethodInfo
}

type methodInfo struct {
	required bool          // if is required protocol method
	hasFunc  reflect.Value // the hasXXX func for this method
	fun      reflect.Value // the func for this method
}

func createClass(t reflect.Type, protocolName string) *classInfo {
	classLock.Lock()
	defer classLock.Unlock()
	if ci, ok := classCache[protocolName]; ok {
		return ci
	}
	class := objc.AllocateClassPair(baseClass, protocolName+"Adaptor", 0)
	protocol := objc.GetProtocol(protocolName)
	class.AddProtocol(protocol)

	var methodInfos = map[string]*methodInfo{} // selector name to method signature
	for _, md := range getProtocolMethods(protocol) {
		selector := md.Name
		selName := selector.GetName()
		goFuncName := internal.SelectorToGoName(selName)
		fun, ok := t.MethodByName(goFuncName)
		if !ok {
			if md.required {
				panic("required method not implemented:" + selName)
			} else {
				continue
			}
		}
		hasFunc, _ := t.MethodByName("Implements" + goFuncName)

		methodInfos[selName] = &methodInfo{
			required: md.required,
			hasFunc:  hasFunc.Func,
			fun:      fun.Func,
		}
	}

	ci := &classInfo{
		class:       class,
		methodInfos: methodInfos,
	}

	classCache[protocolName] = ci
	return ci
}

type methodDescription struct {
	objc.MethodDescription
	required       bool
	instanceMethod bool
}

func getProtocolMethods(protocol objc.Protocol) []methodDescription {
	if protocol.GetName() == "NSObject" {
		return nil
	}
	var mds []methodDescription
	for _, md := range protocol.CopyMethodDescriptionList(true, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          true,
			instanceMethod:    true,
		})
	}

	for _, md := range protocol.CopyMethodDescriptionList(false, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          false,
			instanceMethod:    true,
		})
	}

	for _, parent := range protocol.CopyProtocolList() {
		mds = append(mds, getProtocolMethods(parent)...)
	}
	return mds
}

type instanceInfo struct {
	instance     any        // the go instance
	protocolName string     // go protocol interface name
	classInfo    *classInfo // the class info for this instance
}

// param protocolName: the delegate go interface name
// param d: the delegate go implementaion
func CreateProtocol[T any](d T) objc.Object {
	dv := reflect.ValueOf(d)
	protocolName := getProtocolName[T]()
	ci := createClass(dv.Type(), protocolName)
	ii := &instanceInfo{
		classInfo:    ci,
		instance:     d,
		protocolName: protocolName,
	}
	h := cgo.NewHandle(ii)
	return objc.MakeObject(C.New_ProtocolImpl(ci.class.Ptr(), C.uintptr_t(h)))
}

func getProtocolName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	pkgName := t.PkgPath()
	idx := strings.LastIndexByte(pkgName, '/')
	if idx > 0 {
		pkgName = pkgName[idx+1:]
	}

	pi, ok := GetProtocolInfo(pkgName + "." + t.Name())
	if !ok || len(pi.Name) == 0 {
		panic("protocol meta not found")
	}
	return pi.Name
}

//export respondsTo
func respondsTo(goID uintptr, sel unsafe.Pointer) bool {
	selName := objc.MakeSelector(sel).GetName()
	ii := cgo.Handle(goID).Value().(*instanceInfo)
	mi := ii.classInfo.methodInfos[selName]
	if mi == nil {
		return false
	}

	v := reflect.ValueOf(ii.instance)
	if mi.required {
		return true
	}

	return mi.hasFunc.Call([]reflect.Value{v})[0].Bool()
}

//export internalDeleteHandle
func internalDeleteHandle(hp uintptr) {
	cgo.Handle(hp).Delete()
}

//export handleDelegateInvocation
func handleDelegateInvocation(goID uintptr, sel unsafe.Pointer, invocationPtr unsafe.Pointer) {
	invocation := makeInvocation(invocationPtr)
	selName := objc.MakeSelector(sel).GetName()
	ii := cgo.Handle(goID).Value().(*instanceInfo)
	f := ii.classInfo.methodInfos[selName].fun

	ft := f.Type()
	args := make([]reflect.Value, ft.NumIn())
	args[0] = reflect.ValueOf(ii.instance)

	argsPtr := make([]unsafe.Pointer, len(args)-1)
	for i := 1; i < ft.NumIn(); i++ {
		argsPtr[i-1] = parseGoType(ft.In(i))
	}
	if len(argsPtr) > 0 {
		invocation.getArguments(2, argsPtr)
	}
	for i := 1; i < ft.NumIn(); i++ {
		args[i] = convertToGoValue(argsPtr[i-1], ft.In(i))
	}

	rvs := f.Call(args)
	if len(rvs) > 1 {
		panic("too many return values")
	}
	if len(rvs) == 1 {
		rp := convertToObjcValue(rvs[0].Interface())
		invocation.setReturnValue(rp)
	}
}
