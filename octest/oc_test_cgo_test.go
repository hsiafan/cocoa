package octest

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/quartzcore"
)

func Test_basicTypes(t *testing.T) {
	compareBasicType[int32, CInt](t)
	compareBasicType[int, Integer](t)
	compareBasicType[int64, CLong](t)
	compareBasicType[bool, CBOOL](t)
	compareBasicType[bool, CBool](t)
}

func compareBasicType[X any, Y any](t *testing.T) {
	var x X
	var y Y
	if unsafe.Sizeof(x) != unsafe.Sizeof(y) {
		t.Logf("type not match, %T %T", x, y)
		t.FailNow()
	}
}

func Test_structs(t *testing.T) {
	compareStruct[AffineTransform, coregraphics.AffineTransform](t)
	compareStruct[Point, coregraphics.Point](t)
	compareStruct[Size, coregraphics.Size](t)
	compareStruct[Rect, coregraphics.Rect](t)

	compareStruct[EdgeInsets, foundation.EdgeInsets](t)
	compareStruct[Range, foundation.Range](t)
	compareStruct[AffineTransformStruct, foundation.AffineTransformStruct](t)
	compareStruct[Decimal, foundation.Decimal](t)

	compareStruct[DirectionalEdgeInsets, appkit.DirectionalEdgeInsets](t)

	compareStruct[Transform3D, quartzcore.Transform3D](t)
	compareStruct[FrameRateRange, quartzcore.FrameRateRange](t)
}

func compareStruct[X any, Y any](t *testing.T) {
	var x X
	var y Y
	if unsafe.Sizeof(x) != unsafe.Sizeof(y) {
		t.Logf("struct size unmatch, %T:%v, %T:%v", x, unsafe.Sizeof(x), y, unsafe.Sizeof(y))
		t.FailNow()
	}

	xt := reflect.TypeOf(x)
	yt := reflect.TypeOf(y)
	if xt.NumField() != yt.NumField() {
		t.Logf("struct field count unmatch, %T %T", x, y)
		t.FailNow()
	}

	for i := 0; i < xt.NumField(); i++ {
		xf := xt.Field(i)
		yf := xt.Field(i)
		if !strings.EqualFold(xf.Name, yf.Name) {
			t.Logf("struct field name unmatch, %T:%s %T:%s", x, xf.Name, y, yf.Name)
			t.FailNow()
		}
		if xf.Type.Kind() != yf.Type.Kind() {
			t.Logf("struct field kind unmatch, %T:%s %T:%s", x, xf.Name, y, yf.Name)
			t.FailNow()
		}
	}

}
