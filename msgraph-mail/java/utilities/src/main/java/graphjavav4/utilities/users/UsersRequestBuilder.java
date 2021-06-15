package graphjavav4.utilities.users;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.SerializationWriterFactory;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
/** Builds and executes requests for operations under /users  */
public class UsersRequestBuilder {
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
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @return a void
     */
    public UsersRequestBuilder() {
        this.pathSegment = "/users";
    }
}
