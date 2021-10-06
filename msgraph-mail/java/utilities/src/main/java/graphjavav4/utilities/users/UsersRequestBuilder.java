package graphjavav4.utilities.users;

import com.microsoft.kiota.RequestAdapter;
import java.util.Objects;
/** Builds and executes requests for operations under /users  */
public class UsersRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    /** Whether the current path is a raw URL  */
    private final boolean isRawUrl;
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /** The http core service to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param requestAdapter The http core service to use to execute the requests.
     * @return a void
     */
    public UsersRequestBuilder(final String currentPath, final RequestAdapter requestAdapter) {
        this(currentPath, requestAdapter, true);
    }
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param isRawUrl Whether the current path is a raw URL
     * @param requestAdapter The http core service to use to execute the requests.
     * @return a void
     */
    public UsersRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final RequestAdapter requestAdapter, final boolean isRawUrl) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(requestAdapter);
        this.pathSegment = "/users";
        this.requestAdapter = requestAdapter;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    }
}
