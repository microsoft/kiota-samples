package graphjavav4.utilities.users.item.inferenceclassification;

import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import graphjavav4.utilities.models.microsoft.graph.InferenceClassification;
import graphjavav4.utilities.users.item.inferenceclassification.overrides.item.InferenceClassificationOverrideItemRequestBuilder;
import graphjavav4.utilities.users.item.inferenceclassification.overrides.OverridesRequestBuilder;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/inferenceClassification  */
public class InferenceClassificationRequestBuilder {
    @javax.annotation.Nonnull
    public OverridesRequestBuilder overrides() {
        return new OverridesRequestBuilder(pathParameters, requestAdapter);
    }
    /** Path parameters for the request  */
    private final HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public InferenceClassificationRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/users/{user_id}/inferenceClassification{?select}";
        var urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public InferenceClassificationRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/users/{user_id}/inferenceClassification{?select}";
        var urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation() throws URISyntaxException {
        return createDeleteRequestInformation(null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createDeleteRequestInformation(h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.DELETE;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addRequestOptions(o.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation() throws URISyntaxException {
        return createGetRequestInformation(null, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param q Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) throws URISyntaxException {
        return createGetRequestInformation(q, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param q Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createGetRequestInformation(q, h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.GET;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (q != null) {
            final GetQueryParameters qParams = new GetQueryParameters();
            q.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
        }
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addRequestOptions(o.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final InferenceClassification body) throws URISyntaxException {
        return createPatchRequestInformation(body, null, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        return createPatchRequestInformation(body, h, null);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (h != null) {
            h.accept(requestInfo.headers);
        }
        if (o != null) {
            requestInfo.addRequestOptions(o.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
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
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(h, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(h, o);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(h, o);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @return a CompletableFuture of inferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get() {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(null, null, null);
            return this.requestAdapter.sendAsync(requestInfo, InferenceClassification::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param q Request query parameters
     * @return a CompletableFuture of inferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(q, null, null);
            return this.requestAdapter.sendAsync(requestInfo, InferenceClassification::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param q Request query parameters
     * @return a CompletableFuture of inferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(q, h, null);
            return this.requestAdapter.sendAsync(requestInfo, InferenceClassification::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @return a CompletableFuture of inferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(q, h, o);
            return this.requestAdapter.sendAsync(requestInfo, InferenceClassification::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of inferenceClassification
     */
    public java.util.concurrent.CompletableFuture<InferenceClassification> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(q, h, o);
            return this.requestAdapter.sendAsync(requestInfo, InferenceClassification::createFromDiscriminatorValue, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.inferenceClassification.overrides.item collection
     * @param id Unique identifier of the item
     * @return a inferenceClassificationOverrideItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public InferenceClassificationOverrideItemRequestBuilder overrides(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("inferenceClassificationOverride_id", id);
        return new InferenceClassificationOverrideItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, null, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
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
            final RequestInformation requestInfo = createPatchRequestInformation(body, h, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, h, o);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final InferenceClassification body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nullable final Collection<RequestOption> o, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, h, o);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Select properties to be returned  */
        @javax.annotation.Nullable
        public String[] select;
    }
}
