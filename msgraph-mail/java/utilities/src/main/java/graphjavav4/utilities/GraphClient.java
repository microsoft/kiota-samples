package graphjavav4.utilities;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.item.UserRequestBuilder;
import graphjavav4.utilities.users.UsersRequestBuilder;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
public class GraphClient {
    @javax.annotation.Nonnull
    public UsersRequestBuilder users() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        return new UsersRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
    @javax.annotation.Nonnull
    private final String pathSegment = "https://graph.microsoft.com/v1.0";
    @javax.annotation.Nullable
    public String currentPath;
    @javax.annotation.Nullable
    public HttpCore httpCore;
    @javax.annotation.Nullable
    public Function<String, SerializationWriter> serializerFactory;
    @javax.annotation.Nonnull
    public UserRequestBuilder users(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/users/" + id;
        final HttpCore parentCore = httpCore;
        return new UserRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
}
