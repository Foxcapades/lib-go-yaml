package xyml_test

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMapAppend(t *testing.T) {
	Convey("MapAppend", t, func() {
		node := xyml.NewStringNode("")
		So(xyml.MapAppend(node, "a", "b"), ShouldNotBeNil)

		node = xyml.NewMapNode(1)
		So(xyml.MapAppend(node, "a", "b"), ShouldBeNil)
		So(len(node.Content), ShouldEqual, 2)
		So(node.Content[0].Value, ShouldEqual, "a")
		So(node.Content[1].Value, ShouldEqual, "b")

		node = xyml.NewMapNode(1)
		So(xyml.MapAppend(node, struct{}{}, "foo"), ShouldNotBeNil)

		node = xyml.NewMapNode(1)
		So(xyml.MapAppend(node, "foo", struct{}{}), ShouldNotBeNil)
	})
}

func TestSequenceAppend(t *testing.T) {
	Convey("SequenceAppend", t, func() {
		node := xyml.NewStringNode("")
		So(xyml.SequenceAppend(node, "a"), ShouldNotBeNil)

		node = xyml.NewSequenceNode(1)
		So(xyml.SequenceAppend(node, "a"), ShouldBeNil)
		So(len(node.Content), ShouldEqual, 1)
		So(node.Content[0].Value, ShouldEqual, "a")

		node = xyml.NewSequenceNode(1)
		So(xyml.SequenceAppend(node, struct{}{}), ShouldNotBeNil)
	})
}

func TestMapAppendNode(t *testing.T) {
	Convey("MapAppendNode", t, func() {
		So(xyml.MapAppendNode(xyml.NewStringNode("hi"), xyml.NewNullNode(),
			xyml.NewNullNode()), ShouldNotBeNil)
	})
}

func TestSequenceAppendNode(t *testing.T) {
	Convey("SequenceAppendNode", t, func() {
		So(xyml.SequenceAppendNode(xyml.NewStringNode("hi"), xyml.NewNullNode()),
			ShouldNotBeNil)
	})
}