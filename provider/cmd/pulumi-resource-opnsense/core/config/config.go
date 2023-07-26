package config

import (
	"fmt"
	"github.com/oss4u/go-opnsense/opnsense"
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type Config struct {
	Address           string `pulumi:"address" provider:"secret"`
	Key               string `pulumi:"key" provider:"secret"`
	Secret            string `pulumi:"secret" provider:"secret"`
	Version           string `pulumi:"version"`
	PluginDownloadURL string `pulumi:"pluginDownloadURL"`
	Api               *opnsense.OpnSenseApi
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Address, "The address of the fw. (without /api)")
	a.Describe(&c.Key, "The key to access the api of the fw.")
	a.Describe(&c.Secret, `The secret to access the api of the fw.`)
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx provider.Context) error {
	c.Api = opnsense.GetOpnSenseClient(c.Address, c.Key, c.Secret)
	return nil
}

func (c *Config) diffConfig(ctx provider.Context, req provider.DiffRequest) (provider.DiffResponse, error) {
	return provider.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          false,
		DetailedDiff:        nil,
	}, nil
}

func (c *Config) Diff(ctx provider.Context, id string, olds any, news Config) (provider.DiffResponse, error) {
	ctx.Log(diag.Info, "Running Diff - Config")
	ctx.Log(diag.Info, fmt.Sprintf("ConfigOLD: %+v", olds))
	ctx.Log(diag.Info, fmt.Sprintf("ConfigNEW: %+v", news))
	return provider.DiffResponse{
		HasChanges:          false,
		DeleteBeforeReplace: false,
	}, nil
}

func (c Config) Update(ctx provider.Context, id string, olds any, news any, preview bool) (any, error) {
	return news, nil
}
