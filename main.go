package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	enc := yaml.NewEncoder(os.Stdout)

	nodeA := &yaml.Node{
		Kind: yaml.MappingNode,
		Tag: "!!map",
		Style: yaml.TaggedStyle,
	}

	enc.Encode(nodeA)

	nodeB := &yaml.Node{
		Kind: yaml.MappingNode,
		Tag: "tag:yaml.org,2002:map",
		Style: yaml.TaggedStyle,
	}

	enc.Encode(nodeB)
}
