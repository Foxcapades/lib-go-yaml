package xyml

// YAML type tags.
const (
	TagBinary    = "!!binary"
	TagBool      = "!!bool"
	TagFloat     = "!!float"
	TagInt       = "!!int"
	TagMerge     = "!!merge"
	TagNil       = "!!null"
	TagString    = "!!str"
	TagTimestamp = "!!timestamp"
	TagValue     = "!!value"
	TagYaml      = "!!yaml"

	TagMap        = "!!map"
	TagOrderedMap = "!!omap"
	TagPairs      = "!!pairs"
	TagSet        = "!!set"
	TagSequence   = "!!seq"
)
