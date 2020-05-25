package xyml_test

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
	"time"
)

func TestToBinary(t *testing.T) {
	Convey("ToBinary", t, func() {
		Convey("With a non-binary node", func() {
			val, err := xyml.ToBinary(xyml.NewNullNode())
			So(err, ShouldNotBeNil)
			So(val, ShouldBeNil)
		})

		Convey("with a bad value", func() {
			Convey("and a position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagBinary,
					Value:  "hello",
					Line:   666,
					Column: 69,
				}

				val, err := xyml.ToBinary(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeNil)
				So(strings.Index(err.Error(), "666:69"), ShouldBeGreaterThan, -1)
			})
			Convey("and no position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagBinary,
					Value:  "hello",
				}

				val, err := xyml.ToBinary(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeNil)
			})
		})

		Convey("with a valid value", func() {
			val, err := xyml.ToBinary(xyml.NewBinaryNode([]byte("hello")))

			So(err, ShouldBeNil)
			So(val, ShouldResemble, []byte("hello"))
		})
	})
}

func TestToBoolean(t *testing.T) {
	Convey("ToBoolean", t, func() {
		Convey("With a non-binary node", func() {
			val, err := xyml.ToBoolean(xyml.NewNullNode())
			So(err, ShouldNotBeNil)
			So(val, ShouldBeFalse)
		})

		Convey("with a bad value", func() {
			Convey("and a position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagBool,
					Value:  "hello",
					Line:   666,
					Column: 69,
				}

				val, err := xyml.ToBoolean(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeFalse)
				So(strings.Index(err.Error(), "666:69"), ShouldBeGreaterThan, -1)
			})
			Convey("and no position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagBool,
					Value:  "hello",
				}

				val, err := xyml.ToBoolean(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeFalse)
			})
		})

		Convey("with a valid value", func() {
			Convey("that equates to true", func() {
				validTrue := []string{
					"y", "Y", "yes", "Yes", "YES",
					"on", "On", "ON",
					"true", "True", "TRUE",
				}

				for _, str := range validTrue {
					raw := &yaml.Node{
						Kind:        yaml.ScalarNode,
						Tag:         xyml.TagBool,
						Value:       str,
					}

					val, err := xyml.ToBoolean(raw)

					So(err, ShouldBeNil)
					So(val, ShouldBeTrue)
				}
			})

			Convey("that equates to false", func() {
				validFalse := []string{
					"n", "N", "no", "No", "NO",
					"off", "Off", "OFF",
					"false", "False", "FALSE",
				}

				for _, str := range validFalse {
					raw := &yaml.Node{
						Kind:        yaml.ScalarNode,
						Tag:         xyml.TagBool,
						Value:       str,
					}

					val, err := xyml.ToBoolean(raw)

					So(err, ShouldBeNil)
					So(val, ShouldBeFalse)
				}
			})
		})
	})
}

func TestToFloat(t *testing.T) {
	Convey("ToFloat", t, func() {
		Convey("With a non-binary node", func() {
			val, err := xyml.ToFloat(xyml.NewNullNode())
			So(err, ShouldNotBeNil)
			So(val, ShouldBeZeroValue)
		})

		Convey("with a bad value", func() {
			Convey("and a position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagFloat,
					Value:  "hello",
					Line:   666,
					Column: 69,
				}

				val, err := xyml.ToFloat(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeZeroValue)
				So(strings.Index(err.Error(), "666:69"), ShouldBeGreaterThan, -1)
			})
			Convey("and no position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagFloat,
					Value:  "hello",
				}

				val, err := xyml.ToFloat(raw)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeZeroValue)
			})
		})

		Convey("with a valid value", func() {
			val, err := xyml.ToFloat(xyml.NewFloatNode(3, 'f', 2))

			So(err, ShouldBeNil)
			So(val, ShouldEqual, 3)
		})
	})
}

func TestToInt(t *testing.T) {
	Convey("ToInt", t, func() {
		Convey("With a non-binary node", func() {
			val, err := xyml.ToInt(xyml.NewNullNode(), 10)
			So(err, ShouldNotBeNil)
			So(val, ShouldBeZeroValue)
		})

		Convey("with a bad value", func() {
			Convey("and a position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagInt,
					Value:  "hello",
					Line:   666,
					Column: 69,
				}

				val, err := xyml.ToInt(raw, 10)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeZeroValue)
				So(strings.Index(err.Error(), "666:69"), ShouldBeGreaterThan, -1)
			})
			Convey("and no position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagInt,
					Value:  "hello",
				}

				val, err := xyml.ToInt(raw, 10)

				So(err, ShouldNotBeNil)
				So(val, ShouldBeZeroValue)
			})
		})

		Convey("with a valid value", func() {
			val, err := xyml.ToInt(xyml.NewIntNode(3, 10), 10)

			So(err, ShouldBeNil)
			So(val, ShouldEqual, 3)
		})
	})
}

func TestToTime(t *testing.T) {
	Convey("ToTime", t, func() {
		Convey("With a non-binary node", func() {
			_, err := xyml.ToTime(xyml.NewNullNode(), "")
			So(err, ShouldNotBeNil)
		})

		Convey("with a bad value", func() {
			Convey("and a position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagTimestamp,
					Value:  "hello",
					Line:   666,
					Column: 69,
				}

				_, err := xyml.ToTime(raw, time.RFC3339)

				So(err, ShouldNotBeNil)
				So(strings.Index(err.Error(), "666:69"), ShouldBeGreaterThan, -1)
			})
			Convey("and no position", func() {
				raw := &yaml.Node{
					Kind:   yaml.ScalarNode,
					Tag:    xyml.TagTimestamp,
					Value:  "hello",
				}

				_, err := xyml.ToTime(raw, time.RFC3339)

				So(err, ShouldNotBeNil)
			})
		})

		Convey("with a valid value", func() {
			val, err := xyml.ToTime(xyml.NewTimestampNode(time.Date(1988, 9, 18, 1,
				2, 3, 40000, time.UTC)), time.RFC3339Nano)

			So(err, ShouldBeNil)
			So(val.Year(), ShouldEqual, 1988)
			So(val.Month(), ShouldEqual, 9)
			So(val.Day(), ShouldEqual, 18)
			So(val.Hour(), ShouldEqual, 1)
			So(val.Minute(), ShouldEqual, 2)
			So(val.Second(), ShouldEqual, 3)
			So(val.Nanosecond(), ShouldEqual, 40000)
		})
	})
}

func TestToTime3339Nano(t *testing.T) {
	Convey("ToTime3339Nano", t, func() {
		Convey("with a valid value", func() {
			val, err := xyml.ToTime3339Nano(xyml.NewTimestampNode(time.Date(1988, 9,
				18, 1, 2, 3, 40000, time.UTC)))

			So(err, ShouldBeNil)
			So(val.Year(), ShouldEqual, 1988)
			So(val.Month(), ShouldEqual, 9)
			So(val.Day(), ShouldEqual, 18)
			So(val.Hour(), ShouldEqual, 1)
			So(val.Minute(), ShouldEqual, 2)
			So(val.Second(), ShouldEqual, 3)
			So(val.Nanosecond(), ShouldEqual, 40000)
		})
	})
}

func TestToScalarValue(t *testing.T) {
	Convey("ToScalarValue", t, func() {
		Convey("given a non-scalar yaml node", func() {
			bad := xyml.NewMapNode(0)
			_, err := xyml.ToScalarValue(bad)
			So(err, ShouldNotBeNil)
		})

		Convey("given a scalar node", func() {
			raw := xyml.NewNullNode()
			val, err := xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldBeNil)

			raw = xyml.NewBinaryNode([]byte("hello"))
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldResemble, []byte("hello"))

			raw = xyml.NewBoolNode(true)
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldBeTrue)

			raw = xyml.NewFloatNode(3, 'f', 3)
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldResemble, float64(3))

			raw = xyml.NewIntNode(10, 10)
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldResemble, int64(10))

			raw = xyml.NewStringNode("hello")
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)
			So(val, ShouldResemble, "hello")

			raw = xyml.NewTimestampNode(time.Date(1988, 9, 18, 1, 2, 3, 40000,
				time.UTC))
			val, err = xyml.ToScalarValue(raw)
			So(err, ShouldBeNil)

			v2 := val.(time.Time)
			So(v2.Year(), ShouldEqual, 1988)
			So(v2.Month(), ShouldEqual, 9)
			So(v2.Day(), ShouldEqual, 18)
			So(v2.Hour(), ShouldEqual, 1)
			So(v2.Minute(), ShouldEqual, 2)
			So(v2.Second(), ShouldEqual, 3)
			So(v2.Nanosecond(), ShouldEqual, 40000)
		})
	})
}