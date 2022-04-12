package graphjavav4.utilities.users.item.mailfolders.item;

import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParameter;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import graphjavav4.utilities.models.microsoft.graph.MailFolder;
import graphjavav4.utilities.users.item.mailfolders.item.childfolders.ChildFoldersRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.messagerules.item.MessageRuleItemRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.messagerules.MessageRulesRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.messages.item.MessageItemRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.messages.MessagesRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.multivalueextendedproperties.item.MultiValueLegacyExtendedPropertyItemRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.multivalueextendedproperties.MultiValueExtendedPropertiesRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.singlevalueextendedproperties.item.SingleValueLegacyExtendedPropertyItemRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.singlevalueextendedproperties.SingleValueExtendedPropertiesRequestBuilder;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}  */
public class MailFolderItemRequestBuilder {
    /** The childFolders property  */
    @javax.annotation.Nonnull
    public ChildFoldersRequestBuilder childFolders() {
        return new ChildFoldersRequestBuilder(pathParameters, requestAdapter);
    }
    /** The messageRules property  */
    @javax.annotation.Nonnull
    public MessageRulesRequestBuilder messageRules() {
        return new MessageRulesRequestBuilder(pathParameters, requestAdapter);
    }
    /** The messages property  */
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(pathParameters, requestAdapter);
    }
    /** The multiValueExtendedProperties property  */
    @javax.annotation.Nonnull
    public MultiValueExtendedPropertiesRequestBuilder multiValueExtendedProperties() {
        return new MultiValueExtendedPropertiesRequestBuilder(pathParameters, requestAdapter);
    }
    /** Path parameters for the request  */
    private final HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** The singleValueExtendedProperties property  */
    @javax.annotation.Nonnull
    public SingleValueExtendedPropertiesRequestBuilder singleValueExtendedProperties() {
        return new SingleValueExtendedPropertiesRequestBuilder(pathParameters, requestAdapter);
    }
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.childFolders.item collection
     * @param id Unique identifier of the item
     * @return a mailFolderItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderItemRequestBuilder childFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("mailFolder%2Did1", id);
        return new MailFolderItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new MailFolderItemRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public MailFolderItemRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24select}";
        var urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new MailFolderItemRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public MailFolderItemRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24select}";
        var urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation() throws URISyntaxException {
        return createDeleteRequestInformation(null, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createDeleteRequestInformation(headers, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.DELETE;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation() throws URISyntaxException {
        return createGetRequestInformation(null, null, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters) throws URISyntaxException {
        return createGetRequestInformation(queryParameters, null, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createGetRequestInformation(queryParameters, headers, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.GET;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (queryParameters != null) {
            final GetQueryParameters qParams = new GetQueryParameters();
            queryParameters.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
        }
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final MailFolder body) throws URISyntaxException {
        return createPatchRequestInformation(body, null, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createPatchRequestInformation(body, headers, null);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete() {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(null, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @return a CompletableFuture of mailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get() {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(null, null, null);
            return this.requestAdapter.sendAsync(requestInfo, MailFolder::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of mailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, null, null);
            return this.requestAdapter.sendAsync(requestInfo, MailFolder::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of mailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, null);
            return this.requestAdapter.sendAsync(requestInfo, MailFolder::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of mailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, options);
            return this.requestAdapter.sendAsync(requestInfo, MailFolder::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of mailFolder
     */
    public java.util.concurrent.CompletableFuture<MailFolder> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, options);
            return this.requestAdapter.sendAsync(requestInfo, MailFolder::createFromDiscriminatorValue, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.messageRules.item collection
     * @param id Unique identifier of the item
     * @return a messageRuleItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRuleItemRequestBuilder messageRules(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("messageRule%2Did", id);
        return new MessageRuleItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.messages.item collection
     * @param id Unique identifier of the item
     * @return a messageItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageItemRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("message%2Did", id);
        return new MessageItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.multiValueExtendedProperties.item collection
     * @param id Unique identifier of the item
     * @return a multiValueLegacyExtendedPropertyItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MultiValueLegacyExtendedPropertyItemRequestBuilder multiValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("multiValueLegacyExtendedProperty%2Did", id);
        return new MultiValueLegacyExtendedPropertyItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, null, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final MailFolder body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.singleValueExtendedProperties.item collection
     * @param id Unique identifier of the item
     * @return a singleValueLegacyExtendedPropertyItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public SingleValueLegacyExtendedPropertyItemRequestBuilder singleValueExtendedProperties(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("singleValueLegacyExtendedProperty%2Did", id);
        return new SingleValueLegacyExtendedPropertyItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /** The user's mail folders. Read-only. Nullable.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Select properties to be returned  */
        @QueryParameter(name = "%24select")
        @javax.annotation.Nullable
        public String[] select;
    }
}
