// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class CacheKeyPolicyResponse
    {
        /// <summary>
        /// If true, requests to different hosts will be cached separately.
        /// </summary>
        public readonly bool IncludeHost;
        /// <summary>
        /// If true, http and https requests will be cached separately.
        /// </summary>
        public readonly bool IncludeProtocol;
        /// <summary>
        /// If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist. If neither is set, the entire query string will be included. If false, the query string will be excluded from the cache key entirely.
        /// </summary>
        public readonly bool IncludeQueryString;
        /// <summary>
        /// Names of query string parameters to exclude in cache keys. All other parameters will be included. Either specify query_string_whitelist or query_string_blacklist, not both. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public readonly ImmutableArray<string> QueryStringBlacklist;
        /// <summary>
        /// Names of query string parameters to include in cache keys. All other parameters will be excluded. Either specify query_string_whitelist or query_string_blacklist, not both. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public readonly ImmutableArray<string> QueryStringWhitelist;

        [OutputConstructor]
        private CacheKeyPolicyResponse(
            bool includeHost,

            bool includeProtocol,

            bool includeQueryString,

            ImmutableArray<string> queryStringBlacklist,

            ImmutableArray<string> queryStringWhitelist)
        {
            IncludeHost = includeHost;
            IncludeProtocol = includeProtocol;
            IncludeQueryString = includeQueryString;
            QueryStringBlacklist = queryStringBlacklist;
            QueryStringWhitelist = queryStringWhitelist;
        }
    }
}