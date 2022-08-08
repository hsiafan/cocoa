// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var OrthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}

type _OrthographyClass struct {
	objc.Class
}

type IOrthography interface {
	objc.IObject
	DominantLanguageForScript(script string) string
	LanguagesForScript(script string) []string
	LanguageMap() map[string][]string
	DominantLanguage() string
	DominantScript() string
	AllScripts() []string
	AllLanguages() []string
}

type Orthography struct {
	objc.Object
}

func MakeOrthography(ptr unsafe.Pointer) Orthography {
	return Orthography{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OrthographyClass) DefaultOrthographyForLanguage(language string) Orthography {
	rv := ffi.CallMethod[Orthography](oc, "defaultOrthographyForLanguage:", language)
	return rv
}

func (o_ Orthography) InitWithDominantScript_LanguageMap(script string, _map map[string][]string) Orthography {
	rv := ffi.CallMethod[Orthography](o_, "initWithDominantScript:languageMap:", script, _map)
	return rv
}

func (oc _OrthographyClass) OrthographyWithDominantScript_LanguageMap(script string, _map map[string][]string) Orthography {
	rv := ffi.CallMethod[Orthography](oc, "orthographyWithDominantScript:languageMap:", script, _map)
	return rv
}

func (oc _OrthographyClass) Alloc() Orthography {
	rv := ffi.CallMethod[Orthography](oc, "alloc")
	return rv
}

func (o_ Orthography) Init() Orthography {
	rv := ffi.CallMethod[Orthography](o_, "init")
	return rv
}

func (oc _OrthographyClass) New() Orthography {
	rv := ffi.CallMethod[Orthography](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOrthography() Orthography {
	return OrthographyClass.New()
}

func (o_ Orthography) DominantLanguageForScript(script string) string {
	rv := ffi.CallMethod[string](o_, "dominantLanguageForScript:", script)
	return rv
}

func (o_ Orthography) LanguagesForScript(script string) []string {
	rv := ffi.CallMethod[[]string](o_, "languagesForScript:", script)
	return rv
}

func (o_ Orthography) LanguageMap() map[string][]string {
	rv := ffi.CallMethod[map[string][]string](o_, "languageMap")
	return rv
}

func (o_ Orthography) DominantLanguage() string {
	rv := ffi.CallMethod[string](o_, "dominantLanguage")
	return rv
}

func (o_ Orthography) DominantScript() string {
	rv := ffi.CallMethod[string](o_, "dominantScript")
	return rv
}

func (o_ Orthography) AllScripts() []string {
	rv := ffi.CallMethod[[]string](o_, "allScripts")
	return rv
}

func (o_ Orthography) AllLanguages() []string {
	rv := ffi.CallMethod[[]string](o_, "allLanguages")
	return rv
}
