// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class HostOverride extends pulumi.CustomResource {
    /**
     * Get an existing HostOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HostOverride {
        return new HostOverride(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'opnsense:unbound:HostOverride';

    /**
     * Returns true if the given object is an instance of HostOverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostOverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostOverride.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly domain!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean>;
    public readonly hostname!: pulumi.Output<string>;
    public readonly mx!: pulumi.Output<string>;
    public readonly mx_prio!: pulumi.Output<number>;
    public /*out*/ readonly result!: pulumi.Output<string>;
    public readonly rr!: pulumi.Output<string>;
    public readonly server!: pulumi.Output<string>;

    /**
     * Create a HostOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostOverrideArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.mx === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mx'");
            }
            if ((!args || args.mx_prio === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mx_prio'");
            }
            if ((!args || args.rr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rr'");
            }
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["mx"] = args ? args.mx : undefined;
            resourceInputs["mx_prio"] = args ? args.mx_prio : undefined;
            resourceInputs["rr"] = args ? args.rr : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["result"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["mx"] = undefined /*out*/;
            resourceInputs["mx_prio"] = undefined /*out*/;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["rr"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostOverride.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a HostOverride resource.
 */
export interface HostOverrideArgs {
    description: pulumi.Input<string>;
    domain: pulumi.Input<string>;
    enabled: pulumi.Input<boolean>;
    hostname: pulumi.Input<string>;
    mx: pulumi.Input<string>;
    mx_prio: pulumi.Input<number>;
    rr: pulumi.Input<string>;
    server: pulumi.Input<string>;
}
