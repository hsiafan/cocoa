package foundation

import (
	"bytes"
	"testing"
)

func TestNewMutableString(t *testing.T) {
	ms := NewMutableString()
	ms.AppendString("test")
	if ms.ToString() != "test" {
		t.Fail()
	}

	ms = NewMutableStringWithCapacity(10)
	ms.AppendString("test")
	if ms.ToString() != "test" {
		t.Fail()
	}

	ms = NewMutableStringWithString("test")
	if ms.ToString() != "test" {
		t.Fail()
	}
}

func TestNewMutableData(t *testing.T) {
	data := []byte{1, 2, 3, 4}
	md := NewMutableData()
	md.AppendData(data)
	if !bytes.Equal(data, md.ToBytes()) {
		t.Fail()
	}

	md = NewMutableDataWithCapacity(10)
	md.AppendData(data)
	if !bytes.Equal(data, md.ToBytes()) {
		t.Fail()
	}

	md = NewMutableDataWithData(data)
	if !bytes.Equal(data, md.ToBytes()) {
		t.Fail()
	}
}
