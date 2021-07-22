package graphjavav4.utilities.users.mailFolders.item;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.MailFolder;
import graphjavav4.utilities.users.mailFolders.childFolders.ChildFoldersRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messageRules.item.MessageRuleRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messageRules.MessageRulesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.MessagesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.multiValueExtendedProperties.item.MultiValueLegacyExtendedPropertyRequestBuilder;
import graphjavav4.utilities.users.mailFolders.multiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.singleValueExtendedProperties.item.SingleValueLegacyExtendedPropertyRequestBuilder;
import graphjavav4.utilities.users.mailFolders.singleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}  */
public class MailFolderRequestBuilder {
    @javax.annotation.Nonnull
    public ChildFoldersRequestBuilder childFolders() {
        return new ChildFoldersRequestBuilder(currentPath + pathSegment, httpCore);
    }
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    @javax.annotation.Nonnull
    public MessageRulesRequestBuilder messageRules() {
        return new MessageRulesRequestBuilder(currentPath + pathSegment, httpCore);
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(currentPath + pathSegment, httpCore);
    }
    @javax.annotation.Nonnull
    public MultiValueExtendedPropertiesRequestBuilder multiValueExtendedProperties() {
        return new MultiValueExtendedPropertiesRequestBuilder(currentPath + pathSegment, httpCore);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    @javax.annotation.Nonnull
    public SingleValueExtendedPropertiesRequestBuilder singleValueExtendedProperties() {
        return new SingleValueExtendedPropertiesRequestBuilder(currentPath + pathSegment, httpCore);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.mailFolders.childFolders collection
     * @param id Unique identifier of the item
     * @return a MailFolderRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder childFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MailFolderRequestBuilder(currentPath + pathSegment + "/childFolders/" + id, httpCore);
    }
    /**
     * Instantiates a new MailFolderRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public MailFolderRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final MailFolder body) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.setContentFromParsable(body, httpCore, "application/json");
        return requestInfo;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @return a RequestInfo
     */
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath + pathSegment);
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.setContentFromParsable(body, httpCore, "application/json");
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        return requestInfo;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
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
     * The user's mail folders. Read-only. Nullable.
     * @return a CompletableFuture of MailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, MailFolder.class, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, MailFolder.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                h
            );
            return this.httpCore.sendAsync(requestInfo, MailFolder.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of MailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                q, h
            );
            return this.httpCore.sendAsync(requestInfo, MailFolder.class, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.mailFolders.messageRules collection
     * @param id Unique identifier of the item
     * @return a MessageRuleRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRuleRequestBuilder messageRules(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MessageRuleRequestBuilder(currentPath + pathSegment + "/messageRules/" + id, httpCore);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.mailFolders.messages collection
     * @param id Unique identifier of the item
     * @return a MessageRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MessageRequestBuilder(currentPath + pathSegment + "/messages/" + id, httpCore);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.mailFolders.multiValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @return a MultiValueLegacyExtendedPropertyRequestBuilder
     */
    @javax.annotation.Nonnull
    public MultiValueLegacyExtendedPropertyRequestBuilder multiValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MultiValueLegacyExtendedPropertyRequestBuilder(currentPath + pathSegment + "/multiValueExtendedProperties/" + id, httpCore);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body) {
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
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final ResponseHandler responseHandler) {
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
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final ResponseHandler responseHandler) {
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
     * Gets an item from the graphjavav4.utilities.users.mailFolders.singleValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @return a SingleValueLegacyExtendedPropertyRequestBuilder
     */
    @javax.annotation.Nonnull
    public SingleValueLegacyExtendedPropertyRequestBuilder singleValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new SingleValueLegacyExtendedPropertyRequestBuilder(currentPath + pathSegment + "/singleValueExtendedProperties/" + id, httpCore);
    }
    /** The user's mail folders. Read-only. Nullable.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Expand related entities  */
        @javax.annotation.Nullable
        public String[] expand;
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
    }
}
