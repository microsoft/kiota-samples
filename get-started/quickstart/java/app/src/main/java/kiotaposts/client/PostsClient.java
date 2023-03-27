package kiotaposts.client;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.FormParseNodeFactory;
import com.microsoft.kiota.serialization.FormSerializationWriterFactory;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import com.microsoft.kiota.serialization.TextParseNodeFactory;
import com.microsoft.kiota.serialization.TextSerializationWriterFactory;
import java.util.HashMap;
import java.util.Objects;
import kiotaposts.client.posts.item.PostItemRequestBuilder;
import kiotaposts.client.posts.PostsRequestBuilder;
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
public class PostsClient {
    /** Path parameters for the request */
    private HashMap<String, Object> pathParameters;
    /** The posts property */
    @javax.annotation.Nonnull
    public PostsRequestBuilder posts() {
        return new PostsRequestBuilder(pathParameters, requestAdapter);
    }
    /** The request adapter to use to execute the requests. */
    private RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private String urlTemplate;
    /**
     * Instantiates a new PostsClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public PostsClient(@javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(requestAdapter);
        this.pathParameters = new HashMap<>();
        this.urlTemplate = "{+baseurl}";
        this.requestAdapter = requestAdapter;
        ApiClientBuilder.registerDefaultSerializer(JsonSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultSerializer(TextSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultSerializer(FormSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(JsonParseNodeFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(FormParseNodeFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(TextParseNodeFactory.class);
        if (requestAdapter.getBaseUrl() == null || requestAdapter.getBaseUrl().isEmpty()) {
            requestAdapter.setBaseUrl("https://jsonplaceholder.typicode.com");
        }
        pathParameters.put("baseurl", requestAdapter.getBaseUrl());
    }
    /**
     * Gets an item from the kiotaposts.client.posts.item collection
     * @param id Unique identifier of the item
     * @return a PostItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public PostItemRequestBuilder posts(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("post%2Did", id);
        return new PostItemRequestBuilder(urlTplParams, requestAdapter);
    }
}
