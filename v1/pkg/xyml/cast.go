package xyml

import (
	"fmt"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v3"
)

// ScalarToYamlNode converts the given scalar node to a YAML node.
//
// Returns an error if the given value cannot be converted to a YAML scalar
// node.
func ScalarToYamlNode(val interface{}) (*yaml.Node, error) {
	if val == nil {
		return NewNullNode(), nil
	}

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String:
		return NewStringNode(val.(string)), nil
	case reflect.Int:
		return NewIntNode(int64(val.(int)), 10), nil
	case reflect.Int8:
		return NewIntNode(int64(val.(int8)), 10), nil
	case reflect.Int16:
		return NewIntNode(int64(val.(int16)), 10), nil
	case reflect.Int32:
		return NewIntNode(int64(val.(int32)), 10), nil
	case reflect.Int64:
		return NewIntNode(val.(int64), 10), nil
	case reflect.Uint:
		return &yaml.Node{
			Kind:  yaml.ScalarNode,
			Tag:   TagInt,
			Value: strconv.FormatUint(uint64(val.(uint)), 10),
		}, nil
	case reflect.Uint8:
		return NewIntNode(int64(val.(uint8)), 10), nil
	case reflect.Uint16:
		return NewIntNode(int64(val.(uint16)), 10), nil
	case reflect.Uint32:
		return NewIntNode(int64(val.(uint32)), 10), nil
	case reflect.Uint64:
		return &yaml.Node{
			Kind:  yaml.ScalarNode,
			Tag:   TagInt,
			Value: strconv.FormatUint(val.(uint64), 10),
		}, nil
	case reflect.Bool:
		return NewBoolNode(val.(bool)), nil
	case reflect.Float32:
		return NewFloatNode(float64(val.(float32)), 'g', 10), nil
	case reflect.Float64:
		return NewFloatNode(val.(float64), 'g', 10), nil
	}

	if canNil(v.Kind()) && v.IsNil() {
		return NewNullNode(), nil
	}

	if v.Kind() == reflect.Ptr {
		return ScalarToYamlNode(v.Elem().Interface())
	}

	return nil, fmt.Errorf("cannot convert type %s to a scalar node", v.Type())
}

// MapToYamlNode converts the given map to a YAML node.
//
// Returns an error if the given value is not a map, if the map key type is not
// a scalar value, or if the map contains a value which cannot be automatically
// converted to a YAML node (such as a struct value).
func MapToYamlNode(val interface{}) (*yaml.Node, error) {
	if val == nil {
		return NewMapNode(0), nil
	}

	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Map {
		out := NewMapNode(v.Len())
		it := v.MapRange()

		for it.Next() {
			if key, err := ScalarToYamlNode(it.Key().Interface()); err != nil {
				return nil, err
			} else if val, err := ToYamlNode(it.Value().Interface()); err != nil {
				return nil, err
			} else {
				_ = MapAppendNode(out, key, val)
			}
		}

		return out, nil
	}

	if canNil(v.Kind()) && v.IsNil() {
		return NewMapNode(0), nil
	}

	if v.Kind() == reflect.Ptr {
		return MapToYamlNode(v.Elem().Interface())
	}

	return nil, fmt.Errorf("cannot convert type %s to a map node", v.Type())
}

// SliceToYamlNode converts the given slice or array value into a YAML sequence.
//
// Returns an error if the given value is not an array or slice, or if the
// array/slice contains a value that cannot be automatically converted to a YAML
// node (such as a struct value).
func SliceToYamlNode(val interface{}) (*yaml.Node, error) {
	if val == nil {
		return NewSequenceNode(0), nil
	}

	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		ln := v.Len()
		out := NewSequenceNode(ln)

		for i := 0; i < ln; i++ {
			val, err := ToYamlNode(v.Index(i).Interface())

			if err != nil {
				return nil, err
			}

			_ = SequenceAppendNode(out, val)
		}

		return out, nil
	}

	if canNil(v.Kind()) && v.IsNil() {
		return NewSequenceNode(0), nil
	}

	if v.Kind() == reflect.Ptr {
		return SliceToYamlNode(v.Elem().Interface())
	}

	return nil, fmt.Errorf("cannot convert type %s to a sequence node", v.Type())
}

// ToYamlNode converts arbitrary typed data to a YAML node.
//
// Note: This method is a shorthand call for all 3 of ScalarToYamlNode,
// SliceToYamlNode, and MapToYamlNode.  As such it follows the same rules and
// cannot handle user defined types such as structs.
func ToYamlNode(val interface{}) (*yaml.Node, error) {

	if val == nil {
		return NewNullNode(), nil
	}

	if tmp, ok := val.(*yaml.Node); ok {
		return tmp, nil
	}

	if tmp, ok := val.(Marshaler); ok {
		return tmp.ToYAML()
	}

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Map:
		return MapToYamlNode(val)
	case reflect.Array, reflect.Slice:
		return SliceToYamlNode(val)
	}

	if canNil(v.Kind()) && v.IsNil() {
		return NewNullNode(), nil
	}

	if v.Kind() == reflect.Ptr {
		return ToYamlNode(v.Elem().Interface())
	}

	return ScalarToYamlNode(val)
}

var nillable = map[reflect.Kind]bool{
	reflect.Map:       true,
	reflect.Interface: true,
	reflect.Slice:     true,
	reflect.Ptr:       true,
	reflect.Chan:      true,
	reflect.Func:      true,
	reflect.Invalid:   true,
}

func canNil(k reflect.Kind) bool {
	return nillable[k]
}
