//
// Copyright (c) 2019 Stackinsights to present
// All rights reserved
//

package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"go/types"
	"os"
	"strings"
)

func addingOmitemptyToTag(tag string) string {
	jsonPrefix := strings.Index(tag, "json:\"")
	if jsonPrefix < 0 {
		return tag + "json:\",omitempty\""
	}
	endInx := strings.Index(tag[jsonPrefix+6:], "\"")
	jsonTag := "json:\"" + tag[jsonPrefix+6:endInx+6] + ",omitempty\""
	return jsonTag + tag[endInx+7:]
}

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			switch field.Type.(type) {
			case *types.Pointer:
				field.Tag = addingOmitemptyToTag(field.Tag)
			}
		}
	}

	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg, api.ReplacePlugin(&p))

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
