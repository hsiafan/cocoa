// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	objc.Class
}

type ICharacterSet interface {
	objc.IObject
	CharacterIsMember(aCharacter unichar) bool
	IsSupersetOfSet(theOtherSet ICharacterSet) bool
	LongCharacterIsMember(theLongChar uint32) bool
	BitmapRepresentation() []byte
	InvertedSet() CharacterSet
}

type CharacterSet struct {
	objc.Object
}

func MakeCharacterSet(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "alloc")
	return rv
}

func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCharacterSet() CharacterSet {
	return CharacterSetClass.New()
}

func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, "init")
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithCharactersInString(aString string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "characterSetWithCharactersInString:", aString)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithRange(aRange Range) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "characterSetWithRange:", aRange)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "characterSetWithBitmapRepresentation:", data)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithContentsOfFile(fName string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "characterSetWithContentsOfFile:", fName)
	return rv
}

func (c_ CharacterSet) CharacterIsMember(aCharacter unichar) bool {
	rv := objc.CallMethod[bool](c_, "characterIsMember:", aCharacter)
	return rv
}

func (c_ CharacterSet) IsSupersetOfSet(theOtherSet ICharacterSet) bool {
	rv := objc.CallMethod[bool](c_, "isSupersetOfSet:", theOtherSet)
	return rv
}

func (c_ CharacterSet) LongCharacterIsMember(theLongChar uint32) bool {
	rv := objc.CallMethod[bool](c_, "longCharacterIsMember:", theLongChar)
	return rv
}

func (cc _CharacterSetClass) AlphanumericCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "alphanumericCharacterSet")
	return rv
}

func (cc _CharacterSetClass) CapitalizedLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "capitalizedLetterCharacterSet")
	return rv
}

func (cc _CharacterSetClass) ControlCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "controlCharacterSet")
	return rv
}

func (cc _CharacterSetClass) DecimalDigitCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "decimalDigitCharacterSet")
	return rv
}

func (cc _CharacterSetClass) DecomposableCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "decomposableCharacterSet")
	return rv
}

func (cc _CharacterSetClass) IllegalCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "illegalCharacterSet")
	return rv
}

func (cc _CharacterSetClass) LetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "letterCharacterSet")
	return rv
}

func (cc _CharacterSetClass) LowercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "lowercaseLetterCharacterSet")
	return rv
}

func (cc _CharacterSetClass) NewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "newlineCharacterSet")
	rv.Autorelease()
	return rv
}

func (cc _CharacterSetClass) NonBaseCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "nonBaseCharacterSet")
	return rv
}

func (cc _CharacterSetClass) PunctuationCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "punctuationCharacterSet")
	return rv
}

func (cc _CharacterSetClass) SymbolCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "symbolCharacterSet")
	return rv
}

func (cc _CharacterSetClass) UppercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "uppercaseLetterCharacterSet")
	return rv
}

func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "whitespaceAndNewlineCharacterSet")
	return rv
}

func (cc _CharacterSetClass) WhitespaceCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "whitespaceCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLFragmentAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLFragmentAllowedCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLHostAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLHostAllowedCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLPasswordAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLPasswordAllowedCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLPathAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLPathAllowedCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLQueryAllowedCharacterSet")
	return rv
}

func (cc _CharacterSetClass) URLUserAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, "URLUserAllowedCharacterSet")
	return rv
}

func (c_ CharacterSet) BitmapRepresentation() []byte {
	rv := objc.CallMethod[[]byte](c_, "bitmapRepresentation")
	return rv
}

func (c_ CharacterSet) InvertedSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, "invertedSet")
	return rv
}
