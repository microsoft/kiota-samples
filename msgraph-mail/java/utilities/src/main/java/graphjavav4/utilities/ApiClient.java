package graphjavav4.utilities;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import graphjavav4.utilities.users.item.UserRequestBuilder;
import graphjavav4.utilities.users.UsersRequestBuilder;
import java.util.Objects;
/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
public class ApiClient {
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /** The http core service to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    @javax.annotation.Nonnull
    public UsersRequestBuilder users() {
        return new UsersRequestBuilder(pathSegment, requestAdapter, false);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The http core service to use to execute the requests.
     * @return a void
     */
    public ApiClient(@javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(requestAdapter);
        this.pathSegment = "https://graph.microsoft.com/v1.0";
        this.requestAdapter = requestAdapter;
        ApiClientBuilder.registerDefaultSerializer(JsonSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(JsonParseNodeFactory.class);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item collection
     * @param id Unique identifier of the item
     * @return a userRequestBuilder
     */
    @javax.annotation.Nonnull
    public UserRequestBuilder users(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new UserRequestBuilder(pathSegment + "/users/" + id, requestAdapter, false);
    }
}
