package kiotaposts.client.posts;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.BaseRequestConfiguration;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import kiotaposts.client.models.Post;
import kiotaposts.client.posts.item.PostItemRequestBuilder;
/**
 * Builds and executes requests for operations under /posts
 */
public class PostsRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the kiotaposts.client.posts.item collection
     * @param postId key: id of post
     * @return a PostItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public PostItemRequestBuilder byPostId(@javax.annotation.Nonnull final Integer postId) {
        Objects.requireNonNull(postId);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("post%2Did", postId);
        return new PostItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new PostsRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public PostsRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/posts{?userId*,title*}", pathParameters);
    }
    /**
     * Instantiates a new PostsRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public PostsRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/posts{?userId*,title*}", rawUrl);
    }
    /**
     * Get posts
     * @return a CompletableFuture of post
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<java.util.List<Post>> get() {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(null);
            return this.requestAdapter.sendCollectionAsync(requestInfo, Post::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<java.util.List<Post>> executionException = new java.util.concurrent.CompletableFuture<java.util.List<Post>>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Get posts
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of post
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<java.util.List<Post>> get(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
            return this.requestAdapter.sendCollectionAsync(requestInfo, Post::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<java.util.List<Post>> executionException = new java.util.concurrent.CompletableFuture<java.util.List<Post>>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Create post
     * @param body The request body
     * @return a CompletableFuture of post
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Post> post(@javax.annotation.Nonnull final Post body) {
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, null);
            return this.requestAdapter.sendAsync(requestInfo, Post::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Post> executionException = new java.util.concurrent.CompletableFuture<Post>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Create post
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of post
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Post> post(@javax.annotation.Nonnull final Post body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, requestConfiguration);
            return this.requestAdapter.sendAsync(requestInfo, Post::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Post> executionException = new java.util.concurrent.CompletableFuture<Post>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Get posts
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation() throws URISyntaxException {
        return toGetRequestInformation(null);
    }
    /**
     * Get posts
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        if (requestConfiguration != null) {
            final GetRequestConfiguration requestConfig = new GetRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.addQueryParameters(requestConfig.queryParameters);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Create post
     * @param body The request body
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final Post body) throws URISyntaxException {
        return toPostRequestInformation(body, null);
    }
    /**
     * Create post
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final Post body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (requestConfiguration != null) {
            final PostRequestConfiguration requestConfig = new PostRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Get posts
     */
    public class GetQueryParameters {
        /**
         * Filter results by title
         */
        @javax.annotation.Nullable
        public String title;
        /**
         * Filter results by user ID
         */
        @javax.annotation.Nullable
        public Integer userId;
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class GetRequestConfiguration extends BaseRequestConfiguration {
        /**
         * Request query parameters
         */
        @javax.annotation.Nullable
        public GetQueryParameters queryParameters = new GetQueryParameters();
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class PostRequestConfiguration extends BaseRequestConfiguration {
    }
}
