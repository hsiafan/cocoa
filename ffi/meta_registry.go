package ffi

import (
	"reflect"
	"sync"
)

// Protocol meta data
type ProtocolInfo struct {
	Name string // the objc protocol name
}

var _protocolRegistry = protocolRegistry{
	data: map[reflect.Type]ProtocolInfo{},
}

// protocolRegistry register protocol meta data
type protocolRegistry struct {
	data map[reflect.Type]ProtocolInfo // go protocol interface type to ProtocolInfo
	lock sync.RWMutex
}

func (p *protocolRegistry) register(pt reflect.Type, pi ProtocolInfo) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.data[pt] = pi
}

func (p *protocolRegistry) get(pt reflect.Type) (pi ProtocolInfo, ok bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	pi, ok = p.data[pt]
	return
}

// RegisterProtocol registers a protocol meta data
func RegisterProtocol[T any](pi ProtocolInfo) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	_protocolRegistry.register(t, pi)
}

// GetProtocolInfo return a protocol meta data
func GetProtocolInfo[T any]() (pi ProtocolInfo, ok bool) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return _protocolRegistry.get(t)
}
