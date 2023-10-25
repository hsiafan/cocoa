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
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &SoundDelegateCreator{className: name, class: class}
}

func (c *SoundDelegateCreator) SetSound_DidFinishPlaying(handle func(o objc.ProtocolBase, sound Sound, flag bool)) {
	objc.AddMethod(c.class, objc.SEL("sound:didFinishPlaying:"), handle)
}

func (c *SoundDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
