// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var URLClass = _URLClass{objc.GetClass("NSURL")}

type _URLClass struct {
	objc.Class
}

type IURL interface {
	objc.IObject
	GetFileSystemRepresentation_MaxLength(buffer *byte, maxBufferLength uint) bool
	CheckResourceIsReachableAndReturnError(error *Error) bool
	IsFileReferenceURL() bool
	ResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object
	SetResourceValue_ForKey_Error(value objc.IObject, key URLResourceKey, error *Error) bool
	SetResourceValues_Error(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey)
	SetTemporaryResourceValue_ForKey(value objc.IObject, key URLResourceKey)
	FileReferenceURL() URL
	URLByAppendingPathComponent(pathComponent string) URL
	URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL
	URLByAppendingPathExtension(pathExtension string) URL
	BookmarkDataWithOptions_IncludingResourceValuesForKeys_RelativeToURL_Error(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	CheckPromisedItemIsReachableAndReturnError(error *Error) bool
	GetPromisedItemResourceValue_ForKey_Error(value *objc.Object, key URLResourceKey, error *Error) bool
	PromisedItemResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object
	// deprecated
	LoadResourceDataNotifyingClient_UsingCache(client objc.IObject, shouldUseCache bool)
	// deprecated
	ResourceDataUsingCache(shouldUseCache bool) []byte
	// deprecated
	SetResourceData(data []byte) bool
	// deprecated
	PropertyForKey(propertyKey string) objc.Object
	// deprecated
	SetProperty_ForKey(property objc.IObject, propertyKey string) bool
	DataRepresentation() []byte
	IsFileURL() bool
	AbsoluteString() string
	AbsoluteURL() URL
	BaseURL() URL
	FileSystemRepresentation() *byte
	Fragment() string
	Host() string
	LastPathComponent() string
	// deprecated
	ParameterString() string
	Password() string
	Path() string
	PathComponents() []string
	PathExtension() string
	Port() Number
	Query() string
	RelativePath() string
	RelativeString() string
	ResourceSpecifier() string
	Scheme() string
	StandardizedURL() URL
	User() string
	FilePathURL() URL
	URLByDeletingLastPathComponent() URL
	URLByDeletingPathExtension() URL
	URLByResolvingSymlinksInPath() URL
	URLByStandardizingPath() URL
	HasDirectoryPath() bool
}

type URL struct {
	objc.Object
}

func MakeURL(ptr unsafe.Pointer) URL {
	return URL{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _URLClass) URLWithString(URLString string) URL {
	rv := objc.CallMethod[URL](uc, "URLWithString:", URLString)
	return rv
}

func (u_ URL) InitWithString(URLString string) URL {
	rv := objc.CallMethod[URL](u_, "initWithString:", URLString)
	return rv
}

func (uc _URLClass) URLWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "URLWithString:relativeToURL:", URLString, baseURL)
	return rv
}

func (u_ URL) InitWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initWithString:relativeToURL:", URLString, baseURL)
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](u_, "initFileURLWithPath:isDirectory:", path, isDir)
	return rv
}

func (u_ URL) InitFileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initFileURLWithPath:relativeToURL:", path, baseURL)
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initFileURLWithPath:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (u_ URL) InitFileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](u_, "initFileURLWithPath:", path)
	return rv
}

func (uc _URLClass) URLByResolvingAliasFileAtURL_Options_Error(url IURL, options URLBookmarkResolutionOptions, error *Error) URL {
	rv := objc.CallMethod[URL](uc, "URLByResolvingAliasFileAtURL:options:error:", url, options, unsafe.Pointer(error))
	return rv
}

func (uc _URLClass) URLByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](uc, "URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:", bookmarkData, options, relativeURL, isStale, unsafe.Pointer(error))
	return rv
}

func (u_ URL) InitByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](u_, "initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:", bookmarkData, options, relativeURL, isStale, unsafe.Pointer(error))
	return rv
}

func (u_ URL) InitFileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (u_ URL) InitAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initAbsoluteURLWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (u_ URL) InitWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, "initWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (u_ URL) InitWithScheme_Host_Path(scheme string, host string, path string) URL {
	rv := objc.CallMethod[URL](u_, "initWithScheme:host:path:", scheme, host, path)
	return rv
}

func (uc _URLClass) Alloc() URL {
	rv := objc.CallMethod[URL](uc, "alloc")
	return rv
}

func (uc _URLClass) New() URL {
	rv := objc.CallMethod[URL](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURL() URL {
	return URLClass.New()
}

func (u_ URL) Init() URL {
	rv := objc.CallMethod[URL](u_, "init")
	return rv
}

func (uc _URLClass) FileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithPath:isDirectory:", path, isDir)
	return rv
}

func (uc _URLClass) FileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithPath:relativeToURL:", path, baseURL)
	return rv
}

func (uc _URLClass) FileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithPath:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (uc _URLClass) FileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithPath:", path)
	return rv
}

func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithPathComponents:", components)
	return rv
}

func (uc _URLClass) FileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (u_ URL) GetFileSystemRepresentation_MaxLength(buffer *byte, maxBufferLength uint) bool {
	rv := objc.CallMethod[bool](u_, "getFileSystemRepresentation:maxLength:", buffer, maxBufferLength)
	return rv
}

func (uc _URLClass) AbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "absoluteURLWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (uc _URLClass) URLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, "URLWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (u_ URL) CheckResourceIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, "checkResourceIsReachableAndReturnError:", unsafe.Pointer(error))
	return rv
}

func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.CallMethod[bool](u_, "isFileReferenceURL")
	return rv
}

func (u_ URL) ResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, "resourceValuesForKeys:error:", keys, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValue_ForKey_Error(value objc.IObject, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, "setResourceValue:forKey:error:", value, key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValues_Error(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool {
	rv := objc.CallMethod[bool](u_, "setResourceValues:error:", keyedValues, unsafe.Pointer(error))
	return rv
}

func (u_ URL) RemoveAllCachedResourceValues() {
	objc.CallMethod[objc.Void](u_, "removeAllCachedResourceValues")
}

func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, "removeCachedResourceValueForKey:", key)
}

func (u_ URL) SetTemporaryResourceValue_ForKey(value objc.IObject, key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, "setTemporaryResourceValue:forKey:", value, key)
}

func (u_ URL) FileReferenceURL() URL {
	rv := objc.CallMethod[URL](u_, "fileReferenceURL")
	return rv
}

func (u_ URL) URLByAppendingPathComponent(pathComponent string) URL {
	rv := objc.CallMethod[URL](u_, "URLByAppendingPathComponent:", pathComponent)
	return rv
}

func (u_ URL) URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL {
	rv := objc.CallMethod[URL](u_, "URLByAppendingPathComponent:isDirectory:", pathComponent, isDirectory)
	return rv
}

func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := objc.CallMethod[URL](u_, "URLByAppendingPathExtension:", pathExtension)
	return rv
}

func (uc _URLClass) BookmarkDataWithContentsOfURL_Error(bookmarkFileURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](uc, "bookmarkDataWithContentsOfURL:error:", bookmarkFileURL, unsafe.Pointer(error))
	return rv
}

func (u_ URL) BookmarkDataWithOptions_IncludingResourceValuesForKeys_RelativeToURL_Error(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](u_, "bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:", options, keys, relativeURL, unsafe.Pointer(error))
	return rv
}

func (uc _URLClass) ResourceValuesForKeys_FromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](uc, "resourceValuesForKeys:fromBookmarkData:", keys, bookmarkData)
	return rv
}

func (uc _URLClass) WriteBookmarkData_ToURL_Options_Error(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error *Error) bool {
	rv := objc.CallMethod[bool](uc, "writeBookmarkData:toURL:options:error:", bookmarkData, bookmarkFileURL, options, unsafe.Pointer(error))
	return rv
}

func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.CallMethod[bool](u_, "startAccessingSecurityScopedResource")
	return rv
}

func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.CallMethod[objc.Void](u_, "stopAccessingSecurityScopedResource")
}

func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, "checkPromisedItemIsReachableAndReturnError:", unsafe.Pointer(error))
	return rv
}

func (u_ URL) GetPromisedItemResourceValue_ForKey_Error(value *objc.Object, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, "getPromisedItemResourceValue:forKey:error:", unsafe.Pointer(value), key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) PromisedItemResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, "promisedItemResourceValuesForKeys:error:", keys, unsafe.Pointer(error))
	return rv
}

// deprecated
func (u_ URL) LoadResourceDataNotifyingClient_UsingCache(client objc.IObject, shouldUseCache bool) {
	objc.CallMethod[objc.Void](u_, "loadResourceDataNotifyingClient:usingCache:", client, shouldUseCache)
}

// deprecated
func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) []byte {
	rv := objc.CallMethod[[]byte](u_, "resourceDataUsingCache:", shouldUseCache)
	return rv
}

// deprecated
func (u_ URL) SetResourceData(data []byte) bool {
	rv := objc.CallMethod[bool](u_, "setResourceData:", data)
	return rv
}

// deprecated
func (u_ URL) PropertyForKey(propertyKey string) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, "propertyForKey:", propertyKey)
	return rv
}

// deprecated
func (u_ URL) SetProperty_ForKey(property objc.IObject, propertyKey string) bool {
	rv := objc.CallMethod[bool](u_, "setProperty:forKey:", property, propertyKey)
	return rv
}

func (u_ URL) DataRepresentation() []byte {
	rv := objc.CallMethod[[]byte](u_, "dataRepresentation")
	return rv
}

func (u_ URL) IsFileURL() bool {
	rv := objc.CallMethod[bool](u_, "isFileURL")
	return rv
}

func (u_ URL) AbsoluteString() string {
	rv := objc.CallMethod[string](u_, "absoluteString")
	return rv
}

func (u_ URL) AbsoluteURL() URL {
	rv := objc.CallMethod[URL](u_, "absoluteURL")
	return rv
}

func (u_ URL) BaseURL() URL {
	rv := objc.CallMethod[URL](u_, "baseURL")
	return rv
}

func (u_ URL) FileSystemRepresentation() *byte {
	rv := objc.CallMethod[*byte](u_, "fileSystemRepresentation")
	return rv
}

func (u_ URL) Fragment() string {
	rv := objc.CallMethod[string](u_, "fragment")
	return rv
}

func (u_ URL) Host() string {
	rv := objc.CallMethod[string](u_, "host")
	return rv
}

func (u_ URL) LastPathComponent() string {
	rv := objc.CallMethod[string](u_, "lastPathComponent")
	return rv
}

// deprecated
func (u_ URL) ParameterString() string {
	rv := objc.CallMethod[string](u_, "parameterString")
	return rv
}

func (u_ URL) Password() string {
	rv := objc.CallMethod[string](u_, "password")
	return rv
}

func (u_ URL) Path() string {
	rv := objc.CallMethod[string](u_, "path")
	return rv
}

func (u_ URL) PathComponents() []string {
	rv := objc.CallMethod[[]string](u_, "pathComponents")
	return rv
}

func (u_ URL) PathExtension() string {
	rv := objc.CallMethod[string](u_, "pathExtension")
	return rv
}

func (u_ URL) Port() Number {
	rv := objc.CallMethod[Number](u_, "port")
	return rv
}

func (u_ URL) Query() string {
	rv := objc.CallMethod[string](u_, "query")
	return rv
}

func (u_ URL) RelativePath() string {
	rv := objc.CallMethod[string](u_, "relativePath")
	return rv
}

func (u_ URL) RelativeString() string {
	rv := objc.CallMethod[string](u_, "relativeString")
	return rv
}

func (u_ URL) ResourceSpecifier() string {
	rv := objc.CallMethod[string](u_, "resourceSpecifier")
	return rv
}

func (u_ URL) Scheme() string {
	rv := objc.CallMethod[string](u_, "scheme")
	return rv
}

func (u_ URL) StandardizedURL() URL {
	rv := objc.CallMethod[URL](u_, "standardizedURL")
	return rv
}

func (u_ URL) User() string {
	rv := objc.CallMethod[string](u_, "user")
	return rv
}

func (u_ URL) FilePathURL() URL {
	rv := objc.CallMethod[URL](u_, "filePathURL")
	return rv
}

func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := objc.CallMethod[URL](u_, "URLByDeletingLastPathComponent")
	return rv
}

func (u_ URL) URLByDeletingPathExtension() URL {
	rv := objc.CallMethod[URL](u_, "URLByDeletingPathExtension")
	return rv
}

func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := objc.CallMethod[URL](u_, "URLByResolvingSymlinksInPath")
	return rv
}

func (u_ URL) URLByStandardizingPath() URL {
	rv := objc.CallMethod[URL](u_, "URLByStandardizingPath")
	return rv
}

func (u_ URL) HasDirectoryPath() bool {
	rv := objc.CallMethod[bool](u_, "hasDirectoryPath")
	return rv
}
