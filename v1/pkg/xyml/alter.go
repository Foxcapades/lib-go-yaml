package xyml

import "gopkg.in/yaml.v3"

func MapAppendNode(y, k, v *yaml.Node) {
	y.Content = append(y.Content, k, v)
}

func MapAppend(y *yaml.Node, k, v interface{}) error {
	if key, err := ScalarToYamlNode(k); err != nil {
		return err
	} else if val, err := ToYamlNode(v); err != nil {
		return err
	} else {
		y.Content = append(y.Content, key, val)
	}
	return nil
}

func SequenceAppendNode(y, v *yaml.Node) {
	y.Content = append(y.Content, v)
}

func SequenceAppend(y *yaml.Node, v interface{}) error {
	if val, err := ToYamlNode(v); err != nil {
		return err
	} else {
		y.Content = append(y.Content, val)
	}
	return nil
}