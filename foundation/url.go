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
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithString:"), URLString)
	return rv
}

func (u_ URL) InitWithString(URLString string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithString:"), URLString)
	return rv
}

func (uc _URLClass) URLWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithString:relativeToURL:"), URLString, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithString:relativeToURL:"), URLString, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:isDirectory:"), path, isDir)
	return rv
}

func (u_ URL) InitFileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:relativeToURL:"), path, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitFileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithPath:"), path)
	return rv
}

func (uc _URLClass) URLByResolvingAliasFileAtURL_Options_Error(url IURL, options URLBookmarkResolutionOptions, error *Error) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLByResolvingAliasFileAtURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(error))
	return rv
}

func (uc _URLClass) URLByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.ExtractPtr(relativeURL), isStale, unsafe.Pointer(error))
	return rv
}

func (u_ URL) InitByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, objc.ExtractPtr(relativeURL), isStale, unsafe.Pointer(error))
	return rv
}

func (u_ URL) InitFileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) InitWithScheme_Host_Path(scheme string, host string, path string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("initWithScheme:host:path:"), scheme, host, path)
	return rv
}

func (uc _URLClass) Alloc() URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("alloc"))
	return rv
}

func (uc _URLClass) New() URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewURL() URL {
	return URLClass.New()
}

func (u_ URL) Init() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("init"))
	return rv
}

func (uc _URLClass) FileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:isDirectory:"), path, isDir)
	return rv
}

func (uc _URLClass) FileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:relativeToURL:"), path, objc.ExtractPtr(baseURL))
	return rv
}

func (uc _URLClass) FileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func (uc _URLClass) FileURLWithPath(path string) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPath:"), path)
	return rv
}

func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithPathComponents:"), components)
	return rv
}

func (uc _URLClass) FileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) GetFileSystemRepresentation_MaxLength(buffer *byte, maxBufferLength uint) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}

func (uc _URLClass) AbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("absoluteURLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func (uc _URLClass) URLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := objc.CallMethod[URL](uc, objc.GetSelector("URLWithDataRepresentation:relativeToURL:"), data, objc.ExtractPtr(baseURL))
	return rv
}

func (u_ URL) CheckResourceIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("checkResourceIsReachableAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isFileReferenceURL"))
	return rv
}

func (u_ URL) ResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, objc.GetSelector("resourceValuesForKeys:error:"), keys, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValue_ForKey_Error(value objc.IObject, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setResourceValue:forKey:error:"), objc.ExtractPtr(value), key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) SetResourceValues_Error(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setResourceValues:error:"), keyedValues, unsafe.Pointer(error))
	return rv
}

func (u_ URL) RemoveAllCachedResourceValues() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllCachedResourceValues"))
}

func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeCachedResourceValueForKey:"), key)
}

func (u_ URL) SetTemporaryResourceValue_ForKey(value objc.IObject, key URLResourceKey) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setTemporaryResourceValue:forKey:"), objc.ExtractPtr(value), key)
}

func (u_ URL) FileReferenceURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("fileReferenceURL"))
	return rv
}

func (u_ URL) URLByAppendingPathComponent(pathComponent string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathComponent:"), pathComponent)
	return rv
}

func (u_ URL) URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathComponent:isDirectory:"), pathComponent, isDirectory)
	return rv
}

func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByAppendingPathExtension:"), pathExtension)
	return rv
}

func (uc _URLClass) BookmarkDataWithContentsOfURL_Error(bookmarkFileURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](uc, objc.GetSelector("bookmarkDataWithContentsOfURL:error:"), objc.ExtractPtr(bookmarkFileURL), unsafe.Pointer(error))
	return rv
}

func (u_ URL) BookmarkDataWithOptions_IncludingResourceValuesForKeys_RelativeToURL_Error(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, objc.ExtractPtr(relativeURL), unsafe.Pointer(error))
	return rv
}

func (uc _URLClass) ResourceValuesForKeys_FromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](uc, objc.GetSelector("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}

func (uc _URLClass) WriteBookmarkData_ToURL_Options_Error(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error *Error) bool {
	rv := objc.CallMethod[bool](uc, objc.GetSelector("writeBookmarkData:toURL:options:error:"), bookmarkData, objc.ExtractPtr(bookmarkFileURL), options, unsafe.Pointer(error))
	return rv
}

func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("startAccessingSecurityScopedResource"))
	return rv
}

func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("stopAccessingSecurityScopedResource"))
}

func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("checkPromisedItemIsReachableAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (u_ URL) GetPromisedItemResourceValue_ForKey_Error(value *objc.Object, key URLResourceKey, error *Error) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("getPromisedItemResourceValue:forKey:error:"), unsafe.Pointer(value), key, unsafe.Pointer(error))
	return rv
}

func (u_ URL) PromisedItemResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := objc.CallMethod[map[URLResourceKey]objc.Object](u_, objc.GetSelector("promisedItemResourceValuesForKeys:error:"), keys, unsafe.Pointer(error))
	return rv
}

// deprecated
func (u_ URL) LoadResourceDataNotifyingClient_UsingCache(client objc.IObject, shouldUseCache bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("loadResourceDataNotifyingClient:usingCache:"), objc.ExtractPtr(client), shouldUseCache)
}

// deprecated
func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("resourceDataUsingCache:"), shouldUseCache)
	return rv
}

// deprecated
func (u_ URL) SetResourceData(data []byte) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setResourceData:"), data)
	return rv
}

// deprecated
func (u_ URL) PropertyForKey(propertyKey string) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.GetSelector("propertyForKey:"), propertyKey)
	return rv
}

// deprecated
func (u_ URL) SetProperty_ForKey(property objc.IObject, propertyKey string) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("setProperty:forKey:"), objc.ExtractPtr(property), propertyKey)
	return rv
}

func (u_ URL) DataRepresentation() []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("dataRepresentation"))
	return rv
}

func (u_ URL) IsFileURL() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isFileURL"))
	return rv
}

func (u_ URL) AbsoluteString() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("absoluteString"))
	return rv
}

func (u_ URL) AbsoluteURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("absoluteURL"))
	return rv
}

func (u_ URL) BaseURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("baseURL"))
	return rv
}

func (u_ URL) FileSystemRepresentation() *byte {
	rv := objc.CallMethod[*byte](u_, objc.GetSelector("fileSystemRepresentation"))
	return rv
}

func (u_ URL) Fragment() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("fragment"))
	return rv
}

func (u_ URL) Host() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("host"))
	return rv
}

func (u_ URL) LastPathComponent() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("lastPathComponent"))
	return rv
}

// deprecated
func (u_ URL) ParameterString() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("parameterString"))
	return rv
}

func (u_ URL) Password() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("password"))
	return rv
}

func (u_ URL) Path() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("path"))
	return rv
}

func (u_ URL) PathComponents() []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("pathComponents"))
	return rv
}

func (u_ URL) PathExtension() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("pathExtension"))
	return rv
}

func (u_ URL) Port() Number {
	rv := objc.CallMethod[Number](u_, objc.GetSelector("port"))
	return rv
}

func (u_ URL) Query() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("query"))
	return rv
}

func (u_ URL) RelativePath() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("relativePath"))
	return rv
}

func (u_ URL) RelativeString() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("relativeString"))
	return rv
}

func (u_ URL) ResourceSpecifier() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("resourceSpecifier"))
	return rv
}

func (u_ URL) Scheme() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("scheme"))
	return rv
}

func (u_ URL) StandardizedURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("standardizedURL"))
	return rv
}

func (u_ URL) User() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("user"))
	return rv
}

func (u_ URL) FilePathURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("filePathURL"))
	return rv
}

func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByDeletingLastPathComponent"))
	return rv
}

func (u_ URL) URLByDeletingPathExtension() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByDeletingPathExtension"))
	return rv
}

func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByResolvingSymlinksInPath"))
	return rv
}

func (u_ URL) URLByStandardizingPath() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLByStandardizingPath"))
	return rv
}

func (u_ URL) HasDirectoryPath() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("hasDirectoryPath"))
	return rv
}
