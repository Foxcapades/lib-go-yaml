package xyml_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBinaryNode(t *testing.T) {
	Convey("NewBinaryNode", t, func() {
		test := xyml.NewBinaryNode([]byte("hi"))
		So(test.Kind, ShouldEqual, yaml.ScalarNode)
		So(test.Tag, ShouldEqual, xyml.TagBinary)
		So(test.Value, ShouldResemble, base64.StdEncoding.EncodeToString([]byte("hi")))
	})
}

func TestNewBoolNode(t *testing.T) {
	Convey("NewBoolNode", t, func() {
		test := xyml.NewBoolNode(false)
		So(test.Kind, ShouldEqual, yaml.ScalarNode)
		So(test.Tag, ShouldEqual, xyml.TagBool)
		So(test.Value, ShouldEqual, "false")
	})
}

func TestNewStringNode(t *testing.T) {
	Convey("NewStringNode", t, func() {
		test := xyml.NewStringNode("nope")
		So(test.Kind, ShouldEqual, yaml.ScalarNode)
		So(test.Tag, ShouldEqual, xyml.TagString)
		So(test.Value, ShouldEqual, "nope")
	})
}

func TestNewTimestampNode(t *testing.T) {
	Convey("NewTimestampNode", t, func() {
		test := xyml.NewTimestampNode(time.Date(1988, 9, 18, 3, 23, 58, 123456789, time.UTC))
		So(test.Kind, ShouldEqual, yaml.ScalarNode)
		So(test.Tag, ShouldEqual, xyml.TagTimestamp)
		So(test.Value, ShouldEqual, "1988-09-18T03:23:58.123456789Z")
	})
}

func TestNewOrderedMapNode(t *testing.T) {
	Convey("NewOrderedMapNode", t, func() {
		test := xyml.NewOrderedMapNode(0)
		So(test.Kind, ShouldEqual, yaml.SequenceNode)
		So(test.Tag, ShouldEqual, xyml.TagOrderedMap)
	})
}

func TestNewPairsNode(t *testing.T) {
	Convey("NewPairsNode", t, func() {
		test := xyml.NewPairsNode(0)
		So(test.Kind, ShouldEqual, yaml.SequenceNode)
		So(test.Tag, ShouldEqual, xyml.TagPairs)
	})
}

func TestNewSetNode(t *testing.T) {
	Convey("NewSetNode", t, func() {
		test := xyml.NewSetNode(0)
		So(test.Kind, ShouldEqual, yaml.SequenceNode)
		So(test.Tag, ShouldEqual, xyml.TagSet)
	})
}
