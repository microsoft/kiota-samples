package graphjavav4.utilities.users.mailFolders.messageRules;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.MiddlewareOption;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.MessageRule;
import java.io.InputStream;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messageRules  */
public class MessageRulesRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    /** Whether the current path is a raw URL  */
    private final boolean isRawUrl;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new MessageRulesRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public MessageRulesRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore) {
        this(currentPath, httpCore, true);
    }
    /**
     * Instantiates a new MessageRulesRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @param isRawUrl Whether the current path is a raw URL
     * @return a void
     */
    public MessageRulesRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore, final boolean isRawUrl) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "/messageRules";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo() throws URISyntaxException {
        return createGetRequestInfo(null, null, null);
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) throws URISyntaxException {
        return createGetRequestInfo(q, null, null);
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param h Request headers
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createGetRequestInfo(q, h, null);
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            this.setUri(currentPath, pathSegment, isRawUrl);
            httpMethod = HttpMethod.GET;
        }};
        if (q != null) {
            final GetQueryParameters qParams = new GetQueryParameters();
            q.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
        }
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addMiddlewareOptions(o.toArray(new MiddlewareOption[0]));
        }
        return requestInfo;
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPostRequestInfo(@javax.annotation.Nonnull final MessageRule body) throws URISyntaxException {
        return createPostRequestInfo(body, null, null);
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPostRequestInfo(@javax.annotation.Nonnull final MessageRule body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createPostRequestInfo(body, h, null);
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPostRequestInfo(@javax.annotation.Nonnull final MessageRule body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            this.setUri(currentPath, pathSegment, isRawUrl);
            httpMethod = HttpMethod.POST;
        }};
        requestInfo.setContentFromParsable(httpCore, "application/json", body);
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addMiddlewareOptions(o.toArray(new MiddlewareOption[0]));
        }
        return requestInfo;
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @return a CompletableFuture of MessageRulesResponse
     */
    public java.util.concurrent.CompletableFuture<MessageRulesResponse> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(null, null, null);
            return this.httpCore.sendAsync(requestInfo, MessageRulesResponse.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param q Request query parameters
     * @return a CompletableFuture of MessageRulesResponse
     */
    public java.util.concurrent.CompletableFuture<MessageRulesResponse> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, null, null);
            return this.httpCore.sendAsync(requestInfo, MessageRulesResponse.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param h Request headers
     * @param q Request query parameters
     * @return a CompletableFuture of MessageRulesResponse
     */
    public java.util.concurrent.CompletableFuture<MessageRulesResponse> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, null);
            return this.httpCore.sendAsync(requestInfo, MessageRulesResponse.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @return a CompletableFuture of MessageRulesResponse
     */
    public java.util.concurrent.CompletableFuture<MessageRulesResponse> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, MessageRulesResponse.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MessageRulesResponse
     */
    public java.util.concurrent.CompletableFuture<MessageRulesResponse> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, MessageRulesResponse.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @return a CompletableFuture of MessageRule
     */
    public java.util.concurrent.CompletableFuture<MessageRule> post(@javax.annotation.Nonnull final MessageRule body) {
        try {
            final RequestInfo requestInfo = createPostRequestInfo(body, null, null);
            return this.httpCore.sendAsync(requestInfo, MessageRule.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @param h Request headers
     * @return a CompletableFuture of MessageRule
     */
    public java.util.concurrent.CompletableFuture<MessageRule> post(@javax.annotation.Nonnull final MessageRule body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createPostRequestInfo(body, h, null);
            return this.httpCore.sendAsync(requestInfo, MessageRule.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a CompletableFuture of MessageRule
     */
    public java.util.concurrent.CompletableFuture<MessageRule> post(@javax.annotation.Nonnull final MessageRule body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createPostRequestInfo(body, h, o);
            return this.httpCore.sendAsync(requestInfo, MessageRule.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MessageRule
     */
    public java.util.concurrent.CompletableFuture<MessageRule> post(@javax.annotation.Nonnull final MessageRule body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPostRequestInfo(body, h, o);
            return this.httpCore.sendAsync(requestInfo, MessageRule.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** The collection of rules that apply to the user's Inbox folder.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Include count of items  */
        @javax.annotation.Nullable
        public Boolean count;
        /** Expand related entities  */
        @javax.annotation.Nullable
        public String[] expand;
        /** Filter items by property values  */
        @javax.annotation.Nullable
        public String filter;
        /** Order items by property values  */
        @javax.annotation.Nullable
        public String[] orderby;
        /** Search items by search phrases  */
        @javax.annotation.Nullable
        public String search;
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
        /** Skip the first n items  */
        @javax.annotation.Nullable
        public Integer skip;
        /** Show only the first n items  */
        @javax.annotation.Nullable
        public Integer top;
    }
}
