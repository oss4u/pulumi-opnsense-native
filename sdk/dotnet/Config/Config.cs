// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Oss4u.Opnsense
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("opnsense");

        private static readonly __Value<string?> _address = new __Value<string?>(() => __config.Get("address"));
        /// <summary>
        /// The address of the fw. (without /api)
        /// </summary>
        public static string? Address
        {
            get => _address.Get();
            set => _address.Set(value);
        }

        private static readonly __Value<string?> _key = new __Value<string?>(() => __config.Get("key"));
        /// <summary>
        /// The key to access the api of the fw.
        /// </summary>
        public static string? Key
        {
            get => _key.Get();
            set => _key.Set(value);
        }

        private static readonly __Value<string?> _pluginDownloadURL = new __Value<string?>(() => __config.Get("pluginDownloadURL"));
        public static string? PluginDownloadURL
        {
            get => _pluginDownloadURL.Get();
            set => _pluginDownloadURL.Set(value);
        }

        private static readonly __Value<string?> _secret = new __Value<string?>(() => __config.Get("secret"));
        /// <summary>
        /// The secret to access the api of the fw.
        /// </summary>
        public static string? Secret
        {
            get => _secret.Get();
            set => _secret.Set(value);
        }

        private static readonly __Value<string?> _version = new __Value<string?>(() => __config.Get("version"));
        public static string? Version
        {
            get => _version.Get();
            set => _version.Set(value);
        }

    }
}
