package xyml

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
	"time"
)

// ToScalarValue attempts to parse the given YAML node into a scalar value.
//
// Handles types:
//   !!binary
//   !!bool
//   !!float
//   !!int
//   !!null
//   !!str
//   !!timestamp
func ToScalarValue(y *yaml.Node) (interface{}, error) {
	if err := RequireScalar(y); err != nil {
		return nil, err
	}

	switch y.Tag {
	case TagBinary:
		return ToBinary(y)
	case TagBool:
		return ToBoolean(y)
	case TagFloat:
		return ToFloat(y)
	case TagInt:
		return ToInt(y, 10)
	case TagTimestamp:
		return ToTime3339Nano(y)
	case TagNil:
		return nil, nil
	case TagString:
		fallthrough
	default:
		return y.Value, nil
	}
}

const (
	errParseBin    = "failed to parse binary node: %s"
	errParseBinPos = "failed to parse binary node @ %d:%d: %s"
)

// ToBinary attempts to parse the given node as a base64 encoded binary value.
func ToBinary(y *yaml.Node) ([]byte, error) {
	if err := RequireBinary(y); err != nil {
		return nil, err
	}

	val, err := base64.StdEncoding.DecodeString(y.Value)
	if err != nil {
		if hasPos(y) {
			return nil, fmt.Errorf(errParseBinPos, y.Line, y.Column, err)
		}

		return nil, fmt.Errorf(errParseBin, err)
	}

	return val, nil
}

const (
	errParseBool    = "failed to parse boolean node, unrecognized value \"%s\""
	errParseBoolPos = errParseBool + " @ %d:%d"
)

// ToBoolean attempts to parse the given node as a boolean value.
func ToBoolean(y *yaml.Node) (bool, error) {
	if err := RequireBool(y); err != nil {
		return false, err
	}

	switch strings.ToLower(y.Value) {
	case "y", "yes", "on", "true":
		return true, nil
	case "n", "no", "off", "false":
		return false, nil
	}

	if hasPos(y) {
		return false, fmt.Errorf(errParseBoolPos, y.Value, y.Line, y.Column)
	}

	return false, fmt.Errorf(errParseBool, y.Value)
}

const (
	errParseFloat    = "failed to parse float node: %s"
	errParseFloatPos = "failed to parse float node @ %d:%d: %s"
)

// ToFloat attempts to parse the given node as a float value.
func ToFloat(y *yaml.Node) (float64, error) {
	if err := RequireFloat(y); err != nil {
		return 0, err
	}

	val, err := strconv.ParseFloat(y.Value, 64)
	if err != nil {
		if hasPos(y) {
			return 0, fmt.Errorf(errParseFloatPos, y.Line, y.Column, err)
		}

		return 0, fmt.Errorf(errParseFloat, err)
	}

	return val, nil
}

const (
	errParseInt    = "failed to parse int node: %s"
	errParseIntPos = "failed to parse int node @ %d:%d: %s"
)

// ToInt attempts to parse the given node as a int value.
func ToInt(y *yaml.Node, base int) (int64, error) {
	if err := RequireInt(y); err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(y.Value, base, 64)
	if err != nil {
		if hasPos(y) {
			return 0, fmt.Errorf(errParseIntPos, y.Line, y.Column, err)
		}

		return 0, fmt.Errorf(errParseInt, err)
	}

	return val, nil
}

const (
	errParseTime    = "failed to parse timestamp node: %s"
	errParseTimePos = "failed to parse timestamp node @ %d:%d: %s"
)

// ToTime attempts to parse the given YAML node as a timestamp using the given
// format.
func ToTime(y *yaml.Node, format string) (time.Time, error) {
	if err := RequireTimestamp(y); err != nil {
		return time.Unix(0, 0), err
	}

	val, err := time.Parse(format, y.Value)
	if err != nil {
		if hasPos(y) {
			return time.Unix(0, 0), fmt.Errorf(errParseTimePos, y.Line,
				y.Column, err)
		}

		return time.Unix(0, 0), fmt.Errorf(errParseTime, err)
	}

	return val, nil
}

// ToTime3339Nano calls ToTime with the RFC3339Nano time format.
func ToTime3339Nano(y *yaml.Node) (time.Time, error) {
	return ToTime(y, time.RFC3339Nano)
}
