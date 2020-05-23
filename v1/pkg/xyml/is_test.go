package xyml_test

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestIsBinary(t *testing.T) {
	Convey("IsBinary", t, func() {
		So(xyml.IsBinary(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsBinary(xyml.NewBinaryNode([]byte{})), ShouldBeTrue)
	})
}

func TestIsBool(t *testing.T) {
	Convey("IsBool", t, func() {
		So(xyml.IsBool(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsBool(xyml.NewBoolNode(true)), ShouldBeTrue)
	})
}

func TestIsFloat(t *testing.T) {
	Convey("IsFloat", t, func() {
		So(xyml.IsFloat(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsFloat(xyml.NewFloatNode(3, 'g', 3)), ShouldBeTrue)
	})
}

func TestIsInt(t *testing.T) {
	Convey("IsInt", t, func() {
		So(xyml.IsInt(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsInt(xyml.NewIntNode(3, 10)), ShouldBeTrue)
	})
}

func TestIsMap(t *testing.T) {
	Convey("IsMap", t, func() {
		So(xyml.IsMap(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsMap(xyml.NewOrderedMapNode(0)), ShouldBeTrue)
	})
}

func TestIsNilType(t *testing.T) {
	Convey("IsNilType", t, func() {
		So(xyml.IsNilType(xyml.NewStringNode("hi")), ShouldBeFalse)
		So(xyml.IsNilType(xyml.NewNullNode()), ShouldBeTrue)
	})
}

func TestIsOrderedMap(t *testing.T) {
	Convey("IsOrderedMap", t, func() {
		So(xyml.IsOrderedMap(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsOrderedMap(xyml.NewOrderedMapNode(0)), ShouldBeTrue)
	})
}

func TestIsPairs(t *testing.T) {
	Convey("IsPairs()", t, func() {
		So(xyml.IsPairs(xyml.NewIntNode(0, 10)), ShouldBeFalse)
		So(xyml.IsPairs(xyml.NewPairsNode(0)), ShouldBeTrue)
	})
}

func TestIsScalar(t *testing.T) {
	Convey("IsScalar()", t, func() {
		So(xyml.IsScalar(xyml.NewPairsNode(0)), ShouldBeFalse)
		So(xyml.IsScalar(xyml.NewIntNode(0, 10)), ShouldBeTrue)
	})
}

func TestIsSequence(t *testing.T) {
	Convey("IsSequence()", t, func() {
		So(xyml.IsSequence(xyml.NewStringNode("the")), ShouldBeFalse)
		So(xyml.IsSequence(xyml.NewSequenceNode(0)), ShouldBeTrue)
	})
}

func TestIsString(t *testing.T) {
	Convey("IsString", t, func() {
		So(xyml.IsString(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsString(xyml.NewStringNode("0")), ShouldBeTrue)
	})
}

func TestIsTimestamp(t *testing.T) {
	Convey("IsTimestamp", t, func() {
		So(xyml.IsTimestamp(xyml.NewNullNode()), ShouldBeFalse)
		So(xyml.IsTimestamp(xyml.NewTimestampNode(time.Now())), ShouldBeTrue)
	})
}



