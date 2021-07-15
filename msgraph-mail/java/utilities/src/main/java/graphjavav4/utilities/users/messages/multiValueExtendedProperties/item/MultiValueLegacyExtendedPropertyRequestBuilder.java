package graphjavav4.utilities.users.messages.multiValueExtendedProperties.item;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.MiddlewareOption;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id}  */
public class MultiValueLegacyExtendedPropertyRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new MultiValueLegacyExtendedPropertyRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public MultiValueLegacyExtendedPropertyRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo() throws URISyntaxException {
        return createDeleteRequestInfo(null, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createDeleteRequestInfo(h, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
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
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo() throws URISyntaxException {
        return createGetRequestInfo(null, null, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) throws URISyntaxException {
        return createGetRequestInfo(q, null, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createGetRequestInfo(q, h, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
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
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body) throws URISyntaxException {
        Objects.requireNonNull(body);
        return createPatchRequestInfo(body, null, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        return createPatchRequestInfo(body, h, null);
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.setContentFromParsable(body, httpCore, "application/json");
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addMiddlewareOptions(o.toArray(new MiddlewareOption[0]));
        }
        return requestInfo;
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
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
     * The collection of multi-value extended properties defined for the message. Nullable.
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
     * The collection of multi-value extended properties defined for the message. Nullable.
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
     * The collection of multi-value extended properties defined for the message. Nullable.
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
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @return a CompletableFuture of MultiValueLegacyExtendedProperty
     */
    public java.util.concurrent.CompletableFuture<MultiValueLegacyExtendedProperty> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(null, null, null);
            return this.httpCore.sendAsync(requestInfo, MultiValueLegacyExtendedProperty.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param q Request query parameters
     * @return a CompletableFuture of MultiValueLegacyExtendedProperty
     */
    public java.util.concurrent.CompletableFuture<MultiValueLegacyExtendedProperty> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, null, null);
            return this.httpCore.sendAsync(requestInfo, MultiValueLegacyExtendedProperty.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param q Request query parameters
     * @return a CompletableFuture of MultiValueLegacyExtendedProperty
     */
    public java.util.concurrent.CompletableFuture<MultiValueLegacyExtendedProperty> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, null);
            return this.httpCore.sendAsync(requestInfo, MultiValueLegacyExtendedProperty.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @return a CompletableFuture of MultiValueLegacyExtendedProperty
     */
    public java.util.concurrent.CompletableFuture<MultiValueLegacyExtendedProperty> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, MultiValueLegacyExtendedProperty.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MultiValueLegacyExtendedProperty
     */
    public java.util.concurrent.CompletableFuture<MultiValueLegacyExtendedProperty> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(q, h, o);
            return this.httpCore.sendAsync(requestInfo, MultiValueLegacyExtendedProperty.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, null, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, null);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MultiValueLegacyExtendedProperty body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<MiddlewareOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(body, h, o);
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** The collection of multi-value extended properties defined for the message. Nullable.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Expand related entities  */
        @javax.annotation.Nullable
        public String[] expand;
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
    }
}
