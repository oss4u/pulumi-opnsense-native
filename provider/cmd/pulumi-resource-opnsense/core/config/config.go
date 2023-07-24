package config

import (
	"fmt"
	"github.com/oss4u/go-opnsense/opnsense"
	provider2 "github.com/pulumi/pulumi-go-provider"
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

func (c *Config) Configure(ctx provider2.Context) error {
	msg := fmt.Sprintf("opnsense provider setup with address: %q, a key (its %q) and a secret (its %q)", c.Address, c.Key, c.Secret)
	c.Api = opnsense.GetOpnSenseClient(c.Address, c.Key, c.Secret)
	ctx.Log(diag.Info, msg)
	return nil
}
