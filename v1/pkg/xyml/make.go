package xyml

import (
	"encoding/base64"
	"gopkg.in/yaml.v3"
	"strconv"
	"time"
)

func NewBinaryNode(v []byte) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagBinary,
		Value: base64.StdEncoding.EncodeToString(v),
	}
}

func NewBoolNode(v bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagBool,
		Value: strconv.FormatBool(v),
	}
}

func NewFloatNode(val float64, rep byte, prec int) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagFloat,
		Value: strconv.FormatFloat(val, rep, prec, 64),
	}
}

func NewIntNode(val int64, base int) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagInt,
		Value: strconv.FormatInt(val, base),
	}
}

func NewNullNode() *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: TagNil}
}

func NewStringNode(v string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: TagString, Value: v}
}

func NewTimestampNode(v time.Time) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagTimestamp,
		Value: v.Format(time.RFC3339Nano),
	}
}

func NewMapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     TagMap,
		Content: make([]*yaml.Node, 0, size*2),
	}
}

func NewOrderedMapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     TagOrderedMap,
		Content: make([]*yaml.Node, 0, size*2),
	}
}

func NewPairsNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagPairs,
		Content: make([]*yaml.Node, 0, size),
	}
}

func NewSetNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagSet,
		Content: make([]*yaml.Node, 0, size),
	}
}

func NewSequenceNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagSequence,
		Content: make([]*yaml.Node, 0, size),
	}
}
