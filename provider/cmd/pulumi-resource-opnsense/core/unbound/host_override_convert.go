package unbound

import (
	goopnsense "github.com/oss4u/go-opnsense/opnsense/core/unbound"
)

func OverridesHostToHostOverrideArgs(host *goopnsense.OverridesHost) *HostOverrideArgs {
	args := HostOverrideArgs{
		Enabled:     &host.Host.Enabled,
		Hostname:    &host.Host.Hostname,
		Domain:      &host.Host.Domain,
		Rr:          &host.Host.Rr,
		Description: &host.Host.Description,
	}
	if host.Host.Rr == "A" || host.Host.Rr == "AAAA" {
		args.Server = &host.Host.Server
	} else if host.Host.Rr == "MX" {
		args.Mx = &host.Host.Mx
		args.MxPrio = &host.Host.Mxprio
	}
	return &args
}

func HostOverrideArgsToOverridesHost(args *HostOverrideArgs) *goopnsense.OverridesHost {
	host := goopnsense.OverridesHostDetails{
		Enabled:     *args.Enabled,
		Hostname:    *args.Hostname,
		Domain:      *args.Domain,
		Rr:          *args.Rr,
		Description: *args.Description,
	}
	if *args.Rr == "A" || *args.Rr == "AAAA" {
		host.Server = *args.Server
	} else if *args.Rr == "MX" {
		host.Mx = *args.Mx
		host.Mxprio = *args.MxPrio
	}
	return &goopnsense.OverridesHost{Host: host}
}
