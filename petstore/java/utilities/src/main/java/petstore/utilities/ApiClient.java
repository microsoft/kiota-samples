package petstore.utilities;

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
import petstore.utilities.pet.item.WithPetItemRequestBuilder;
import petstore.utilities.pet.PetRequestBuilder;
import petstore.utilities.store.StoreRequestBuilder;
import petstore.utilities.user.item.WithUsernameItemRequestBuilder;
import petstore.utilities.user.UserRequestBuilder;
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
public class ApiClient {
    /** Path parameters for the request */
    private HashMap<String, Object> pathParameters;
    /** The pet property */
    @javax.annotation.Nonnull
    public PetRequestBuilder pet() {
        return new PetRequestBuilder(pathParameters, requestAdapter);
    }
    /** The request adapter to use to execute the requests. */
    private RequestAdapter requestAdapter;
    /** The store property */
    @javax.annotation.Nonnull
    public StoreRequestBuilder store() {
        return new StoreRequestBuilder(pathParameters, requestAdapter);
    }
    /** Url template to use to build the URL for the current request builder */
    private String urlTemplate;
    /** The user property */
    @javax.annotation.Nonnull
    public UserRequestBuilder user() {
        return new UserRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public ApiClient(@javax.annotation.Nonnull final RequestAdapter requestAdapter) {
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
            requestAdapter.setBaseUrl("https://petstore.swagger.io/v2");
        }
        pathParameters.put("baseurl", requestAdapter.getBaseUrl());
    }
    /**
     * Gets an item from the petstore.utilities.pet.item collection
     * @param id Unique identifier of the item
     * @return a WithPetItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public WithPetItemRequestBuilder pet(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("petId", id);
        return new WithPetItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Gets an item from the petstore.utilities.user.item collection
     * @param id Unique identifier of the item
     * @return a WithUsernameItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public WithUsernameItemRequestBuilder user(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("username", id);
        return new WithUsernameItemRequestBuilder(urlTplParams, requestAdapter);
    }
}
