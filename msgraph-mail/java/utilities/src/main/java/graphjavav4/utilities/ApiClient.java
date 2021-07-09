package graphjavav4.utilities;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactory;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import graphjavav4.utilities.users.item.UserRequestBuilder;
import graphjavav4.utilities.users.UsersRequestBuilder;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
public class ApiClient {
    /** Current path for the request  */
    @javax.annotation.Nullable
    public String currentPath;
    /** Core service to use to execute the requests  */
    @javax.annotation.Nullable
    public HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /** Factory to use to get a serializer for payload serialization  */
    @javax.annotation.Nullable
    public SerializationWriterFactory serializerFactory;
    @javax.annotation.Nonnull
    public UsersRequestBuilder users() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new UsersRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
    /**
     * Instantiates a new Api client and sets the default values.
     * @param httpCore The http core service to use to execute the requests.
     * @param serializationWriterFactory Factory to use to get a serializer for payload serialization.
     * @return a void
     */
    public ApiClient(@javax.annotation.Nonnull final HttpCore httpCore, @javax.annotation.Nullable final SerializationWriterFactory serializationWriterFactory) {
        Objects.requireNonNull(httpCore);
        this.pathSegment = "https://graph.microsoft.com/v1.0";
        this.httpCore = httpCore;
        ApiClientBuilder.registerDefaultSerializer(JsonSerializationWriterFactory.class);
        ApiClientBuilder.registerDefaultDeserializer(JsonParseNodeFactory.class);
        if(serializationWriterFactory == null && SerializationWriterFactoryRegistry.defaultInstance.contentTypeAssociatedFactories.isEmpty()) throw new RuntimeException("The Serialization Writer factory has not been initialized for this client.");
        if(ParseNodeFactoryRegistry.defaultInstance.contentTypeAssociatedFactories.isEmpty()) throw new RuntimeException("The Parse Node factory has not been initialized for this client.");
        this.serializerFactory = (serializationWriterFactory != null ? serializationWriterFactory : SerializationWriterFactoryRegistry.defaultInstance);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users collection
     * @param id Unique identifier of the item
     * @return a UserRequestBuilder
     */
    @javax.annotation.Nonnull
    public UserRequestBuilder users(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/users/" + id;
        final HttpCore parentCore = httpCore;
        final SerializationWriterFactory parentSerializationFactory = serializerFactory;
        return new UserRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; serializerFactory = parentSerializationFactory; }};
    }
}
