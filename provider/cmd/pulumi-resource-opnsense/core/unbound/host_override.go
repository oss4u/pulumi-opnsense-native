package unbound

import (
	"fmt"
	"github.com/oss4u/go-opnsense/opnsense"
	"github.com/oss4u/go-opnsense/opnsense/core/unbound"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
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
	fmt.Printf("Running CREATE")
	cfg := infer.GetConfig[config.Config](ctx)
	state := HostOverrideState{HostOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	var err error
	state.Result, err = createHostOverride(&input, cfg.Api)
	return state.Result, state, err
}

func (HostOverride) Delete(ctx p.Context, id string, input HostOverrideArgs) error {
	fmt.Printf("Running DELETE")
	cfg := infer.GetConfig[config.Config](ctx)
	err := deleteHostOverride(id, cfg.Api)
	return err
}

func (HostOverride) Update(ctx p.Context, id string, olds HostOverrideArgs, news HostOverrideArgs, preview bool) (HostOverrideArgs, error) {
	ctx.Log(diag.Info, "Running UPDATE")
	fmt.Printf("Running UPDATE")
	if preview {
		return news, nil
	}
	cfg := infer.GetConfig[config.Config](ctx)
	overrides := unbound.Get_HostOverrides(cfg.Api)
	host := HostOverrideArgsToOverridesHost(&news)
	host.Host.Uuid = id
	_, err := overrides.Update(host)
	return news, err
}

func (HostOverride) Diff(ctx p.Context, id string, old HostOverrideArgs, new HostOverrideArgs) (p.DiffResponse, error) {
	ctx.Log(diag.Info, "Running DIFF")
	fmt.Printf("Running DIFF")
	diffs := map[string]p.PropertyDiff{}
	if old.Hostname != new.Hostname {
		ctx.Log(diag.Info, "Hostname differs")
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Domain != new.Domain {
		ctx.Log(diag.Info, "Domain differs")
		diffs["domain"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Description != new.Description {
		ctx.Log(diag.Info, "Description differs")
		diffs["description"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Enabled != new.Enabled {
		ctx.Log(diag.Info, "Enabled differs")
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.UpdateReplace,
		}
	}
	if old.Rr == "A" {
		if old.Server != new.Server {
			ctx.Log(diag.Info, "Server differs")
			diffs["server"] = p.PropertyDiff{
				Kind: p.UpdateReplace,
			}
		}
	} else if old.Rr == "MX" {
		if old.Mx != new.Mx {
			ctx.Log(diag.Info, "Mx differs")
			diffs["mx"] = p.PropertyDiff{
				Kind: p.UpdateReplace,
			}
		}
		if old.MxPrio != new.MxPrio {
			ctx.Log(diag.Info, "MxPrio differs")
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

func deleteHostOverride(id string, api *opnsense.OpnSenseApi) error {
	overrides := unbound.Get_HostOverrides(api)
	err := overrides.DeleteByID(id)
	return err
}

func createHostOverride(args *HostOverrideArgs, api *opnsense.OpnSenseApi) (string, error) {
	overrides := unbound.Get_HostOverrides(api)
	newHost := HostOverrideArgsToOverridesHost(args)
	createdHost, err := overrides.Create(newHost)
	if err != nil {
		return "", err
	}
	return createdHost.Host.GetUUID(), err
}
