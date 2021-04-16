package graphjavav4.utilities.users.mailFolders.item;

import java.util.Objects;
import graphjavav4.utilities.users.mailFolders.childFolders.ChildFoldersRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messageRules.MessageRulesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.MessagesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.multiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.singleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.MailFolder;
import graphjavav4.utilities.users.mailFolders.messageRules.item.MessageRuleRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.multiValueExtendedProperties.item.MultiValueLegacyExtendedPropertyRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.singleValueExtendedProperties.item.SingleValueLegacyExtendedPropertyRequestBuilder;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.QueryParametersBase;
import java.util.Map;
import java.net.URI;
import java.net.URISyntaxException;
public class MailFolderRequestBuilder {
    @javax.annotation.Nonnull
    public ChildFoldersRequestBuilder childFolders() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new ChildFoldersRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessageRulesRequestBuilder messageRules() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new MessageRulesRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new MessagesRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MultiValueExtendedPropertiesRequestBuilder multiValueExtendedProperties() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new MultiValueExtendedPropertiesRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public SingleValueExtendedPropertiesRequestBuilder singleValueExtendedProperties() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new SingleValueExtendedPropertiesRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nonnull final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(q);
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            final GetQueryParameters qParams = new GetQueryParameters();
            q.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    private final String pathSegment = "";
    @javax.annotation.Nullable
    public String currentPath;
    @javax.annotation.Nullable
    public HttpCore httpCore;
    public class GetQueryParameters extends QueryParametersBase {
        @javax.annotation.Nullable
        public String[] select;
        @javax.annotation.Nullable
        public String[] expand;
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<MailFolder> get() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder childFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/childFolders/" + id;
        return new MailFolderRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessageRuleRequestBuilder messageRules(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/messageRules/" + id;
        return new MessageRuleRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/messages/" + id;
        return new MessageRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MultiValueLegacyExtendedPropertyRequestBuilder multiValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/multiValueExtendedProperties/" + id;
        return new MultiValueLegacyExtendedPropertyRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public SingleValueLegacyExtendedPropertyRequestBuilder singleValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/singleValueExtendedProperties/" + id;
        return new SingleValueLegacyExtendedPropertyRequestBuilder() {{ currentPath = parentPath; }};
    }
}
