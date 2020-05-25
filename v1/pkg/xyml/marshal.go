package xyml

import "gopkg.in/yaml.v3"

// Marshaler provides a way for xyml's YAML utilities to convert custom types
// into YAML nodes suitable for use with go-yaml v3.
type Marshaler interface {
	// ToYAML provides a way to convert non-builtin types to YAML nodes using the
	// utilities in the xyml package.
	ToYAML() (*yaml.Node, error)
}
