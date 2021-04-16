package graphjavav4.utilities.users.inferenceClassification.overrides.item;

import java.util.Objects;
import graphjavav4.utilities.users.inferenceClassification.InferenceClassificationOverride;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.QueryParametersBase;
import java.util.Map;
import java.net.URI;
import java.net.URISyntaxException;
import java.io.InputStream;
public class InferenceClassificationOverrideRequestBuilder {
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<InferenceClassificationOverride> get(@javax.annotation.Nonnull final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(q);
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                q, h
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nonnull final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(q);
        Objects.requireNonNull(h);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.GET;
        }};
        final GetQueryParameters qParams = new GetQueryParameters();
        q.accept(qParams);
        qParams.AddQueryParameters(requestInfo.queryParameters);
        h.accept(requestInfo.headers);
        return requestInfo;
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final InferenceClassificationOverride body, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body, h
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final InferenceClassificationOverride body, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(body);
        Objects.requireNonNull(h);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.PATCH;
            content = (InputStream)(Object)body;
        }};
        h.accept(requestInfo.headers);
        return requestInfo;
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
                h
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(h);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.DELETE;
        }};
        h.accept(requestInfo.headers);
        return requestInfo;
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
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<InferenceClassificationOverride> get(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
                h
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<InferenceClassificationOverride> get(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final InferenceClassificationOverride body, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<InferenceClassificationOverride> get() {
        try {
            final RequestInfo requestInfo = createGetRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final InferenceClassificationOverride body) {
        Objects.requireNonNull(body);
        try {
            final RequestInfo requestInfo = createPatchRequestInfo(
                body
            );
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Object> delete() {
        try {
            final RequestInfo requestInfo = createDeleteRequestInfo(
            );
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h) throws URISyntaxException {
        Objects.requireNonNull(h);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.GET;
        }};
        h.accept(requestInfo.headers);
        return requestInfo;
    }
    @javax.annotation.Nonnull
    public RequestInfo createGetRequestInfo() throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.GET;
        }};
        return requestInfo;
    }
    @javax.annotation.Nonnull
    public RequestInfo createPatchRequestInfo(@javax.annotation.Nonnull final InferenceClassificationOverride body) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.PATCH;
            content = (InputStream)(Object)body;
        }};
        return requestInfo;
    }
    @javax.annotation.Nonnull
    public RequestInfo createDeleteRequestInfo() throws URISyntaxException {
        final RequestInfo requestInfo = new RequestInfo() {{
            uri = new URI(currentPath);
            httpMethod = HttpMethod.DELETE;
        }};
        return requestInfo;
    }
}
