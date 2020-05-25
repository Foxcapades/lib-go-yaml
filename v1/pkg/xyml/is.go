package xyml

import "gopkg.in/yaml.v3"

// IsBinary returns whether the given YAML node is a binary node.
func IsBinary(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagBinary
}

// IsBool returns whether the given YAML node is a boolean node.
func IsBool(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagBool
}

// IsFloat returns whether the given YAML node is a float node.
func IsFloat(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagFloat
}

// IsInt returns whether the given YAML node is an int node.
func IsInt(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagInt
}

// IsMap returns whether the given YAML node is a map node.
func IsMap(y *yaml.Node) bool {
	return y.Kind == yaml.MappingNode
}

// IsNilType returns whether the given YAML node is a nil typed node.
func IsNilType(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagNil
}

// IsOrderedMap returns whether the given YAML node is an ordered map node.
func IsOrderedMap(y *yaml.Node) bool {
	return IsSequence(y) && y.Tag == TagOrderedMap
}

// IsPairs returns whether the given YAML node is a pairs node.
func IsPairs(y *yaml.Node) bool {
	return IsSequence(y) && y.Tag == TagPairs
}

// IsScalar returns whether the given YAML node is a scalar type.
func IsScalar(y *yaml.Node) bool {
	return y.Kind == yaml.ScalarNode
}

// IsSequence returns whether the given YAML node is a sequence node.
func IsSequence(y *yaml.Node) bool {
	return y.Kind == yaml.SequenceNode
}

// IsTimestamp returns whether the given YAML node is a timestamp node.
func IsTimestamp(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagTimestamp
}

// IsString returns whether the given YAML node is a string node.
func IsString(y *yaml.Node) bool {
	return IsScalar(y) && y.Tag == TagString
}
