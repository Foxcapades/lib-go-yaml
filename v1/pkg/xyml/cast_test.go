package xyml_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"

	"gopkg.in/yaml.v3"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScalarToYamlNode(t *testing.T) {
	Convey("ScalarToYamlNode", t, func() {
		anchor := "hi"
		tests := []struct {
			msg            string
			raw            interface{}
			expVal, expTag string
		}{
			{"Given a string value", "hi", "hi", xyml.TagString},
			{"Given an untyped int value", 13, "13", xyml.TagInt},
			{"Given an int8 value", int8(27), "27", xyml.TagInt},
			{"Given an int16 value", int16(27), "27", xyml.TagInt},
			{"Given an int32 value", int32(27), "27", xyml.TagInt},
			{"Given an int64 value", int64(27), "27", xyml.TagInt},
			{"Given an untyped uint value", uint(13), "13", xyml.TagInt},
			{"Given an uint8 value", uint8(27), "27", xyml.TagInt},
			{"Given an uint16 value", uint16(27), "27", xyml.TagInt},
			{"Given an uint32 value", uint32(27), "27", xyml.TagInt},
			{"Given an uint64 value", uint64(27), "27", xyml.TagInt},
			{"Given a bool value", true, "true", xyml.TagBool},
			{"Given a float32 value", float32(3), "3", xyml.TagFloat},
			{"Given a float64 value", float64(4.4), "4.4", xyml.TagFloat},
			{"Given a nil value", nil, "", xyml.TagNil},
			{"Given a typed nil value", (*string)(nil), "", xyml.TagNil},
			{"Given a pointer value", &anchor, anchor, xyml.TagString},
		}

		for _, test := range tests {
			Convey(test.msg, func() {
				val, err := xyml.ScalarToYamlNode(test.raw)

				So(err, ShouldBeNil)
				So(val.Kind, ShouldEqual, yaml.ScalarNode)
				So(val.Tag, ShouldEqual, test.expTag)
				So(val.Value, ShouldEqual, test.expVal)
			})
		}

		Convey("Given a non-scalar value", func() {
			val, err := xyml.ScalarToYamlNode(struct{}{})

			So(val, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestMapToYamlNode(t *testing.T) {
	Convey("MapToYamlNode", t, func() {
		Convey("scalar keys", func() {
			var map1 map[string]string

			val, err := xyml.MapToYamlNode(map1)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(val.Content, ShouldBeEmpty)

			map2 := map[interface{}]string{struct{}{}: "hi"}
			val, err = xyml.MapToYamlNode(map2)

			So(err, ShouldNotBeNil)
			So(val, ShouldBeNil)

			map3 := map[int]interface{}{3: struct{}{}}
			val, err = xyml.MapToYamlNode(map3)

			So(err, ShouldNotBeNil)
			So(val, ShouldBeNil)

			map4 := map[int][]string{3: {"hi"}}
			val, err = xyml.MapToYamlNode(map4)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(len(val.Content), ShouldEqual, 2)
			So(val.Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[0].Tag, ShouldEqual, xyml.TagInt)
			So(val.Content[0].Value, ShouldEqual, "3")
			So(val.Content[1].Kind, ShouldEqual, yaml.SequenceNode)
			So(val.Content[1].Tag, ShouldEqual, xyml.TagSequence)
			So(len(val.Content[1].Content), ShouldEqual, 1)
			So(val.Content[1].Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[1].Content[0].Tag, ShouldEqual, xyml.TagString)
			So(val.Content[1].Content[0].Value, ShouldEqual, "hi")

			map5 := map[int][]interface{}{3: {struct{}{}}}
			val, err = xyml.MapToYamlNode(map5)

			So(err, ShouldNotBeNil)
			So(val, ShouldBeNil)

			str := "hi"
			map6 := map[int][]*string{3: {nil, &str}}
			val, err = xyml.MapToYamlNode(map6)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(len(val.Content), ShouldEqual, 2)
			So(val.Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[0].Tag, ShouldEqual, xyml.TagInt)
			So(val.Content[0].Value, ShouldEqual, "3")
			So(val.Content[1].Kind, ShouldEqual, yaml.SequenceNode)
			So(val.Content[1].Tag, ShouldEqual, xyml.TagSequence)
			So(len(val.Content[1].Content), ShouldEqual, 2)
			So(val.Content[1].Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[1].Content[0].Tag, ShouldEqual, xyml.TagNil)
			So(val.Content[1].Content[0].Value, ShouldEqual, "")
			So(val.Content[1].Content[1].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[1].Content[1].Tag, ShouldEqual, xyml.TagString)
			So(val.Content[1].Content[1].Value, ShouldEqual, "hi")

			val, err = xyml.MapToYamlNode(nil)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(val.Content, ShouldBeEmpty)

			map7 := map[string]string{}
			val, err = xyml.MapToYamlNode(&map7)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(val.Content, ShouldBeEmpty)

			map8 := map[string]map[string]string{"hi": {"bye": "farewell"}}
			val, err = xyml.MapToYamlNode(&map8)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(len(val.Content), ShouldEqual, 2)
			So(val.Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[0].Tag, ShouldEqual, xyml.TagString)
			So(val.Content[0].Value, ShouldEqual, "hi")
			So(val.Content[1].Kind, ShouldEqual, yaml.MappingNode)
			So(val.Content[1].Tag, ShouldEqual, xyml.TagMap)
			So(len(val.Content[1].Content), ShouldEqual, 2)
			So(val.Content[1].Content[0].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[1].Content[0].Tag, ShouldEqual, xyml.TagString)
			So(val.Content[1].Content[0].Value, ShouldEqual, "bye")
			So(val.Content[1].Content[1].Kind, ShouldEqual, yaml.ScalarNode)
			So(val.Content[1].Content[1].Tag, ShouldEqual, xyml.TagString)
			So(val.Content[1].Content[1].Value, ShouldEqual, "farewell")

			var map9 interface{} = (*string)(nil)
			val, err = xyml.MapToYamlNode(map9)

			So(err, ShouldBeNil)
			So(val.Kind, ShouldEqual, yaml.MappingNode)
			So(val.Tag, ShouldEqual, xyml.TagMap)
			So(val.Content, ShouldBeEmpty)

			map10 := "hi"
			val, err = xyml.MapToYamlNode(map10)

			So(err, ShouldNotBeNil)
			So(val, ShouldBeNil)
		})
	})
}

func TestSliceToYamlNode(t *testing.T) {
	Convey("SliceToYamlNode", t, func() {
		var slice1 []string

		val, err := xyml.SliceToYamlNode(slice1)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.SequenceNode)
		So(val.Tag, ShouldEqual, xyml.TagSequence)
		So(val.Content, ShouldBeEmpty)

		slice2 := []interface{}{struct{}{}}
		val, err = xyml.SliceToYamlNode(slice2)

		So(err, ShouldNotBeNil)
		So(val, ShouldBeNil)

		slice3 := []int{3}
		val, err = xyml.SliceToYamlNode(slice3)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.SequenceNode)
		So(val.Tag, ShouldEqual, xyml.TagSequence)
		So(len(val.Content), ShouldEqual, 1)
		So(val.Content[0].Kind, ShouldEqual, yaml.ScalarNode)
		So(val.Content[0].Tag, ShouldEqual, xyml.TagInt)
		So(val.Content[0].Value, ShouldEqual, "3")

		val, err = xyml.SliceToYamlNode(nil)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.SequenceNode)
		So(val.Tag, ShouldEqual, xyml.TagSequence)
		So(val.Content, ShouldBeEmpty)

		slice4 := []string{}
		val, err = xyml.SliceToYamlNode(&slice4)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.SequenceNode)
		So(val.Tag, ShouldEqual, xyml.TagSequence)
		So(val.Content, ShouldBeEmpty)

		var slice5 interface{} = (*string)(nil)
		val, err = xyml.SliceToYamlNode(slice5)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.SequenceNode)
		So(val.Tag, ShouldEqual, xyml.TagSequence)
		So(val.Content, ShouldBeEmpty)

		slice6 := "hi"
		val, err = xyml.SliceToYamlNode(slice6)

		So(err, ShouldNotBeNil)
		So(val, ShouldBeNil)
	})
}

func TestToYamlNode(t *testing.T) {
	Convey("ToYamlNode", t, func() {
		val, err := xyml.ToYamlNode(nil)

		So(err, ShouldBeNil)
		So(val.Kind, ShouldEqual, yaml.ScalarNode)
		So(val.Tag, ShouldEqual, xyml.TagNil)
		So(val.Content, ShouldBeEmpty)
	})
}
