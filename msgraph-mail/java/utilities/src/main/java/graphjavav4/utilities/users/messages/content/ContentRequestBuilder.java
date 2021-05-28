package graphjavav4.utilities.users.messages.Content;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactory;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/$value  */
public class ContentRequestBuilder {
    /** Current path for the request  */
    @javax.annotation.Nullable
    public String currentPath;
    /** Core service to use to execute the requests  */
    @javax.annotation.Nullable
    public HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    @javax.annotation.Nonnull
    private final String pathSegment = "/$value";
    /** Factory to use to get a serializer for payload serialization  */
    @javax.annotation.Nullable
    public SerializationWriterFactory serializerFactory;
    /**
     * Get media content for the navigation property messages from users
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
     * Get media content for the navigation property messages from users
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
     * Update media content for the navigation property messages in users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param body Binary request body
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public <U extends InputStream> RequestInfo createPutRequestInfo(@javax.annotation.Nonnull final U body) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PUT;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "");
        return requestInfo;
    }
    /**
     * Update media content for the navigation property messages in users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param body Binary request body
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public <U extends InputStream> RequestInfo createPutRequestInfo(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PUT;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "");
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * Get media content for the navigation property messages from users
     * @param <U> The generic type for the model, must inherit InputStream
     * @return a CompletableFuture of InputStream
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<U> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, InputStream.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get media content for the navigation property messages from users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of InputStream
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<U> get(@javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, InputStream.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get media content for the navigation property messages from users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of InputStream
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<U> get(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                h
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, InputStream.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Update media content for the navigation property messages in users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param body Binary request body
     * @return a CompletableFuture of void
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<Void> put(@javax.annotation.Nonnull final U body) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPutRequestInfo(
                body
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Update media content for the navigation property messages in users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param body Binary request body
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<Void> put(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPutRequestInfo(
                body
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Update media content for the navigation property messages in users
     * @param <U> The generic type for the model, must inherit InputStream
     * @param body Binary request body
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public <U extends InputStream> java.util.concurrent.CompletableFuture<Void> put(@javax.annotation.Nonnull final U body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPutRequestInfo(
                body, h
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
}
