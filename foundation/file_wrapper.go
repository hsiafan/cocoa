// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	objc.Class
}

type IFileWrapper interface {
	objc.IObject
	// deprecated
	InitWithPath(path string) objc.Object
	// deprecated
	InitSymbolicLinkWithDestination(path string) objc.Object
	AddFileWrapper(child IFileWrapper) string
	RemoveFileWrapper(child IFileWrapper)
	// deprecated
	AddFileWithPath(path string) string
	AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string
	// deprecated
	AddSymbolicLinkWithDestination_PreferredFilename(path string, filename string) string
	KeyForFileWrapper(child IFileWrapper) string
	// deprecated
	SymbolicLinkDestination() string
	// deprecated
	NeedsToBeUpdatedFromPath(path string) bool
	MatchesContentsOfURL(url IURL) bool
	// deprecated
	UpdateFromPath(path string) bool
	ReadFromURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) bool
	// deprecated
	WriteToFile_Atomically_UpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool
	WriteToURL_Options_OriginalContentsURL_Error(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool
	IsRegularFile() bool
	IsDirectory() bool
	IsSymbolicLink() bool
	FileWrappers() map[string]FileWrapper
	SymbolicLinkDestinationURL() URL
	SerializedRepresentation() []byte
	Filename() string
	SetFilename(value string)
	PreferredFilename() string
	SetPreferredFilename(value string)
	FileAttributes() map[string]objc.Object
	SetFileAttributes(value map[string]objc.IObject)
	RegularFileContents() []byte
}

type FileWrapper struct {
	objc.Object
}

func MakeFileWrapper(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		Object: objc.MakeObject(ptr),
	}
}

func (f_ FileWrapper) InitWithURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "initWithURL:options:error:", url, options, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "initDirectoryWithFileWrappers:", childrenByPreferredName)
	return rv
}

func (f_ FileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "initRegularFileWithContents:", contents)
	return rv
}

func (f_ FileWrapper) InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "initSymbolicLinkWithDestinationURL:", url)
	return rv
}

func (f_ FileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "initWithSerializedRepresentation:", serializeRepresentation)
	return rv
}

func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := ffi.CallMethod[FileWrapper](fc, "alloc")
	return rv
}

func (f_ FileWrapper) Init() FileWrapper {
	rv := ffi.CallMethod[FileWrapper](f_, "init")
	return rv
}

func (fc _FileWrapperClass) New() FileWrapper {
	rv := ffi.CallMethod[FileWrapper](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFileWrapper() FileWrapper {
	return FileWrapperClass.New()
}

// deprecated
func (f_ FileWrapper) InitWithPath(path string) objc.Object {
	rv := ffi.CallMethod[objc.Object](f_, "initWithPath:", path)
	return rv
}

// deprecated
func (f_ FileWrapper) InitSymbolicLinkWithDestination(path string) objc.Object {
	rv := ffi.CallMethod[objc.Object](f_, "initSymbolicLinkWithDestination:", path)
	return rv
}

func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) string {
	rv := ffi.CallMethod[string](f_, "addFileWrapper:", child)
	return rv
}

func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	ffi.CallMethod[ffi.Void](f_, "removeFileWrapper:", child)
}

// deprecated
func (f_ FileWrapper) AddFileWithPath(path string) string {
	rv := ffi.CallMethod[string](f_, "addFileWithPath:", path)
	return rv
}

func (f_ FileWrapper) AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string {
	rv := ffi.CallMethod[string](f_, "addRegularFileWithContents:preferredFilename:", data, fileName)
	return rv
}

// deprecated
func (f_ FileWrapper) AddSymbolicLinkWithDestination_PreferredFilename(path string, filename string) string {
	rv := ffi.CallMethod[string](f_, "addSymbolicLinkWithDestination:preferredFilename:", path, filename)
	return rv
}

func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) string {
	rv := ffi.CallMethod[string](f_, "keyForFileWrapper:", child)
	return rv
}

// deprecated
func (f_ FileWrapper) SymbolicLinkDestination() string {
	rv := ffi.CallMethod[string](f_, "symbolicLinkDestination")
	return rv
}

// deprecated
func (f_ FileWrapper) NeedsToBeUpdatedFromPath(path string) bool {
	rv := ffi.CallMethod[bool](f_, "needsToBeUpdatedFromPath:", path)
	return rv
}

func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := ffi.CallMethod[bool](f_, "matchesContentsOfURL:", url)
	return rv
}

// deprecated
func (f_ FileWrapper) UpdateFromPath(path string) bool {
	rv := ffi.CallMethod[bool](f_, "updateFromPath:", path)
	return rv
}

func (f_ FileWrapper) ReadFromURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) bool {
	rv := ffi.CallMethod[bool](f_, "readFromURL:options:error:", url, options, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (f_ FileWrapper) WriteToFile_Atomically_UpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool {
	rv := ffi.CallMethod[bool](f_, "writeToFile:atomically:updateFilenames:", path, atomicFlag, updateFilenamesFlag)
	return rv
}

func (f_ FileWrapper) WriteToURL_Options_OriginalContentsURL_Error(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool {
	rv := ffi.CallMethod[bool](f_, "writeToURL:options:originalContentsURL:error:", url, options, originalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) IsRegularFile() bool {
	rv := ffi.CallMethod[bool](f_, "isRegularFile")
	return rv
}

func (f_ FileWrapper) IsDirectory() bool {
	rv := ffi.CallMethod[bool](f_, "isDirectory")
	return rv
}

func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := ffi.CallMethod[bool](f_, "isSymbolicLink")
	return rv
}

func (f_ FileWrapper) FileWrappers() map[string]FileWrapper {
	rv := ffi.CallMethod[map[string]FileWrapper](f_, "fileWrappers")
	return rv
}

func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := ffi.CallMethod[URL](f_, "symbolicLinkDestinationURL")
	return rv
}

func (f_ FileWrapper) SerializedRepresentation() []byte {
	rv := ffi.CallMethod[[]byte](f_, "serializedRepresentation")
	return rv
}

func (f_ FileWrapper) Filename() string {
	rv := ffi.CallMethod[string](f_, "filename")
	return rv
}

func (f_ FileWrapper) SetFilename(value string) {
	ffi.CallMethod[ffi.Void](f_, "setFilename:", value)
}

func (f_ FileWrapper) PreferredFilename() string {
	rv := ffi.CallMethod[string](f_, "preferredFilename")
	return rv
}

func (f_ FileWrapper) SetPreferredFilename(value string) {
	ffi.CallMethod[ffi.Void](f_, "setPreferredFilename:", value)
}

func (f_ FileWrapper) FileAttributes() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](f_, "fileAttributes")
	return rv
}

func (f_ FileWrapper) SetFileAttributes(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "setFileAttributes:", value)
}

func (f_ FileWrapper) RegularFileContents() []byte {
	rv := ffi.CallMethod[[]byte](f_, "regularFileContents")
	return rv
}
