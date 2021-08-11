package graphjavav4.utilities.users;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.MiddlewareOption;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users  */
public class UsersRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** The http core service to use to execute the requests.  */
    private final HttpCore httpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @return a void
     */
    public UsersRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final HttpCore httpCore) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(httpCore);
        this.pathSegment = "/users";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
    }
}
