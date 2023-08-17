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

type SoundDelegateWrapper struct {
	objc.Object
}

func (s_ SoundDelegateWrapper) Sound_DidFinishPlaying(sound ISound, flag bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sound:didFinishPlaying:"), objc.ExtractPtr(sound), flag)
}
