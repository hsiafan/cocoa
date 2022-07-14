// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[URL](uc, "URLWithString:", URLString)
	return rv
}

func (u_ URL) InitWithString(URLString string) URL {
	rv := ffi.CallMethod[URL](u_, "initWithString:", URLString)
	rv.Autorelease()
	return rv
}

func (uc _URLClass) URLWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "URLWithString:relativeToURL:", URLString, baseURL)
	return rv
}

func (u_ URL) InitWithString_RelativeToURL(URLString string, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initWithString:relativeToURL:", URLString, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := ffi.CallMethod[URL](u_, "initFileURLWithPath:isDirectory:", path, isDir)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitFileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initFileURLWithPath:relativeToURL:", path, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initFileURLWithPath:isDirectory:relativeToURL:", path, isDir, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitFileURLWithPath(path string) URL {
	rv := ffi.CallMethod[URL](u_, "initFileURLWithPath:", path)
	rv.Autorelease()
	return rv
}

func (uc _URLClass) URLByResolvingAliasFileAtURL_Options_Error(url IURL, options URLBookmarkResolutionOptions, error *Error) URL {
	rv := ffi.CallMethod[URL](uc, "URLByResolvingAliasFileAtURL:options:error:", url, options, error)
	return rv
}

func (uc _URLClass) URLByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := ffi.CallMethod[URL](uc, "URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:", bookmarkData, options, relativeURL, isStale, error)
	return rv
}

func (u_ URL) InitByResolvingBookmarkData_Options_RelativeToURL_BookmarkDataIsStale_Error(bookmarkData []byte, options URLBookmarkResolutionOptions, relativeURL IURL, isStale *bool, error *Error) URL {
	rv := ffi.CallMethod[URL](u_, "initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:", bookmarkData, options, relativeURL, isStale, error)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitFileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:", path, isDir, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initAbsoluteURLWithDataRepresentation:relativeToURL:", data, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](u_, "initWithDataRepresentation:relativeToURL:", data, baseURL)
	rv.Autorelease()
	return rv
}

func (u_ URL) InitWithScheme_Host_Path(scheme string, host string, path string) URL {
	rv := ffi.CallMethod[URL](u_, "initWithScheme:host:path:", scheme, host, path)
	rv.Autorelease()
	return rv
}

func (uc _URLClass) Alloc() URL {
	rv := ffi.CallMethod[URL](uc, "alloc")
	return rv
}

func (u_ URL) Init() URL {
	rv := ffi.CallMethod[URL](u_, "init")
	rv.Autorelease()
	return rv
}

func (uc _URLClass) New() URL {
	rv := ffi.CallMethod[URL](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURL() URL {
	return URLClass.New()
}

func (uc _URLClass) FileURLWithPath_IsDirectory(path string, isDir bool) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithPath:isDirectory:", path, isDir)
	return rv
}

func (uc _URLClass) FileURLWithPath_RelativeToURL(path string, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithPath:relativeToURL:", path, baseURL)
	return rv
}

func (uc _URLClass) FileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithPath:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (uc _URLClass) FileURLWithPath(path string) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithPath:", path)
	return rv
}

func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithPathComponents:", components)
	return rv
}

func (uc _URLClass) FileURLWithFileSystemRepresentation_IsDirectory_RelativeToURL(path *byte, isDir bool, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:", path, isDir, baseURL)
	return rv
}

func (u_ URL) GetFileSystemRepresentation_MaxLength(buffer *byte, maxBufferLength uint) bool {
	rv := ffi.CallMethod[bool](u_, "getFileSystemRepresentation:maxLength:", buffer, maxBufferLength)
	return rv
}

func (uc _URLClass) AbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "absoluteURLWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (uc _URLClass) URLWithDataRepresentation_RelativeToURL(data []byte, baseURL IURL) URL {
	rv := ffi.CallMethod[URL](uc, "URLWithDataRepresentation:relativeToURL:", data, baseURL)
	return rv
}

func (u_ URL) CheckResourceIsReachableAndReturnError(error *Error) bool {
	rv := ffi.CallMethod[bool](u_, "checkResourceIsReachableAndReturnError:", error)
	return rv
}

func (u_ URL) IsFileReferenceURL() bool {
	rv := ffi.CallMethod[bool](u_, "isFileReferenceURL")
	return rv
}

func (u_ URL) ResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := ffi.CallMethod[map[URLResourceKey]objc.Object](u_, "resourceValuesForKeys:error:", keys, error)
	return rv
}

func (u_ URL) SetResourceValue_ForKey_Error(value objc.IObject, key URLResourceKey, error *Error) bool {
	rv := ffi.CallMethod[bool](u_, "setResourceValue:forKey:error:", value, key, error)
	return rv
}

func (u_ URL) SetResourceValues_Error(keyedValues map[URLResourceKey]objc.IObject, error *Error) bool {
	rv := ffi.CallMethod[bool](u_, "setResourceValues:error:", keyedValues, error)
	return rv
}

func (u_ URL) RemoveAllCachedResourceValues() {
	ffi.CallMethod[ffi.Void](u_, "removeAllCachedResourceValues")
}

func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	ffi.CallMethod[ffi.Void](u_, "removeCachedResourceValueForKey:", key)
}

func (u_ URL) SetTemporaryResourceValue_ForKey(value objc.IObject, key URLResourceKey) {
	ffi.CallMethod[ffi.Void](u_, "setTemporaryResourceValue:forKey:", value, key)
}

func (u_ URL) FileReferenceURL() URL {
	rv := ffi.CallMethod[URL](u_, "fileReferenceURL")
	return rv
}

func (u_ URL) URLByAppendingPathComponent(pathComponent string) URL {
	rv := ffi.CallMethod[URL](u_, "URLByAppendingPathComponent:", pathComponent)
	return rv
}

func (u_ URL) URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL {
	rv := ffi.CallMethod[URL](u_, "URLByAppendingPathComponent:isDirectory:", pathComponent, isDirectory)
	return rv
}

func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := ffi.CallMethod[URL](u_, "URLByAppendingPathExtension:", pathExtension)
	return rv
}

func (uc _URLClass) BookmarkDataWithContentsOfURL_Error(bookmarkFileURL IURL, error *Error) []byte {
	rv := ffi.CallMethod[[]byte](uc, "bookmarkDataWithContentsOfURL:error:", bookmarkFileURL, error)
	return rv
}

func (u_ URL) BookmarkDataWithOptions_IncludingResourceValuesForKeys_RelativeToURL_Error(options URLBookmarkCreationOptions, keys []URLResourceKey, relativeURL IURL, error *Error) []byte {
	rv := ffi.CallMethod[[]byte](u_, "bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:", options, keys, relativeURL, error)
	return rv
}

func (uc _URLClass) ResourceValuesForKeys_FromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	rv := ffi.CallMethod[map[URLResourceKey]objc.Object](uc, "resourceValuesForKeys:fromBookmarkData:", keys, bookmarkData)
	return rv
}

func (uc _URLClass) WriteBookmarkData_ToURL_Options_Error(bookmarkData []byte, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error *Error) bool {
	rv := ffi.CallMethod[bool](uc, "writeBookmarkData:toURL:options:error:", bookmarkData, bookmarkFileURL, options, error)
	return rv
}

func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := ffi.CallMethod[bool](u_, "startAccessingSecurityScopedResource")
	return rv
}

func (u_ URL) StopAccessingSecurityScopedResource() {
	ffi.CallMethod[ffi.Void](u_, "stopAccessingSecurityScopedResource")
}

func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error *Error) bool {
	rv := ffi.CallMethod[bool](u_, "checkPromisedItemIsReachableAndReturnError:", error)
	return rv
}

func (u_ URL) GetPromisedItemResourceValue_ForKey_Error(value *objc.Object, key URLResourceKey, error *Error) bool {
	rv := ffi.CallMethod[bool](u_, "getPromisedItemResourceValue:forKey:error:", value, key, error)
	return rv
}

func (u_ URL) PromisedItemResourceValuesForKeys_Error(keys []URLResourceKey, error *Error) map[URLResourceKey]objc.Object {
	rv := ffi.CallMethod[map[URLResourceKey]objc.Object](u_, "promisedItemResourceValuesForKeys:error:", keys, error)
	return rv
}

// deprecated
func (u_ URL) LoadResourceDataNotifyingClient_UsingCache(client objc.IObject, shouldUseCache bool) {
	ffi.CallMethod[ffi.Void](u_, "loadResourceDataNotifyingClient:usingCache:", client, shouldUseCache)
}

// deprecated
func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) []byte {
	rv := ffi.CallMethod[[]byte](u_, "resourceDataUsingCache:", shouldUseCache)
	return rv
}

// deprecated
func (u_ URL) SetResourceData(data []byte) bool {
	rv := ffi.CallMethod[bool](u_, "setResourceData:", data)
	return rv
}

// deprecated
func (u_ URL) PropertyForKey(propertyKey string) objc.Object {
	rv := ffi.CallMethod[objc.Object](u_, "propertyForKey:", propertyKey)
	return rv
}

// deprecated
func (u_ URL) SetProperty_ForKey(property objc.IObject, propertyKey string) bool {
	rv := ffi.CallMethod[bool](u_, "setProperty:forKey:", property, propertyKey)
	return rv
}

func (u_ URL) DataRepresentation() []byte {
	rv := ffi.CallMethod[[]byte](u_, "dataRepresentation")
	return rv
}

func (u_ URL) IsFileURL() bool {
	rv := ffi.CallMethod[bool](u_, "isFileURL")
	return rv
}

func (u_ URL) AbsoluteString() string {
	rv := ffi.CallMethod[string](u_, "absoluteString")
	return rv
}

func (u_ URL) AbsoluteURL() URL {
	rv := ffi.CallMethod[URL](u_, "absoluteURL")
	return rv
}

func (u_ URL) BaseURL() URL {
	rv := ffi.CallMethod[URL](u_, "baseURL")
	return rv
}

func (u_ URL) FileSystemRepresentation() *byte {
	rv := ffi.CallMethod[*byte](u_, "fileSystemRepresentation")
	return rv
}

func (u_ URL) Fragment() string {
	rv := ffi.CallMethod[string](u_, "fragment")
	return rv
}

func (u_ URL) Host() string {
	rv := ffi.CallMethod[string](u_, "host")
	return rv
}

func (u_ URL) LastPathComponent() string {
	rv := ffi.CallMethod[string](u_, "lastPathComponent")
	return rv
}

// deprecated
func (u_ URL) ParameterString() string {
	rv := ffi.CallMethod[string](u_, "parameterString")
	return rv
}

func (u_ URL) Password() string {
	rv := ffi.CallMethod[string](u_, "password")
	return rv
}

func (u_ URL) Path() string {
	rv := ffi.CallMethod[string](u_, "path")
	return rv
}

func (u_ URL) PathComponents() []string {
	rv := ffi.CallMethod[[]string](u_, "pathComponents")
	return rv
}

func (u_ URL) PathExtension() string {
	rv := ffi.CallMethod[string](u_, "pathExtension")
	return rv
}

func (u_ URL) Port() Number {
	rv := ffi.CallMethod[Number](u_, "port")
	return rv
}

func (u_ URL) Query() string {
	rv := ffi.CallMethod[string](u_, "query")
	return rv
}

func (u_ URL) RelativePath() string {
	rv := ffi.CallMethod[string](u_, "relativePath")
	return rv
}

func (u_ URL) RelativeString() string {
	rv := ffi.CallMethod[string](u_, "relativeString")
	return rv
}

func (u_ URL) ResourceSpecifier() string {
	rv := ffi.CallMethod[string](u_, "resourceSpecifier")
	return rv
}

func (u_ URL) Scheme() string {
	rv := ffi.CallMethod[string](u_, "scheme")
	return rv
}

func (u_ URL) StandardizedURL() URL {
	rv := ffi.CallMethod[URL](u_, "standardizedURL")
	return rv
}

func (u_ URL) User() string {
	rv := ffi.CallMethod[string](u_, "user")
	return rv
}

func (u_ URL) FilePathURL() URL {
	rv := ffi.CallMethod[URL](u_, "filePathURL")
	return rv
}

func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := ffi.CallMethod[URL](u_, "URLByDeletingLastPathComponent")
	return rv
}

func (u_ URL) URLByDeletingPathExtension() URL {
	rv := ffi.CallMethod[URL](u_, "URLByDeletingPathExtension")
	return rv
}

func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := ffi.CallMethod[URL](u_, "URLByResolvingSymlinksInPath")
	return rv
}

func (u_ URL) URLByStandardizingPath() URL {
	rv := ffi.CallMethod[URL](u_, "URLByStandardizingPath")
	return rv
}

func (u_ URL) HasDirectoryPath() bool {
	rv := ffi.CallMethod[bool](u_, "hasDirectoryPath")
	return rv
}

var URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	objc.Class
}

type IURLRequest interface {
	objc.IObject
	ValueForHTTPHeaderField(field string) string
	CachePolicy() URLRequestCachePolicy
	HTTPMethod() string
	URL() URL
	HTTPBody() []byte
	HTTPBodyStream() InputStream
	MainDocumentURL() URL
	AllHTTPHeaderFields() map[string]string
	TimeoutInterval() TimeInterval
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	AllowsCellularAccess() bool
	AllowsConstrainedNetworkAccess() bool
	AllowsExpensiveNetworkAccess() bool
	NetworkServiceType() URLRequestNetworkServiceType
	Attribution() URLRequestAttribution
	AssumesHTTP3Capable() bool
}

type URLRequest struct {
	objc.Object
}

func MakeURLRequest(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _URLRequestClass) RequestWithURL(URL IURL) URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "requestWithURL:", URL)
	return rv
}

func (u_ URLRequest) InitWithURL(URL IURL) URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "initWithURL:", URL)
	rv.Autorelease()
	return rv
}

func (uc _URLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "requestWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (u_ URLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "initWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	rv.Autorelease()
	return rv
}

func (uc _URLRequestClass) Alloc() URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "alloc")
	return rv
}

func (u_ URLRequest) Init() URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "init")
	rv.Autorelease()
	return rv
}

func (uc _URLRequestClass) New() URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURLRequest() URLRequest {
	return URLRequestClass.New()
}

func (u_ URLRequest) ValueForHTTPHeaderField(field string) string {
	rv := ffi.CallMethod[string](u_, "valueForHTTPHeaderField:", field)
	return rv
}

func (u_ URLRequest) CachePolicy() URLRequestCachePolicy {
	rv := ffi.CallMethod[URLRequestCachePolicy](u_, "cachePolicy")
	return rv
}

func (u_ URLRequest) HTTPMethod() string {
	rv := ffi.CallMethod[string](u_, "HTTPMethod")
	return rv
}

func (u_ URLRequest) URL() URL {
	rv := ffi.CallMethod[URL](u_, "URL")
	return rv
}

func (u_ URLRequest) HTTPBody() []byte {
	rv := ffi.CallMethod[[]byte](u_, "HTTPBody")
	return rv
}

func (u_ URLRequest) HTTPBodyStream() InputStream {
	rv := ffi.CallMethod[InputStream](u_, "HTTPBodyStream")
	return rv
}

func (u_ URLRequest) MainDocumentURL() URL {
	rv := ffi.CallMethod[URL](u_, "mainDocumentURL")
	return rv
}

func (u_ URLRequest) AllHTTPHeaderFields() map[string]string {
	rv := ffi.CallMethod[map[string]string](u_, "allHTTPHeaderFields")
	return rv
}

func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](u_, "timeoutInterval")
	return rv
}

func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := ffi.CallMethod[bool](u_, "HTTPShouldHandleCookies")
	return rv
}

func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := ffi.CallMethod[bool](u_, "HTTPShouldUsePipelining")
	return rv
}

func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsCellularAccess")
	return rv
}

func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsConstrainedNetworkAccess")
	return rv
}

func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsExpensiveNetworkAccess")
	return rv
}

func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := ffi.CallMethod[URLRequestNetworkServiceType](u_, "networkServiceType")
	return rv
}

func (uc _URLRequestClass) SupportsSecureCoding() bool {
	rv := ffi.CallMethod[bool](uc, "supportsSecureCoding")
	return rv
}

func (u_ URLRequest) Attribution() URLRequestAttribution {
	rv := ffi.CallMethod[URLRequestAttribution](u_, "attribution")
	return rv
}

func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := ffi.CallMethod[bool](u_, "assumesHTTP3Capable")
	return rv
}

var MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}

type _MutableURLRequestClass struct {
	objc.Class
}

type IMutableURLRequest interface {
	IURLRequest
	AddValue_ForHTTPHeaderField(value string, field string)
	SetValue_ForHTTPHeaderField(value string, field string)
	SetCachePolicy(value URLRequestCachePolicy)
	SetHTTPMethod(value string)
	SetURL(value IURL)
	SetHTTPBody(value []byte)
	SetHTTPBodyStream(value IInputStream)
	SetMainDocumentURL(value IURL)
	SetAllHTTPHeaderFields(value map[string]string)
	SetTimeoutInterval(value TimeInterval)
	SetHTTPShouldHandleCookies(value bool)
	SetHTTPShouldUsePipelining(value bool)
	SetAllowsCellularAccess(value bool)
	SetAllowsConstrainedNetworkAccess(value bool)
	SetAllowsExpensiveNetworkAccess(value bool)
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SetAttribution(value URLRequestAttribution)
	SetAssumesHTTP3Capable(value bool)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
}

type MutableURLRequest struct {
	URLRequest
}

func MakeMutableURLRequest(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: MakeURLRequest(ptr),
	}
}

func (mc _MutableURLRequestClass) RequestWithURL(URL IURL) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "requestWithURL:", URL)
	return rv
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "initWithURL:", URL)
	rv.Autorelease()
	return rv
}

func (mc _MutableURLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "requestWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (m_ MutableURLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "initWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	rv.Autorelease()
	return rv
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "alloc")
	return rv
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) AddValue_ForHTTPHeaderField(value string, field string) {
	ffi.CallMethod[ffi.Void](m_, "addValue:forHTTPHeaderField:", value, field)
}

func (m_ MutableURLRequest) SetValue_ForHTTPHeaderField(value string, field string) {
	ffi.CallMethod[ffi.Void](m_, "setValue:forHTTPHeaderField:", value, field)
}

func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	ffi.CallMethod[ffi.Void](m_, "setCachePolicy:", value)
}

func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPMethod:", value)
}

func (m_ MutableURLRequest) SetURL(value IURL) {
	ffi.CallMethod[ffi.Void](m_, "setURL:", value)
}

func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPBody:", value)
}

func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPBodyStream:", value)
}

func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	ffi.CallMethod[ffi.Void](m_, "setMainDocumentURL:", value)
}

func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	ffi.CallMethod[ffi.Void](m_, "setAllHTTPHeaderFields:", value)
}

func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	ffi.CallMethod[ffi.Void](m_, "setTimeoutInterval:", value)
}

func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPShouldHandleCookies:", value)
}

func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPShouldUsePipelining:", value)
}

func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsCellularAccess:", value)
}

func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsConstrainedNetworkAccess:", value)
}

func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsExpensiveNetworkAccess:", value)
}

func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	ffi.CallMethod[ffi.Void](m_, "setNetworkServiceType:", value)
}

func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	ffi.CallMethod[ffi.Void](m_, "setAttribution:", value)
}

func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAssumesHTTP3Capable:", value)
}

func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := ffi.CallMethod[bool](m_, "requiresDNSSECValidation")
	return rv
}

func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setRequiresDNSSECValidation:", value)
}

var URLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}

type _URLResponseClass struct {
	objc.Class
}

type IURLResponse interface {
	objc.IObject
	ExpectedContentLength() int64
	SuggestedFilename() string
	MIMEType() string
	TextEncodingName() string
	URL() URL
}

type URLResponse struct {
	objc.Object
}

func MakeURLResponse(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ URLResponse) InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	rv := ffi.CallMethod[URLResponse](u_, "initWithURL:MIMEType:expectedContentLength:textEncodingName:", URL, MIMEType, length, name)
	rv.Autorelease()
	return rv
}

func (uc _URLResponseClass) Alloc() URLResponse {
	rv := ffi.CallMethod[URLResponse](uc, "alloc")
	return rv
}

func (u_ URLResponse) Init() URLResponse {
	rv := ffi.CallMethod[URLResponse](u_, "init")
	rv.Autorelease()
	return rv
}

func (uc _URLResponseClass) New() URLResponse {
	rv := ffi.CallMethod[URLResponse](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURLResponse() URLResponse {
	return URLResponseClass.New()
}

func (u_ URLResponse) ExpectedContentLength() int64 {
	rv := ffi.CallMethod[int64](u_, "expectedContentLength")
	return rv
}

func (u_ URLResponse) SuggestedFilename() string {
	rv := ffi.CallMethod[string](u_, "suggestedFilename")
	return rv
}

func (u_ URLResponse) MIMEType() string {
	rv := ffi.CallMethod[string](u_, "MIMEType")
	return rv
}

func (u_ URLResponse) TextEncodingName() string {
	rv := ffi.CallMethod[string](u_, "textEncodingName")
	return rv
}

func (u_ URLResponse) URL() URL {
	rv := ffi.CallMethod[URL](u_, "URL")
	return rv
}

var HTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}

type _HTTPURLResponseClass struct {
	objc.Class
}

type IHTTPURLResponse interface {
	IURLResponse
	ValueForHTTPHeaderField(field string) string
	StatusCode() int
}

type HTTPURLResponse struct {
	URLResponse
}

func MakeHTTPURLResponse(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		URLResponse: MakeURLResponse(ptr),
	}
}

func (h_ HTTPURLResponse) InitWithURL_StatusCode_HTTPVersion_HeaderFields(url IURL, statusCode int, HTTPVersion string, headerFields map[string]string) HTTPURLResponse {
	rv := ffi.CallMethod[HTTPURLResponse](h_, "initWithURL:statusCode:HTTPVersion:headerFields:", url, statusCode, HTTPVersion, headerFields)
	rv.Autorelease()
	return rv
}

func (h_ HTTPURLResponse) InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(URL IURL, MIMEType string, length int, name string) HTTPURLResponse {
	rv := ffi.CallMethod[HTTPURLResponse](h_, "initWithURL:MIMEType:expectedContentLength:textEncodingName:", URL, MIMEType, length, name)
	rv.Autorelease()
	return rv
}

func (hc _HTTPURLResponseClass) Alloc() HTTPURLResponse {
	rv := ffi.CallMethod[HTTPURLResponse](hc, "alloc")
	return rv
}

func (h_ HTTPURLResponse) Init() HTTPURLResponse {
	rv := ffi.CallMethod[HTTPURLResponse](h_, "init")
	rv.Autorelease()
	return rv
}

func (hc _HTTPURLResponseClass) New() HTTPURLResponse {
	rv := ffi.CallMethod[HTTPURLResponse](hc, "new")
	rv.Autorelease()
	return rv
}

func NewHTTPURLResponse() HTTPURLResponse {
	return HTTPURLResponseClass.New()
}

func (h_ HTTPURLResponse) ValueForHTTPHeaderField(field string) string {
	rv := ffi.CallMethod[string](h_, "valueForHTTPHeaderField:", field)
	return rv
}

func (hc _HTTPURLResponseClass) LocalizedStringForStatusCode(statusCode int) string {
	rv := ffi.CallMethod[string](hc, "localizedStringForStatusCode:", statusCode)
	return rv
}

func (h_ HTTPURLResponse) StatusCode() int {
	rv := ffi.CallMethod[int](h_, "statusCode")
	return rv
}

var HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	objc.Class
}

type IHTTPCookie interface {
	objc.IObject
	Domain() string
	Path() string
	PortList() []Number
	Name() string
	Value() string
	Version() uint
	ExpiresDate() Date
	IsSessionOnly() bool
	IsHTTPOnly() bool
	IsSecure() bool
	SameSitePolicy() HTTPCookieStringPolicy
	Properties() map[HTTPCookiePropertyKey]objc.Object
	Comment() string
	CommentURL() URL
}

type HTTPCookie struct {
	objc.Object
}

func MakeHTTPCookie(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		Object: objc.MakeObject(ptr),
	}
}

func (h_ HTTPCookie) InitWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := ffi.CallMethod[HTTPCookie](h_, "initWithProperties:", properties)
	rv.Autorelease()
	return rv
}

func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := ffi.CallMethod[HTTPCookie](hc, "alloc")
	return rv
}

func (h_ HTTPCookie) Init() HTTPCookie {
	rv := ffi.CallMethod[HTTPCookie](h_, "init")
	rv.Autorelease()
	return rv
}

func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := ffi.CallMethod[HTTPCookie](hc, "new")
	rv.Autorelease()
	return rv
}

func NewHTTPCookie() HTTPCookie {
	return HTTPCookieClass.New()
}

func (hc _HTTPCookieClass) CookiesWithResponseHeaderFields_ForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	rv := ffi.CallMethod[[]HTTPCookie](hc, "cookiesWithResponseHeaderFields:forURL:", headerFields, URL)
	return rv
}

func (hc _HTTPCookieClass) CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := ffi.CallMethod[HTTPCookie](hc, "cookieWithProperties:", properties)
	return rv
}

func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	rv := ffi.CallMethod[map[string]string](hc, "requestHeaderFieldsWithCookies:", cookies)
	return rv
}

func (h_ HTTPCookie) Domain() string {
	rv := ffi.CallMethod[string](h_, "domain")
	return rv
}

func (h_ HTTPCookie) Path() string {
	rv := ffi.CallMethod[string](h_, "path")
	return rv
}

func (h_ HTTPCookie) PortList() []Number {
	rv := ffi.CallMethod[[]Number](h_, "portList")
	return rv
}

func (h_ HTTPCookie) Name() string {
	rv := ffi.CallMethod[string](h_, "name")
	return rv
}

func (h_ HTTPCookie) Value() string {
	rv := ffi.CallMethod[string](h_, "value")
	return rv
}

func (h_ HTTPCookie) Version() uint {
	rv := ffi.CallMethod[uint](h_, "version")
	return rv
}

func (h_ HTTPCookie) ExpiresDate() Date {
	rv := ffi.CallMethod[Date](h_, "expiresDate")
	return rv
}

func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := ffi.CallMethod[bool](h_, "isSessionOnly")
	return rv
}

func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := ffi.CallMethod[bool](h_, "isHTTPOnly")
	return rv
}

func (h_ HTTPCookie) IsSecure() bool {
	rv := ffi.CallMethod[bool](h_, "isSecure")
	return rv
}

func (h_ HTTPCookie) SameSitePolicy() HTTPCookieStringPolicy {
	rv := ffi.CallMethod[HTTPCookieStringPolicy](h_, "sameSitePolicy")
	return rv
}

func (h_ HTTPCookie) Properties() map[HTTPCookiePropertyKey]objc.Object {
	rv := ffi.CallMethod[map[HTTPCookiePropertyKey]objc.Object](h_, "properties")
	return rv
}

func (h_ HTTPCookie) Comment() string {
	rv := ffi.CallMethod[string](h_, "comment")
	return rv
}

func (h_ HTTPCookie) CommentURL() URL {
	rv := ffi.CallMethod[URL](h_, "commentURL")
	return rv
}
