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
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/unbound"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	p.RunProvider("opnsense", Version,
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[unbound.HostAlias, unbound.HostAliasArgs, unbound.HostAliasState](),
				infer.Resource[unbound.HostOverride, unbound.HostOverrideArgs, unbound.HostOverrideState](),
			},
		}))
}
