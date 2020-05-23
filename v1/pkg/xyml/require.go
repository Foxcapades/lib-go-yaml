package xyml

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

const (
	errNoPos   = `expected a node of type %s, instead got type %s`
	errWithPos = errNoPos + ` @ %d:%d`
	errKind    = `expected a node of kind %s, instead got %s`
	errKindPos = errKind + ` @ %d:%d`
)

// RequireBinary returns an error if the given node is not of type binary.
func RequireBinary(y *yaml.Node) error {
	if !IsBinary(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagBinary, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagBinary, y.Tag)
	}
	return nil
}

// RequireBool returns an error if the given node is not of type bool.
func RequireBool(y *yaml.Node) error {
	if !IsBool(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagBool, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagBool, y.Tag)
	}
	return nil
}

// RequireFloat returns an error if the given node is not of type float.
func RequireFloat(y *yaml.Node) error {
	if !IsFloat(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagFloat, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagFloat, y.Tag)
	}
	return nil
}

// RequireBool returns an error if the given node is not of type int.
func RequireInt(y *yaml.Node) error {
	if !IsInt(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagInt, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagInt, y.Tag)
	}
	return nil
}

// RequireMap returns an error if the given node is not a mapping node.
func RequireMap(y *yaml.Node) error {
	if !IsMap(y) {
		if hasPos(y) {
			return fmt.Errorf(errKindPos, kindToString(yaml.MappingNode),
				kindToString(y.Kind), y.Line, y.Column)
		}
		return fmt.Errorf(errKind, kindToString(yaml.MappingNode),
			kindToString(y.Kind))
	}
	return nil
}

// RequireNilType returns an error if the given node is not of type null.
func RequireNilType(y *yaml.Node) error {
	if !IsNilType(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagNil, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagNil, y.Tag)
	}
	return nil
}

// RequireOrderedMap returns an error if the given node is not of type omap.
func RequireOrderedMap(y *yaml.Node) error {
	if !IsOrderedMap(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagOrderedMap, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagOrderedMap, y.Tag)
	}
	return nil
}

// RequirePairs returns an error if the given node is not of type pairs.
func RequirePairs(y *yaml.Node) error {
	if !IsPairs(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagPairs, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagPairs, y.Tag)
	}
	return nil
}

// RequireScalar returns an error if the given node is not a scalar node.
func RequireScalar(y *yaml.Node) error {
	if !IsScalar(y) {
		if hasPos(y) {
			return fmt.Errorf(errKindPos, kindToString(yaml.ScalarNode),
				kindToString(y.Kind), y.Line, y.Column)
		}
		return fmt.Errorf(errKind, kindToString(yaml.ScalarNode),
			kindToString(y.Kind))
	}
	return nil
}

// RequireSequence returns an error if the given node is not a sequence node.
func RequireSequence(y *yaml.Node) error {
	if !IsSequence(y) {
		if hasPos(y) {
			return fmt.Errorf(errKindPos, kindToString(yaml.SequenceNode),
				kindToString(y.Kind), y.Line, y.Column)
		}
		return fmt.Errorf(errKind, kindToString(yaml.SequenceNode),
			kindToString(y.Kind))
	}
	return nil
}

// RequireTimestamp returns an error if the given node is not of type timestamp.
func RequireTimestamp(y *yaml.Node) error {
	if !IsTimestamp(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagTimestamp, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagTimestamp, y.Tag)
	}
	return nil
}


// RequireString returns an error if the given node is not of type string.
func RequireString(y *yaml.Node) error {
	if !IsString(y) {
		if hasPos(y) {
			return fmt.Errorf(errWithPos, TagString, y.Tag, y.Line, y.Column)
		}
		return fmt.Errorf(errNoPos, TagString, y.Tag)
	}
	return nil
}

func hasPos(y *yaml.Node) bool {
	return y.Line > 0
}

func kindToString(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.SequenceNode:
		return "sequence"
	case yaml.MappingNode:
		return "mapping"
	case yaml.ScalarNode:
		return "scalar"
	case yaml.AliasNode:
		return "alias"
	}
	return "unknown"
}