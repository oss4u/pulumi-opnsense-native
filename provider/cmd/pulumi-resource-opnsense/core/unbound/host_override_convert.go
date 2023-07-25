package unbound

import (
	goopnsense "github.com/oss4u/go-opnsense/opnsense/core/unbound"
)

func OverridesHostToHostOverrideArgs(host *goopnsense.OverridesHost) *HostOverrideArgs {
	args := HostOverrideArgs{
		Enabled:     host.Host.Enabled,
		Hostname:    host.Host.Hostname,
		Domain:      host.Host.Domain,
		Rr:          host.Host.Rr,
		MxPrio:      host.Host.Mxprio,
		Mx:          host.Host.Mx,
		Server:      host.Host.Server,
		Description: host.Host.Description,
	}
	return &args
}

func HostOverrideArgsToOverridesHost(args *HostOverrideArgs) *goopnsense.OverridesHost {
	host := goopnsense.OverridesHostDetails{
		Enabled:     args.Enabled,
		Hostname:    args.Hostname,
		Domain:      args.Domain,
		Rr:          args.Rr,
		Description: args.Description,
		Mxprio:      args.MxPrio,
		Mx:          args.Mx,
		Server:      args.Server,
	}
	return &goopnsense.OverridesHost{Host: host}
}
