package unbound

import (
	"fmt"
	goopnsense "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/go-opnsense/opnsense/types"
	"strconv"
)

func OverridesHostToHostOverrideArgs(host *goopnsense.OverridesHost) *HostOverrideArgs {
	args := HostOverrideArgs{
		Enabled:     setEnabledOnArgs(&host.Host.Enabled),
		Hostname:    &host.Host.Hostname,
		Domain:      &host.Host.Domain,
		Rr:          setRROnArgs(&host.Host.Rr),
		Description: &host.Host.Description,
	}
	if host.Host.Rr == "A" || host.Host.Rr == "AAAA" {
		args.Server = &host.Host.Server
	} else if host.Host.Rr == "MX" {
		args.Mx = &host.Host.Mx
		args.MxPrio = setMxPrioOnArgs(&host.Host.Mxprio)
	}
	return &args
}

func setRROnArgs(r *goopnsense.Rr) *string {
	var result string
	result = fmt.Sprintf("%s", *r)
	return &result
}

func setMxPrioOnArgs(m *goopnsense.MxPrio) *int {
	var result int
	result, _ = strconv.Atoi(fmt.Sprintf("%d", *m))
	return &result
}

func setEnabledOnArgs(t *types.Bool) *bool {
	var b bool
	if *t == true {
		b = true
	} else {
		b = false
	}
	return &b

}

func HostOverrideArgsToOverridesHost(args *HostOverrideArgs) *goopnsense.OverridesHost {
	host := goopnsense.OverridesHostDetails{
		Enabled:     setEnabledOnHost(args.Enabled),
		Hostname:    *args.Hostname,
		Domain:      *args.Domain,
		Rr:          setRrOnHost(args.Rr),
		Description: *args.Description,
	}
	if *args.Rr == "A" || *args.Rr == "AAAA" {
		host.Server = *args.Server
	} else if *args.Rr == "MX" {
		host.Mx = *args.Mx
		host.Mxprio = setMxPrioOnHost(args.MxPrio)
	}
	return &goopnsense.OverridesHost{Host: host}
}

func setEnabledOnHost(i *bool) types.Bool {
	if *i {
		return true
	}
	return false
}

func setRrOnHost(i *string) goopnsense.Rr {
	return goopnsense.Rr(*i)
}

func setMxPrioOnHost(i *int) goopnsense.MxPrio {
	return goopnsense.MxPrio(*i)
}
