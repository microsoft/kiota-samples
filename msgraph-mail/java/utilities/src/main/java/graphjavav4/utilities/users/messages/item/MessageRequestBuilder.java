package graphjavav4.utilities.users.messages.item;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactory;
import graphjavav4.utilities.users.Message;
import graphjavav4.utilities.users.messages.attachments.AttachmentsRequestBuilder;
import graphjavav4.utilities.users.messages.attachments.item.AttachmentRequestBuilder;
import graphjavav4.utilities.users.messages.Content.ContentRequestBuilder;
import graphjavav4.utilities.users.messages.extensions.ExtensionsRequestBuilder;
import graphjavav4.utilities.users.messages.extensions.item.ExtensionRequestBuilder;
import graphjavav4.utilities.users.messages.multiValueExtendedProperties.item.MultiValueLegacyExtendedPropertyRequestBuilder;
import graphjavav4.utilities.users.messages.multiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder;
import graphjavav4.utilities.users.messages.singleValueExtendedProperties.item.SingleValueLegacyExtendedPropertyRequestBuilder;
import graphjavav4.utilities.users.messages.singleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}  */
public class MessageRequestBuilder {
    @javax.annotation.Nonnull
    public AttachmentsRequestBuilder attachments() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new AttachmentsRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    @javax.annotation.Nonnull
    public ContentRequestBuilder content() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new ContentRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /** Current path for the request  */
    @javax.annotation.Nullable
    public String currentPath;
    @javax.annotation.Nonnull
    public ExtensionsRequestBuilder extensions() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new ExtensionsRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /** Core service to use to execute the requests  */
    @javax.annotation.Nullable
    public HttpCore httpCore;
    @javax.annotation.Nonnull
    public MultiValueExtendedPropertiesRequestBuilder multiValueExtendedProperties() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new MultiValueExtendedPropertiesRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /** Path segment to use to build the URL for the current request builder  */
    @javax.annotation.Nonnull
    private final String pathSegment = "";
    /** Factory to use to get a serializer for payload serialization  */
    @javax.annotation.Nullable
    public SerializationWriterFactory serializerFactory;
    @javax.annotation.Nonnull
    public SingleValueExtendedPropertiesRequestBuilder singleValueExtendedProperties() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new SingleValueExtendedPropertiesRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.messages.attachments collection
     * @param id Unique identifier of the item
     * @return a AttachmentRequestBuilder
     */
    @javax.annotation.Nonnull
    public AttachmentRequestBuilder attachments(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/attachments/" + id;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new AttachmentRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /**
     * Delete navigation property messages for users
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo() throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.DELETE;
        }};
        return requestInfo;
    }
    /**
     * Delete navigation property messages for users
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.DELETE;
        }};
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * Get messages from users
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
     * Get messages from users
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
     * Get messages from users
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
     * Update the navigation property messages in users
     * @param body 
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final Message body) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "application/json");
        return requestInfo;
    }
    /**
     * Update the navigation property messages in users
     * @param body 
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final Message body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.setContentFromParsable(body, serializerFactory, "application/json");
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * Delete navigation property messages for users
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete() {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Delete navigation property messages for users
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Delete navigation property messages for users
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
                h
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.messages.extensions collection
     * @param id Unique identifier of the item
     * @return a ExtensionRequestBuilder
     */
    @javax.annotation.Nonnull
    public ExtensionRequestBuilder extensions(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/extensions/" + id;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new ExtensionRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /**
     * Get messages from users
     * @return a CompletableFuture of Message
     */
    public java.util.concurrent.CompletableFuture<Message> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, Message.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get messages from users
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of Message
     */
    public java.util.concurrent.CompletableFuture<Message> get(@javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, Message.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get messages from users
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of Message
     */
    public java.util.concurrent.CompletableFuture<Message> get(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                h
            );
            return this.httpCore.sendAsync(requestInfo, Message.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Get messages from users
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of Message
     */
    public java.util.concurrent.CompletableFuture<Message> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                q, h
            );
            return this.httpCore.sendAsync(requestInfo, Message.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.messages.multiValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @return a MultiValueLegacyExtendedPropertyRequestBuilder
     */
    @javax.annotation.Nonnull
    public MultiValueLegacyExtendedPropertyRequestBuilder multiValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/multiValueExtendedProperties/" + id;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new MultiValueLegacyExtendedPropertyRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /**
     * Update the navigation property messages in users
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Message body) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Update the navigation property messages in users
     * @param body 
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Message body, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Update the navigation property messages in users
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Message body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body, h
            );
            return this.httpCore.sendPrimitiveAsync(requestInfo, Void.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.messages.singleValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @return a SingleValueLegacyExtendedPropertyRequestBuilder
     */
    @javax.annotation.Nonnull
    public SingleValueLegacyExtendedPropertyRequestBuilder singleValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/singleValueExtendedProperties/" + id;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new SingleValueLegacyExtendedPropertyRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /** Get messages from users  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Expand related entities  */
        @javax.annotation.Nullable
        public String[] expand;
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
    }
}
