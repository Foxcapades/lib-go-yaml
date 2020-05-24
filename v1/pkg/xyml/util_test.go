package xyml_test

import (
	"errors"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMapForEach(t *testing.T) {
	Convey("MapForEach", t, func() {
		Convey("given a map node", func() {
			Convey("and a function that does not return an error", func() {
				test := xyml.NewMapNode(5)
				So(xyml.MapAppend(test, "bugs", "fish"), ShouldBeNil)
				So(xyml.MapAppend(test, "fish", "cats"), ShouldBeNil)
				So(xyml.MapAppend(test, "cats", "dogs"), ShouldBeNil)
				So(xyml.MapAppend(test, "dogs", "alligators"), ShouldBeNil)
				So(xyml.MapAppend(test, "alligators", "Steve Irwin"), ShouldBeNil)

				hits := 0

				So(xyml.MapForEach(test, func(k, v *yaml.Node) error {
					switch hits {
					case 0:
						So(k.Value, ShouldEqual, "bugs")
						So(v.Value, ShouldEqual, "fish")
					case 1:
						So(k.Value, ShouldEqual, "fish")
						So(v.Value, ShouldEqual, "cats")
					case 2:
						So(k.Value, ShouldEqual, "cats")
						So(v.Value, ShouldEqual, "dogs")
					case 3:
						So(k.Value, ShouldEqual, "dogs")
						So(v.Value, ShouldEqual, "alligators")
					case 4:
						So(k.Value, ShouldEqual, "alligators")
						So(v.Value, ShouldEqual, "Steve Irwin")
					case 5:
						So(hits, ShouldEqual, 5)
					}

					hits++

					return nil
				}), ShouldBeNil)

				So(hits, ShouldEqual, 5)
			})

			Convey("given a function that returns an error", func() {
				test := xyml.NewMapNode(1)
				So(xyml.MapAppend(test, "hi", "bye"), ShouldBeNil)
				So(xyml.MapForEach(test, func(k, v *yaml.Node) error {
					return errors.New("test error")
				}).Error(), ShouldEqual, "test error")
			})
		})

		Convey("given a non-map node", func() {
			test := xyml.NewNullNode()
			So(xyml.MapForEach(test, func(_, _ *yaml.Node) error { return nil }),
				ShouldNotBeNil)
		})
	})
}


func TestSequenceForEach(t *testing.T) {
	Convey("SequenceForEach", t, func() {
		Convey("given a sequence node", func() {
			Convey("and a function that does not return an error", func() {
				test := xyml.NewSequenceNode(5)
				So(xyml.SequenceAppend(test, "apple"), ShouldBeNil)
				So(xyml.SequenceAppend(test, "worm"), ShouldBeNil)
				So(xyml.SequenceAppend(test, "bird"), ShouldBeNil)
				So(xyml.SequenceAppend(test, "cat"), ShouldBeNil)
				So(xyml.SequenceAppend(test, "cat lady"), ShouldBeNil)

				hits := 0

				So(xyml.SequenceForEach(test, func(v *yaml.Node) error {
					switch hits {
					case 0:
						So(v.Value, ShouldEqual, "apple")
					case 1:
						So(v.Value, ShouldEqual, "worm")
					case 2:
						So(v.Value, ShouldEqual, "bird")
					case 3:
						So(v.Value, ShouldEqual, "cat")
					case 4:
						So(v.Value, ShouldEqual, "cat lady")
					case 5:
						So(hits, ShouldEqual, 5)
					}

					hits++

					return nil
				}), ShouldBeNil)

				So(hits, ShouldEqual, 5)
			})

			Convey("given a function that returns an error", func() {
				test := xyml.NewSequenceNode(1)
				So(xyml.SequenceAppend(test, "hi"), ShouldBeNil)
				So(xyml.SequenceForEach(test, func(v *yaml.Node) error {
					return errors.New("test error")
				}).Error(), ShouldEqual, "test error")
			})
		})

		Convey("given a non-sequence node", func() {
			test := xyml.NewNullNode()
			So(xyml.SequenceForEach(test, func(*yaml.Node) error {return nil}),
				ShouldNotBeNil)
		})
	})
}
