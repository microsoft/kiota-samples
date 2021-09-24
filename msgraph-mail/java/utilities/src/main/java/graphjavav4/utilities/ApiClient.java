package graphjavav4.utilities;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.MiddlewareOption;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import graphjavav4.utilities.users.item.UserRequestBuilder;
import graphjavav4.utilities.users.UsersRequestBuilder;
import java.io.InputStream;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
public class ApiClient {
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    @javax.annotation.Nonnull
    public UsersRequestBuilder users() {
        return new UsersRequestBuilder(pathSegment, httpCore, false);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public ApiClient(@javax.annotation.Nonnull final HttpCore httpCore) {
        Objects.requireNonNull(httpCore);
        this.pathSegment = "https://graph.microsoft.com/v1.0";
        this.httpCore = httpCore;
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
        return new UserRequestBuilder(pathSegment + "/users/" + id, httpCore, false);
    }
}
