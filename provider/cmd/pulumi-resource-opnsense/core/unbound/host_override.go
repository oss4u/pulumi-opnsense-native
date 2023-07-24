package unbound

import (
	"github.com/oss4u/go-opnsense/opnsense"
	"github.com/oss4u/go-opnsense/opnsense/core/unbound"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
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
	MxPrio      int    `pulumi:"mx_prio,optional"`
	Mx          string `pulumi:"mx,optional"`
	Server      string `pulumi:"server,optional"`
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
	config := infer.GetConfig[config.Config](ctx)
	state := HostOverrideState{HostOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = createHostOverride(input, config.Api)
	return name, state, nil
}

func (HostOverride) Delete(ctx p.Context, id string, input HostOverrideArgs) error {
	config := infer.GetConfig[config.Config](ctx)
	err := deleteHostOverride(ctx, id, config.Api)
	return err
}

func (HostOverride) Diff(ctx p.Context, id string, old HostOverrideArgs, new HostOverrideArgs) (p.DiffResponse, error) {

	diffs := map[string]p.PropertyDiff{}
	if old.Hostname != new.Hostname {
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Domain != new.Domain {
		diffs["domain"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Description != new.Description {
		diffs["description"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Enabled != new.Enabled {
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Rr == "A" {
		if old.Server != new.Server {
			diffs["server"] = p.PropertyDiff{
				Kind: p.UpdateReplace,
			}
		}
	} else if old.Rr == "MX" {
		if old.Mx != new.Mx {
			diffs["mx"] = p.PropertyDiff{
				Kind: p.UpdateReplace,
			}
		}
		if old.MxPrio != new.MxPrio {
			diffs["mxprio"] = p.PropertyDiff{
				Kind: p.UpdateReplace,
			}
		}

	}
	diff := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}
	diff.DeleteBeforeReplace = true
	return diff, nil
}

func deleteHostOverride(ctx p.Context, id string, api *opnsense.OpnSenseApi) error {
	overrides := unbound.Get_HostOverrides(api)
	err := overrides.DeleteByID(id)
	return err
}

func createHostOverride(host HostOverrideArgs, api *opnsense.OpnSenseApi) string {
	overrides := unbound.Get_HostOverrides(api)
	newHost := unbound.OverridesHost{Host: unbound.OverridesHostDetails{
		Uuid:        "",
		Enabled:     host.Enabled,
		Hostname:    host.Hostname,
		Domain:      host.Domain,
		Rr:          host.Rr,
		Description: host.Description,
	}}
	if host.Rr == "A" {
		newHost.Host.Server = host.Server
	} else if host.Rr == "AAAA" {
		newHost.Host.Server = host.Server
	} else if host.Rr == "MX" {
		newHost.Host.Mx = host.Mx
		newHost.Host.Mxprio = host.MxPrio
	}
	created_host, _ := overrides.Create(&newHost)
	return created_host.Host.GetUUID()
}
