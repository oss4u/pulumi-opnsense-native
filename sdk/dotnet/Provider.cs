// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Oss4u.Opnsense
{
    [OpnsenseResourceType("pulumi:providers:opnsense")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The address of the fw. (without /api)
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The key to access the api of the fw.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        [Output("pluginDownloadURL")]
        public Output<string> PluginDownloadURL { get; private set; } = null!;

        /// <summary>
        /// The secret to access the api of the fw.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("opnsense", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/oss4u/pulumi-opnsense-native",
                AdditionalSecretOutputs =
                {
                    "address",
                    "key",
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("address", required: true)]
        private Input<string>? _address;

        /// <summary>
        /// The address of the fw. (without /api)
        /// </summary>
        public Input<string>? Address
        {
            get => _address;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _address = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("key", required: true)]
        private Input<string>? _key;

        /// <summary>
        /// The key to access the api of the fw.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("pluginDownloadURL", required: true)]
        public Input<string> PluginDownloadURL { get; set; } = null!;

        [Input("secret", required: true)]
        private Input<string>? _secret;

        /// <summary>
        /// The secret to access the api of the fw.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
