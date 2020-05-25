package xyml_test

import (
	"gopkg.in/yaml.v3"
	"testing"
	"time"

	. "github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequireBinary(t *testing.T) {
	Convey("RequireBinary", t, func() {
		So(RequireBinary(NewBinaryNode([]byte{})), ShouldBeNil)
		So(RequireBinary(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireBinary(tmp), ShouldNotBeNil)
	})
}

func TestRequireBool(t *testing.T) {
	Convey("RequireBool", t, func() {
		So(RequireBool(NewBoolNode(true)), ShouldBeNil)
		So(RequireBool(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireBool(tmp), ShouldNotBeNil)
	})
}

func TestRequireFloat(t *testing.T) {
	Convey("RequireFloat", t, func() {
		So(RequireFloat(NewFloatNode(3, 'g', 3)), ShouldBeNil)
		So(RequireFloat(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireFloat(tmp), ShouldNotBeNil)
	})
}

func TestRequireInt(t *testing.T) {
	Convey("RequireInt", t, func() {
		So(RequireInt(NewIntNode(0, 2)), ShouldBeNil)
		So(RequireInt(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireInt(tmp), ShouldNotBeNil)
	})
}

func TestRequireMap(t *testing.T) {
	Convey("RequireMap", t, func() {
		So(RequireMap(NewMapNode(0)), ShouldBeNil)
		So(RequireMap(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireMap(tmp), ShouldNotBeNil)

		raw := &yaml.Node{Kind: yaml.DocumentNode}
		So(RequireMap(raw), ShouldNotBeNil)

		raw = &yaml.Node{Kind: yaml.AliasNode}
		So(RequireMap(raw), ShouldNotBeNil)

		raw = &yaml.Node{Kind: 69}
		So(RequireMap(raw), ShouldNotBeNil)
	})
}

func TestRequireNilType(t *testing.T) {
	Convey("RequireBool", t, func() {
		So(RequireNilType(NewNullNode()), ShouldBeNil)
		So(RequireNilType(NewStringNode("")), ShouldNotBeNil)

		tmp := NewMapNode(0)
		tmp.Line = 1
		tmp.Column = 2
		So(RequireNilType(tmp), ShouldNotBeNil)
	})
}

func TestRequireOrderedMap(t *testing.T) {
	Convey("RequireOrderedMap", t, func() {
		So(RequireOrderedMap(NewOrderedMapNode(0)), ShouldBeNil)
		So(RequireOrderedMap(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireOrderedMap(tmp), ShouldNotBeNil)
	})
}

func TestRequirePairs(t *testing.T) {
	Convey("RequirePairs", t, func() {
		So(RequirePairs(NewPairsNode(0)), ShouldBeNil)
		So(RequirePairs(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequirePairs(tmp), ShouldNotBeNil)
	})
}

func TestRequireScalar(t *testing.T) {
	Convey("RequireScalar", t, func() {
		So(RequireScalar(NewIntNode(0, 10)), ShouldBeNil)
		So(RequireScalar(NewMapNode(0)), ShouldNotBeNil)

		tmp := NewSequenceNode(0)
		tmp.Line = 1
		tmp.Column = 2
		So(RequireScalar(tmp), ShouldNotBeNil)
	})
}

func TestRequireSequence(t *testing.T) {
	Convey("RequireSequence", t, func() {
		So(RequireSequence(NewSequenceNode(0)), ShouldBeNil)
		So(RequireSequence(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireSequence(tmp), ShouldNotBeNil)
	})
}

func TestRequireString(t *testing.T) {
	Convey("RequireString", t, func() {
		So(RequireString(NewStringNode("0")), ShouldBeNil)
		So(RequireString(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireString(tmp), ShouldNotBeNil)
	})
}

func TestRequireTimestamp(t *testing.T) {
	Convey("RequireTimestamp", t, func() {
		So(RequireTimestamp(NewTimestampNode(time.Now())), ShouldBeNil)
		So(RequireTimestamp(NewNullNode()), ShouldNotBeNil)

		tmp := NewNullNode()
		tmp.Line = 1
		tmp.Column = 2
		So(RequireTimestamp(tmp), ShouldNotBeNil)
	})
}
