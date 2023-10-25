package objc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsgSend(t *testing.T) {
	o := NewObject()
	o.RetainCount()
	count := MsgSend[uint](o, SEL("retainCount"))
	assert.Equal(t, uint(1), count)
}
