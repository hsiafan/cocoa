package ffi

import (
	"sync"
)

// Protocol meta data
type ProtocolInfo struct {
	Name string // the objc protocol name
}

var _protocolRegistry = protocolRegistry{
	typeMap: map[string]ProtocolInfo{},
}

// protocolRegistry register protocol meta data
type protocolRegistry struct {
	typeMap map[string]ProtocolInfo // go protocol interface type name to ProtocolInfo
	lock    sync.RWMutex
}

func (p *protocolRegistry) register(goTypeName string, pi ProtocolInfo) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.typeMap[goTypeName] = pi
}

func (p *protocolRegistry) get(goTypeName string) (pi ProtocolInfo, ok bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	pi, ok = p.typeMap[goTypeName]
	return
}

// RegisterProtocol registers a protocol meta data
func RegisterProtocol(goTypeName string, pi ProtocolInfo) {
	_protocolRegistry.register(goTypeName, pi)
}

// GetProtocolInfo return a protocol meta data
func GetProtocolInfo(goTypeName string) (pi ProtocolInfo, ok bool) {
	return _protocolRegistry.get(goTypeName)
}
