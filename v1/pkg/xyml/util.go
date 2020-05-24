package xyml

import "gopkg.in/yaml.v3"

// MapForEach calls the given function for each key/value pair in the given YAML
// map node.
//
// Returns an error if the given YAML node is not a map type, or if the passed
// function itself returns an error.
func MapForEach(y *yaml.Node, fn func(k, v *yaml.Node) error) error {
	if err := RequireMap(y); err != nil {
		return err
	}

	for i := 0; i < len(y.Content); i += 2 {
		if err := fn(y.Content[i], y.Content[i+1]); err != nil {
			return err
		}
	}

	return nil
}

// SequenceForEach calls the given function for each value in the given YAML
// sequence node.
//
// Returns an error if the given YAML node is not a sequence type, or if the
// passed function itself returns an error.
func SequenceForEach(y *yaml.Node, fn func(v *yaml.Node) error) error {
	if err := RequireSequence(y); err != nil {
		return err
	}

	for _, n := range y.Content {
		if err := fn(n); err != nil {
			return err
		}
	}

	return nil
}
