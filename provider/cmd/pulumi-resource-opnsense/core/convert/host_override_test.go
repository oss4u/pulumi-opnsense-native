package convert

import (
	goopnsense "github.com/oss4u/go-opnsense/opnsense/core/unbound"
	plugin "github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/unbound"
	"reflect"
	"testing"
)

func TestHostOverrideArgsToOverridesHost(t *testing.T) {
	type args struct {
		args *plugin.HostOverrideArgs
	}
	tests := []struct {
		name string
		args args
		want *goopnsense.OverridesHost
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HostOverrideArgsToOverridesHost(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HostOverrideArgsToOverridesHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverridesHostToHostOverrideArgs(t *testing.T) {
	type args struct {
		host *goopnsense.OverridesHost
	}
	tests := []struct {
		name string
		args args
		want *plugin.HostOverrideArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OverridesHostToHostOverrideArgs(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OverridesHostToHostOverrideArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
