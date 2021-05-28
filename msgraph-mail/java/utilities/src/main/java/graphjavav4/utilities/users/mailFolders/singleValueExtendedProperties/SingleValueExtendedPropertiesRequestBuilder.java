package graphjavav4.utilities.users.mailFolders.singleValueExtendedProperties;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactory;
import graphjavav4.utilities.users.SingleValueLegacyExtendedProperty;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties  */
public class SingleValueExtendedPropertiesRequestBuilder {
    /** Current path for the request  */
    @javax.annotation.Nullable
    public String currentPath;
    /** Core service to use to execute the requests  */
    @javax.annotation.Nullable
    public HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    @javax.annotation.Nonnull
    private final String pathSegment = "/singleValueExtendedProperties";
    /** Factory to use to get a serializer for payload serialization  */
    @javax.annotation.Nullable
    public SerializationWriterFactory serializerFactory;
    /**
     * Get singleValueExtendedProperties from users
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo() throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.GET;
        }};
        return requestInfo;
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.GET;
        }};
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param h Request headers
     * @param q Request query parameters
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
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
        return requestInfo;
    }
    /**
     * Create new navigation property to singleValueExtendedProperties for users
     * @param <U> The generic type for the model, must inherit SingleValueLegacyExtendedProperty
     * @param body The request body
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public <U extends SingleValueLegacyExtendedProperty> RequestInfo createPostRequestInfo(@javax.annotation.Nonnull final U body) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.POST;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "application/json");
        return requestInfo;
    }
    /**
     * Create new navigation property to singleValueExtendedProperties for users
     * @param <U> The generic type for the model, must inherit SingleValueLegacyExtendedProperty
     * @param body The request body
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public <U extends SingleValueLegacyExtendedProperty> RequestInfo createPostRequestInfo(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.POST;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "application/json");
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param <U> The generic type for the model, must inherit SingleValueExtendedPropertiesResponse
     * @return a CompletableFuture of SingleValueExtendedPropertiesResponse
     */
    public <U extends SingleValueExtendedPropertiesResponse> java.util.concurrent.CompletableFuture<U> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueExtendedPropertiesResponse.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param <U> The generic type for the model, must inherit SingleValueExtendedPropertiesResponse
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of SingleValueExtendedPropertiesResponse
     */
    public <U extends SingleValueExtendedPropertiesResponse> java.util.concurrent.CompletableFuture<U> get(@javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueExtendedPropertiesResponse.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param <U> The generic type for the model, must inherit SingleValueExtendedPropertiesResponse
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of SingleValueExtendedPropertiesResponse
     */
    public <U extends SingleValueExtendedPropertiesResponse> java.util.concurrent.CompletableFuture<U> get(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                h
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueExtendedPropertiesResponse.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get singleValueExtendedProperties from users
     * @param <U> The generic type for the model, must inherit SingleValueExtendedPropertiesResponse
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of SingleValueExtendedPropertiesResponse
     */
    public <U extends SingleValueExtendedPropertiesResponse> java.util.concurrent.CompletableFuture<U> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                q, h
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueExtendedPropertiesResponse.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Create new navigation property to singleValueExtendedProperties for users
     * @param <U> The generic type for the model, must inherit SingleValueLegacyExtendedProperty
     * @param body The request body
     * @return a CompletableFuture of SingleValueLegacyExtendedProperty
     */
    public <U extends SingleValueLegacyExtendedProperty> java.util.concurrent.CompletableFuture<U> post(@javax.annotation.Nonnull final U body) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPostRequestInfo(
                body
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueLegacyExtendedProperty.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Create new navigation property to singleValueExtendedProperties for users
     * @param <U> The generic type for the model, must inherit SingleValueLegacyExtendedProperty
     * @param body The request body
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of SingleValueLegacyExtendedProperty
     */
    public <U extends SingleValueLegacyExtendedProperty> java.util.concurrent.CompletableFuture<U> post(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPostRequestInfo(
                body
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueLegacyExtendedProperty.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Create new navigation property to singleValueExtendedProperties for users
     * @param <U> The generic type for the model, must inherit SingleValueLegacyExtendedProperty
     * @param body The request body
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of SingleValueLegacyExtendedProperty
     */
    public <U extends SingleValueLegacyExtendedProperty> java.util.concurrent.CompletableFuture<U> post(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPostRequestInfo(
                body, h
            );
            return this.httpCore.sendAsync(requestInfo, SingleValueLegacyExtendedProperty.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** Get singleValueExtendedProperties from users  */
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
