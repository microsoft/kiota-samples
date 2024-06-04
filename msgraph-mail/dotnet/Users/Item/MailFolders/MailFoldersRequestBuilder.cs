// <auto-generated/>
using Graphdotnetv4.Models.ODataErrors;
using Graphdotnetv4.Models;
using Graphdotnetv4.Users.Item.MailFolders.Count;
using Graphdotnetv4.Users.Item.MailFolders.Item;
using Microsoft.Kiota.Abstractions.Serialization;
using Microsoft.Kiota.Abstractions;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using System.Threading;
using System;
namespace Graphdotnetv4.Users.Item.MailFolders
{
    /// <summary>
    /// Builds and executes requests for operations under \users\{user-id}\mailFolders
    /// </summary>
    public class MailFoldersRequestBuilder : BaseRequestBuilder
    {
        /// <summary>The Count property</summary>
        public Graphdotnetv4.Users.Item.MailFolders.Count.CountRequestBuilder Count
        {
            get => new Graphdotnetv4.Users.Item.MailFolders.Count.CountRequestBuilder(PathParameters, RequestAdapter);
        }
        /// <summary>Gets an item from the Graphdotnetv4.users.item.mailFolders.item collection</summary>
        /// <param name="position">The unique identifier of mailFolder</param>
        /// <returns>A <see cref="Graphdotnetv4.Users.Item.MailFolders.Item.MailFolderItemRequestBuilder"/></returns>
        public Graphdotnetv4.Users.Item.MailFolders.Item.MailFolderItemRequestBuilder this[string position]
        {
            get
            {
                var urlTplParams = new Dictionary<string, object>(PathParameters);
                urlTplParams.Add("mailFolder%2Did", position);
                return new Graphdotnetv4.Users.Item.MailFolders.Item.MailFolderItemRequestBuilder(urlTplParams, RequestAdapter);
            }
        }
        /// <summary>
        /// Instantiates a new <see cref="Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder"/> and sets the default values.
        /// </summary>
        /// <param name="pathParameters">Path parameters for the request</param>
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        public MailFoldersRequestBuilder(Dictionary<string, object> pathParameters, IRequestAdapter requestAdapter) : base(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders{?%24count,%24expand,%24filter,%24orderby,%24select,%24skip,%24top,includeHiddenFolders}", pathParameters)
        {
        }
        /// <summary>
        /// Instantiates a new <see cref="Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder"/> and sets the default values.
        /// </summary>
        /// <param name="rawUrl">The raw URL to use for the request builder.</param>
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        public MailFoldersRequestBuilder(string rawUrl, IRequestAdapter requestAdapter) : base(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders{?%24count,%24expand,%24filter,%24orderby,%24select,%24skip,%24top,includeHiddenFolders}", rawUrl)
        {
        }
        /// <summary>
        /// The user&apos;s mail folders. Read-only. Nullable.
        /// </summary>
        /// <returns>A <see cref="Graphdotnetv4.Models.MailFolderCollectionResponse"/></returns>
        /// <param name="cancellationToken">Cancellation token to use when cancelling requests</param>
        /// <param name="requestConfiguration">Configuration for the request such as headers, query parameters, and middleware options.</param>
        /// <exception cref="Graphdotnetv4.Models.ODataErrors.ODataError">When receiving a 4XX or 5XX status code</exception>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public async Task<Graphdotnetv4.Models.MailFolderCollectionResponse?> GetAsync(Action<RequestConfiguration<Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder.MailFoldersRequestBuilderGetQueryParameters>>? requestConfiguration = default, CancellationToken cancellationToken = default)
        {
#nullable restore
#else
        public async Task<Graphdotnetv4.Models.MailFolderCollectionResponse> GetAsync(Action<RequestConfiguration<Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder.MailFoldersRequestBuilderGetQueryParameters>> requestConfiguration = default, CancellationToken cancellationToken = default)
        {
#endif
            var requestInfo = ToGetRequestInformation(requestConfiguration);
            var errorMapping = new Dictionary<string, ParsableFactory<IParsable>>
            {
                { "XXX", Graphdotnetv4.Models.ODataErrors.ODataError.CreateFromDiscriminatorValue },
            };
            return await RequestAdapter.SendAsync<Graphdotnetv4.Models.MailFolderCollectionResponse>(requestInfo, Graphdotnetv4.Models.MailFolderCollectionResponse.CreateFromDiscriminatorValue, errorMapping, cancellationToken).ConfigureAwait(false);
        }
        /// <summary>
        /// Create new navigation property to mailFolders for users
        /// </summary>
        /// <returns>A <see cref="Graphdotnetv4.Models.MailFolder"/></returns>
        /// <param name="body">The request body</param>
        /// <param name="cancellationToken">Cancellation token to use when cancelling requests</param>
        /// <param name="requestConfiguration">Configuration for the request such as headers, query parameters, and middleware options.</param>
        /// <exception cref="Graphdotnetv4.Models.ODataErrors.ODataError">When receiving a 4XX or 5XX status code</exception>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public async Task<Graphdotnetv4.Models.MailFolder?> PostAsync(Graphdotnetv4.Models.MailFolder body, Action<RequestConfiguration<DefaultQueryParameters>>? requestConfiguration = default, CancellationToken cancellationToken = default)
        {
#nullable restore
#else
        public async Task<Graphdotnetv4.Models.MailFolder> PostAsync(Graphdotnetv4.Models.MailFolder body, Action<RequestConfiguration<DefaultQueryParameters>> requestConfiguration = default, CancellationToken cancellationToken = default)
        {
#endif
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = ToPostRequestInformation(body, requestConfiguration);
            var errorMapping = new Dictionary<string, ParsableFactory<IParsable>>
            {
                { "XXX", Graphdotnetv4.Models.ODataErrors.ODataError.CreateFromDiscriminatorValue },
            };
            return await RequestAdapter.SendAsync<Graphdotnetv4.Models.MailFolder>(requestInfo, Graphdotnetv4.Models.MailFolder.CreateFromDiscriminatorValue, errorMapping, cancellationToken).ConfigureAwait(false);
        }
        /// <summary>
        /// The user&apos;s mail folders. Read-only. Nullable.
        /// </summary>
        /// <returns>A <see cref="RequestInformation"/></returns>
        /// <param name="requestConfiguration">Configuration for the request such as headers, query parameters, and middleware options.</param>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public RequestInformation ToGetRequestInformation(Action<RequestConfiguration<Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder.MailFoldersRequestBuilderGetQueryParameters>>? requestConfiguration = default)
        {
#nullable restore
#else
        public RequestInformation ToGetRequestInformation(Action<RequestConfiguration<Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder.MailFoldersRequestBuilderGetQueryParameters>> requestConfiguration = default)
        {
#endif
            var requestInfo = new RequestInformation(Method.GET, UrlTemplate, PathParameters);
            requestInfo.Configure(requestConfiguration);
            requestInfo.Headers.TryAdd("Accept", "application/json");
            return requestInfo;
        }
        /// <summary>
        /// Create new navigation property to mailFolders for users
        /// </summary>
        /// <returns>A <see cref="RequestInformation"/></returns>
        /// <param name="body">The request body</param>
        /// <param name="requestConfiguration">Configuration for the request such as headers, query parameters, and middleware options.</param>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public RequestInformation ToPostRequestInformation(Graphdotnetv4.Models.MailFolder body, Action<RequestConfiguration<DefaultQueryParameters>>? requestConfiguration = default)
        {
#nullable restore
#else
        public RequestInformation ToPostRequestInformation(Graphdotnetv4.Models.MailFolder body, Action<RequestConfiguration<DefaultQueryParameters>> requestConfiguration = default)
        {
#endif
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = new RequestInformation(Method.POST, UrlTemplate, PathParameters);
            requestInfo.Configure(requestConfiguration);
            requestInfo.Headers.TryAdd("Accept", "application/json");
            requestInfo.SetContentFromParsable(RequestAdapter, "application/json", body);
            return requestInfo;
        }
        /// <summary>
        /// Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
        /// </summary>
        /// <returns>A <see cref="Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder"/></returns>
        /// <param name="rawUrl">The raw URL to use for the request builder.</param>
        public Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder WithUrl(string rawUrl)
        {
            return new Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder(rawUrl, RequestAdapter);
        }
        /// <summary>
        /// The user&apos;s mail folders. Read-only. Nullable.
        /// </summary>
        public class MailFoldersRequestBuilderGetQueryParameters 
        {
            /// <summary>Include count of items</summary>
            [QueryParameter("%24count")]
            public bool? Count { get; set; }
            /// <summary>Expand related entities</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
            [QueryParameter("%24expand")]
            public string[]? Expand { get; set; }
#nullable restore
#else
            [QueryParameter("%24expand")]
            public string[] Expand { get; set; }
#endif
            /// <summary>Filter items by property values</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
            [QueryParameter("%24filter")]
            public string? Filter { get; set; }
#nullable restore
#else
            [QueryParameter("%24filter")]
            public string Filter { get; set; }
#endif
            /// <summary>Include Hidden Folders</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
            [QueryParameter("includeHiddenFolders")]
            public string? IncludeHiddenFolders { get; set; }
#nullable restore
#else
            [QueryParameter("includeHiddenFolders")]
            public string IncludeHiddenFolders { get; set; }
#endif
            /// <summary>Order items by property values</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
            [QueryParameter("%24orderby")]
            public string[]? Orderby { get; set; }
#nullable restore
#else
            [QueryParameter("%24orderby")]
            public string[] Orderby { get; set; }
#endif
            /// <summary>Select properties to be returned</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
            [QueryParameter("%24select")]
            public string[]? Select { get; set; }
#nullable restore
#else
            [QueryParameter("%24select")]
            public string[] Select { get; set; }
#endif
            /// <summary>Skip the first n items</summary>
            [QueryParameter("%24skip")]
            public int? Skip { get; set; }
            /// <summary>Show only the first n items</summary>
            [QueryParameter("%24top")]
            public int? Top { get; set; }
        }
        /// <summary>
        /// Configuration for the request such as headers, query parameters, and middleware options.
        /// </summary>
        [Obsolete("This class is deprecated. Please use the generic RequestConfiguration class generated by the generator.")]
        public class MailFoldersRequestBuilderGetRequestConfiguration : RequestConfiguration<Graphdotnetv4.Users.Item.MailFolders.MailFoldersRequestBuilder.MailFoldersRequestBuilderGetQueryParameters>
        {
        }
        /// <summary>
        /// Configuration for the request such as headers, query parameters, and middleware options.
        /// </summary>
        [Obsolete("This class is deprecated. Please use the generic RequestConfiguration class generated by the generator.")]
        public class MailFoldersRequestBuilderPostRequestConfiguration : RequestConfiguration<DefaultQueryParameters>
        {
        }
    }
}
