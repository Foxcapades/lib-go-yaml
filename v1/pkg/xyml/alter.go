package xyml

import "gopkg.in/yaml.v3"

// MapAppendNode appends the given key/value nodes (k, v) to the given map node
// (y).
func MapAppendNode(y, k, v *yaml.Node) error {
	if err := RequireMap(y); err != nil {
		return err
	}
	y.Content = append(y.Content, k, v)
	return nil
}

// MapAppend converts the given key/value args (k, v) to YAML nodes then appends
// them to the given map node (y).
func MapAppend(y *yaml.Node, k, v interface{}) error {
	if err := RequireMap(y); err != nil {
		return err
	}
	if key, err := ScalarToYamlNode(k); err != nil {
		return err
	} else if val, err := ToYamlNode(v); err != nil {
		return err
	} else {
		y.Content = append(y.Content, key, val)
	}
	return nil
}

// SequenceAppendNode appends the given YAML node value to the given sequence
// node.
func SequenceAppendNode(y, v *yaml.Node) error {
	if err := RequireSequence(y); err != nil {
		return err
	}
	y.Content = append(y.Content, v)
	return nil
}

func SequenceAppend(y *yaml.Node, v interface{}) error {
	if err := RequireSequence(y); err != nil {
		return err
	}
	if val, err := ToYamlNode(v); err != nil {
		return err
	} else {
		y.Content = append(y.Content, val)
	}
	return nil
}