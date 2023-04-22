package objc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CallMethod(t *testing.T) {
	// call method
	var o = NewObject()
	var count = CallMethod[uint](o, GetSelector("retainCount"))
	assert.Equal(t, uint(1), count)

}

func Test_AddMethod(t *testing.T) {
	class := AllocateClassPair(GetClass("NSObject"), "MyClass1", 0)
	var o = class.CreateInstance(0)
	sel := GetSelector("plus:and:")
	ok := AddMethod(class, sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2
	})
	assert.True(t, ok)
	ret := CallMethod[int](o, sel, 1, 2)
	assert.Equal(t, 3, ret)

	//replace method
	ReplaceMethod(o.GetClass(), sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2 + 10
	})
	ret = CallMethod[int](o, sel, 1, 2)
	assert.Equal(t, 13, ret)
}

func Test_AddClassMethod(t *testing.T) {
	class := AllocateClassPair(GetClass("NSObject"), "MyClass2", 0)
	//TODO: why need call class.GetMethodImplementation to make meta class exists?
	o := class.CreateInstance(0)
	o.RetainCount()
	sel := GetSelector("plus:and:")
	ok := AddClassMethod(class, sel, func(c Class, v1 int, v2 int) int {
		return v1 + v2
	})
	assert.True(t, ok)
	ret := CallMethod[int](class, sel, 4, 5)
	assert.Equal(t, 9, ret)

	// replace method
	ReplaceClassMethod(class, sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2 + 10
	})
	ret = CallMethod[int](class, sel, 4, 5)
	assert.Equal(t, 19, ret)
}
