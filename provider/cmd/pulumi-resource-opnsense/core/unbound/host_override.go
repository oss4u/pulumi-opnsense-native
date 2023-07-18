package unbound

import (
	p "github.com/pulumi/pulumi-go-provider"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type HostOverride struct{}

type HostOverrideArgs struct {
	Enabled     bool   `pulumi:"enabled"`
	Hostname    string `pulumi:"hostname"`
	Domain      string `pulumi:"domain"`
	Rr          string `pulumi:"rr"`
	MxPrio      int    `pulumi:"mx_prio"`
	Mx          string `pulumi:"mx"`
	Server      string `pulumi:"server"`
	Description string `pulumi:"description"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type HostOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostOverrideArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minumum.
func (HostOverride) Create(ctx p.Context, name string, input HostOverrideArgs, preview bool) (string, HostOverrideState, error) {
	state := HostOverrideState{HostOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = createHostOverride(input)
	return name, state, nil
}

func createHostOverride(host HostOverrideArgs) string {
	if host.Rr == "A" {

	} else if host.Rr == "AAAA" {

	} else if host.Rr == "MX" {

	}
	return string("result")
}
