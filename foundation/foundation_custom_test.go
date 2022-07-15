package foundation

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestJSONObjectWithData(t *testing.T) {
	json := "[1, 2, 3]"
	o, err := JSONObjectWithData([]byte(json), JSONReadingFragmentsAllowed)
	assert.NoError(t, err)
	assert.False(t, o.IsNil())
	array := MakeArray(o.Ptr())
	assert.Equal(t, uint(3), array.Count())

	json = "[1, 2,"
	o, err = JSONObjectWithData([]byte(json), JSONReadingFragmentsAllowed)
	assert.Error(t, err)
	assert.True(t, o.IsNil())
}

func TestDataWithJSONObject(t *testing.T) {
	array := ArrayOf([]string{"1", "2"})
	d, err := DataWithJSONObject(array, JSONWritingFragmentsAllowed)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(d), "1"))
}
