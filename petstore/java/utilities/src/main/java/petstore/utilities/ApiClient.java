package petstore.utilities;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.FormParseNodeFactory;
import com.microsoft.kiota.serialization.FormSerializationWriterFactory;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.MultipartSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import com.microsoft.kiota.serialization.TextParseNodeFactory;
import com.microsoft.kiota.serialization.TextSerializationWriterFactory;
import java.util.HashMap;
import java.util.Objects;
import petstore.utilities.pet.PetRequestBuilder;
import petstore.utilities.store.StoreRequestBuilder;
import petstore.utilities.user.UserRequestBuilder;
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ApiClient extends BaseRequestBuilder {
    /**
     * The pet property
     * @return a {@link PetRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public PetRequestBuilder pet() {
        return new PetRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The store property
     * @return a {@link StoreRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public StoreRequestBuilder store() {
        return new StoreRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The user property
     * @return a {@link UserRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public UserRequestBuilder user() {
        return new UserRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link ApiClient} and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ApiClient(@jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}");
        this.pathParameters = new HashMap<>();
        ApiClientBuilder.registerDefaultSerializer(() -> new JsonSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new TextSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new FormSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new MultipartSerializationWriterFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new JsonParseNodeFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new FormParseNodeFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new TextParseNodeFactory());
        if (requestAdapter.getBaseUrl() == null || requestAdapter.getBaseUrl().isEmpty()) {
            requestAdapter.setBaseUrl("https://petstore.swagger.io/v2");
        }
        pathParameters.put("baseurl", requestAdapter.getBaseUrl());
    }
}
