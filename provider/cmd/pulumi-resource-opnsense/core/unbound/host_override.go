package unbound

import (
	"fmt"
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
	Enabled     *bool   `pulumi:"enabled"`
	Hostname    *string `pulumi:"hostname"`
	Domain      *string `pulumi:"domain"`
	Rr          *string `pulumi:"rr"`
	MxPrio      *int    `pulumi:"mx_prio,optional"`
	Mx          *string `pulumi:"mx,optional"`
	Server      *string `pulumi:"server,optional"`
	Description *string `pulumi:"description"`
	//Aliases     *[]HostAliasArgs `pulumi:"aliases,optional"`
}

var _ = (infer.CustomRead[HostOverrideArgs, HostOverrideState])((*HostOverride)(nil))
var _ = (infer.CustomUpdate[HostOverrideArgs, HostOverrideState])((*HostOverride)(nil))
var _ = (infer.CustomDelete[HostOverrideState])((*HostOverride)(nil))

// Each resource has a state, describing the fields that exist on the created resource.
type HostOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostOverrideArgs
	// Here we define a required output called result.
	Id string `pulumi:"result"`
}

func (HostOverride) GetApi(ctx p.Context) unbound.Overrides {
	cfg := infer.GetConfig[config.Config](ctx)
	return unbound.Get_HostOverrides(cfg.Api)
}

// All resources must implement Create at a minumum.
func (h HostOverride) Create(ctx p.Context, name string, input HostOverrideArgs, preview bool) (string, HostOverrideState, error) {
	ctx.Log(diag.Info, "Running CREATE")
	state := HostOverrideState{HostOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	var err error
	state.Id, err = h.createHostOverride(ctx, &input)
	return state.Id, state, err
}

func (h HostOverride) Delete(ctx p.Context, id string, args HostOverrideState) error {
	ctx.Log(diag.Info, "Running DELETE")
	err := h.deleteHostOverride(ctx, id)
	return err
}

func (h HostOverride) Update(ctx p.Context, id string, old HostOverrideState, news HostOverrideArgs, preview bool) (HostOverrideState, error) {
	ctx.Log(diag.Info, "Running UPDATE")
	if preview {
		return HostOverrideState{
			HostOverrideArgs: news,
		}, nil
	}
	overrides := h.GetApi(ctx)
	host := HostOverrideArgsToOverridesHost(&news)
	host.Host.Uuid = id
	_, err := overrides.Update(host)
	return HostOverrideState{
		HostOverrideArgs: news,
	}, err
}

func (h HostOverride) Read(ctx p.Context, id string, inputs HostOverrideArgs, state HostOverrideState) (canonicalID string, normalizedInputs HostOverrideArgs, normalizedState HostOverrideState, err error) {
	ctx.Log(diag.Info, "Running READ")
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(id)
	newArgs := OverridesHostToHostOverrideArgs(host)
	return id, *newArgs, HostOverrideState{
		HostOverrideArgs: *newArgs,
		Result:           id,
	}, err
}

func (HostOverride) Diff(ctx p.Context, _ string, old HostOverrideState, new HostOverrideArgs) (p.DiffResponse, error) {
	ctx.Log(diag.Info, "Running DIFF")
	diffs := map[string]p.PropertyDiff{}
	if *old.Hostname != *new.Hostname {
		ctx.Log(diag.Debug, fmt.Sprintf("Hostname differs: %s/%s", *old.Hostname, *new.Hostname))
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if *old.Domain != *new.Domain {
		ctx.Log(diag.Debug, fmt.Sprintf("Domain differs: %s/%s", *old.Domain, *new.Domain))
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if *old.Description != *new.Description {
		ctx.Log(diag.Debug, fmt.Sprintf("Description differs: %s/%s", *old.Description, *new.Description))
		diffs["description"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if *old.Enabled != *new.Enabled {
		ctx.Log(diag.Debug, fmt.Sprintf("Enabled differs: %s/%s", *old.Enabled, *new.Enabled))
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if *old.Rr == "A" {
		if *old.Server != *new.Server {
			ctx.Log(diag.Debug, fmt.Sprintf("Server differs: %s/%s", *old.Server, *new.Server))
			diffs["server"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
	} else if *old.Rr == "MX" {
		if *old.Mx != *new.Mx {
			ctx.Log(diag.Debug, fmt.Sprintf("Mx differs: %s/%s", *old.Mx, *new.Mx))
			diffs["mx"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
		if *old.MxPrio != *new.MxPrio {
			ctx.Log(diag.Debug, fmt.Sprintf("MxPrio differs: %s/%s", *old.MxPrio, *new.MxPrio))
			diffs["mxprio"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}

	}
	diff := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}
	return diff, nil
}

func (h HostOverride) deleteHostOverride(ctx p.Context, id string) error {
	overrides := h.GetApi(ctx)
	err := overrides.DeleteByID(id)
	return err
}

func (h HostOverride) createHostOverride(ctx p.Context, args *HostOverrideArgs) (string, error) {
	overrides := h.GetApi(ctx)
	newHost := HostOverrideArgsToOverridesHost(args)
	createdHost, err := overrides.Create(newHost)
	if err != nil {
		return "", err
	}
	return createdHost.Host.GetUUID(), err
}
