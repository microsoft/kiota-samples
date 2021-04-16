package graphjavav4.utilities;

import java.util.Objects;
import graphjavav4.utilities.users.UsersRequestBuilder;
import graphjavav4.utilities.users.item.UserRequestBuilder;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.QueryParametersBase;
import java.util.Map;
import java.net.URI;
import java.net.URISyntaxException;
import java.io.InputStream;
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
    @javax.annotation.Nonnull
    public UserRequestBuilder users(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/users/" + id;
        final HttpCore parentCore = httpCore;
        return new UserRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
}
