package main

import (
	"testing"

	"github.com/hsiafan/cocoa/generate/data"
	"github.com/stretchr/testify/assert"
)

func Test_getClassGen(t *testing.T) {
	ti := loadOneTypeInfo("AppKit", "NSAlert").(*data.Class)
	cg := getClassGen(ti)
	cg.Init()
	assert.Equal(t, "NSAlert", cg.Type.Name)
	assert.Equal(t, "NSObject", cg.Parent.Type.Name)
}

func Test_getProtolcolGen(t *testing.T) {
	ti := loadOneTypeInfo("AppKit", "NSApplicationDelegate").(*data.Protocol)
	cg := getProtocolGen(ti)
	cg.Init()
	assert.Equal(t, "NSApplicationDelegate", cg.Type.Name)
	assert.Equal(t, 0, len(cg.Parents))
}
