package xyml

import (
	"encoding/base64"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// NewBinaryNode returns a new binary typed YAML node with the given content.
func NewBinaryNode(v []byte) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagBinary,
		Value: base64.StdEncoding.EncodeToString(v),
	}
}

// NewBoolNode returns a new boolean typed YAML node with the given value.
func NewBoolNode(v bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagBool,
		Value: strconv.FormatBool(v),
	}
}

// NewFloatNode returns a new float typed YAML node with the given value.
func NewFloatNode(val float64, rep byte, prec int) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagFloat,
		Value: strconv.FormatFloat(val, rep, prec, 64),
	}
}

// NewIntNode returns a new int typed YAML node with the given value.
func NewIntNode(val int64, base int) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagInt,
		Value: strconv.FormatInt(val, base),
	}
}

// NewFloatNode returns a new null typed YAML node.
func NewNullNode() *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: TagNil}
}

// NewStringNode returns a new string typed YAML node with the given value.
func NewStringNode(v string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: TagString, Value: v}
}

// NewFloatNode returns a new float typed YAML node with the given value.
func NewTimestampNode(v time.Time) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   TagTimestamp,
		Value: v.Format(time.RFC3339Nano),
	}
}

// NewMapNode returns a new mapping typed YAML node presized to the given size.
func NewMapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     TagMap,
		Content: make([]*yaml.Node, 0, size*2),
	}
}

// NewMapNode returns a new ordered map typed YAML node presized to the given
// size.
func NewOrderedMapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     TagOrderedMap,
		Content: make([]*yaml.Node, 0, size*2),
	}
}

// NewPairsNode returns a new pairs typed YAML node presized to the given size.
func NewPairsNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagPairs,
		Content: make([]*yaml.Node, 0, size),
	}
}

// NewSetNode returns a new set typed YAML node presized to the given size.
func NewSetNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagSet,
		Content: make([]*yaml.Node, 0, size),
	}
}

// NewMapNode returns a new sequence typed YAML node presized to the given size.
func NewSequenceNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     TagSequence,
		Content: make([]*yaml.Node, 0, size),
	}
}
