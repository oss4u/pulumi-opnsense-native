package unbound

import (
	p "github.com/pulumi/pulumi-go-provider"
	"math/rand"
	"time"
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
type HostAlias struct{}

// Each resource has in input struct, defining what arguments it accepts.
type HostAliasArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type HostAliasState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostAliasArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minumum.
func (HostAlias) Create(ctx p.Context, name string, input HostAliasArgs, preview bool) (string, HostAliasState, error) {
	state := HostAliasState{HostAliasArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = createHostAlias(input.Length)
	return name, state, nil
}

func createHostAlias(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]rune, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
