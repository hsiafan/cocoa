// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type SoundDelegate interface {
	ImplementsSound_DidFinishPlaying() bool
	// optional
	Sound_DidFinishPlaying(sound Sound, flag bool)
}

func WrapSoundDelegate(v SoundDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSoundDelegate", v)
}

type SoundDelegateBase struct {
}

func (p *SoundDelegateBase) ImplementsSound_DidFinishPlaying() bool {
	return false
}

func (p *SoundDelegateBase) Sound_DidFinishPlaying(sound Sound, flag bool) {
	panic("unimpemented protocol method")
}

type SoundDelegateCreator struct {
	className string
	class     objc.Class
}

func NewSoundDelegateCreator(name string) *SoundDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &SoundDelegateCreator{className: name, class: class}
}

func (c *SoundDelegateCreator) SetSound_DidFinishPlaying(handle func(o objc.Object, sound Sound, flag bool)) {
	objc.AddMethod(c.class, objc.SEL("sound:didFinishPlaying:"), handle)
}

func (c *SoundDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
