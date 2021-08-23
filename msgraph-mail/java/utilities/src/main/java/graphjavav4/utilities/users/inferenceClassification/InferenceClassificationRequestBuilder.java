package graphjavav4.utilities.users.inferenceClassification;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.MiddlewareOption;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.inferenceClassification.overrides.item.InferenceClassificationOverrideRequestBuilder;
import graphjavav4.utilities.users.inferenceClassification.overrides.OverridesRequestBuilder;
import java.io.InputStream;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/inferenceClassification  */
public class InferenceClassificationRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    /** Whether the current path is a raw URL  */
    private final boolean isRawUrl;
    @javax.annotation.Nonnull
    public OverridesRequestBuilder overrides() {
        return new OverridesRequestBuilder(currentPath + pathSegment, httpCore, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public InferenceClassificationRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore) {
        this(currentPath, httpCore, true);
    }
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @param isRawUrl Whether the current path is a raw URL
     * @return a void
     */
    public InferenceClassificationRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore, final boolean isRawUrl) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "/inferenceClassification";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo() throws URISyntaxException {
        return createDeleteRequestInfo(null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createDeleteRequestInfo(h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            this.setUri(currentPath, pathSegment, isRawUrl);
            httpMethod = HttpMethod.DELETE;
        }};
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addMiddlewareOptions(o.toArray(new MiddlewareOption[0]));
        }
        return requestInfo;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo() throws URISyntaxException {
        return createGetRequestInfo(null, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) throws URISyntaxException {
        return createGetRequestInfo(q, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createGetRequestInfo(q, h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
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
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final InferenceClassification body) throws URISyntaxException {
        return createPatchRequestInfo(body, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createPatchRequestInfo(body, h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            this.setUri(currentPath, pathSegment, isRawUrl);
            httpMethod = HttpMethod.PATCH;
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
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete() {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(null, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(h, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a CompletableFuture of InferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(null, null, null);
            return this.httpCore.sendAsync(requestInfo, InferenceClassification.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param q Request query parameters
     * @return a CompletableFuture of InferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, null, null);
            return this.httpCore.sendAsync(requestInfo, InferenceClassification.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param q Request query parameters
     * @return a CompletableFuture of InferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, null);
            return this.httpCore.sendAsync(requestInfo, InferenceClassification.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @return a CompletableFuture of InferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, InferenceClassification.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of InferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, InferenceClassification.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.inferenceClassification.overrides collection
     * @param id Unique identifier of the item
     * @return a InferenceClassificationOverrideRequestBuilder
     */
    @javax.annotation.Nonnull
    public InferenceClassificationOverrideRequestBuilder overrides(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new InferenceClassificationOverrideRequestBuilder(currentPath + pathSegment + "/overrides/" + id, httpCore, false);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body) {
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, null, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Expand related entities  */
        @javax.annotation.Nullable
        public String[] expand;
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
    }
}
