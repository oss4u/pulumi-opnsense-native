// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/config"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/unbound"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"os"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	err := p.RunProvider("opnsense", Version, provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {

	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "Opnsense",
			License:     "Apache-2.0",
			Repository:  "https://github.com/oss4u/pulumi-opnsense-native",
			Publisher:   "Oss4u",
			LanguageMap: map[string]any{
				"nodejs": map[string]any{
					"packageName": "@oss4u/opnsense",
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/oss4u/pulumi-opnsense-native/sdk/go/opnsense",
				},
				"csharp": map[string]any{
					"rootNamespace": "Oss4u",
				},
			},
			PluginDownloadURL: "github://api.github.com/oss4u/pulumi-opnsense-native",
		},
		Resources: []infer.InferredResource{
			infer.Resource[unbound.HostAlias, unbound.HostAliasArgs, unbound.HostAliasState](),
			infer.Resource[unbound.HostOverride, unbound.HostOverrideArgs, unbound.HostOverrideState](),
		},
		Config: infer.Config[*config.Config](),
	})
}
