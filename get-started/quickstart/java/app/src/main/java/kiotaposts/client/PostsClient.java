package kiotaposts.client;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import java.util.HashMap;
import java.util.Objects;
import kiotaposts.client.posts.PostsRequestBuilder;
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class PostsClient extends BaseRequestBuilder {
    /**
     * The posts property
     * @return a {@link PostsRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public PostsRequestBuilder posts() {
        return new PostsRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link PostsClient} and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public PostsClient(@jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}");
        this.pathParameters = new HashMap<>();
        if (requestAdapter.getBaseUrl() == null || requestAdapter.getBaseUrl().isEmpty()) {
            requestAdapter.setBaseUrl("https://jsonplaceholder.typicode.com");
        }
        pathParameters.put("baseurl", requestAdapter.getBaseUrl());
    }
}
