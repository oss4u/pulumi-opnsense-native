package config

import (
	"fmt"
	"github.com/oss4u/go-opnsense/opnsense"
	provider2 "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type Config struct {
	Address           string `pulumi:"fw_api_address" provider:"secret"`
	Key               string `pulumi:"fw_api_key" provider:"secret"`
	Secret            string `pulumi:"fw_api_secret" provider:"secret"`
	Version           string `pulumi:"version"`
	PluginDownloadURL string `pulumi:"pluginDownloadURL"`
	Api               *opnsense.OpnSenseApi
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Address, "The username. Its important but not secret.")
	a.Describe(&c.Key, "The password. It is very secret.")
	a.Describe(&c.Secret, `The (entirely uncryptographic) hash function used to encode the "password".`)
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx provider2.Context) error {
	msg := fmt.Sprintf("opnsense provider setup with address: %q, a key (its %q) and a secret (its %q)", c.Address, c.Key, c.Secret)
	c.Api = opnsense.GetOpnSenseClient(c.Address, c.Key, c.Secret)
	ctx.Log(diag.Info, msg)
	return nil
}
